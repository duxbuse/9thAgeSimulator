data "google_project" "project" {}

data "local_file" "cluster-name"{
  filename = "./../terraform-data/cluster-name.tfdata"
}
resource "google_sourcerepo_repository" "new_git_repository" {
  name = "${local.local-repository}"
}

resource "google_cloudbuild_trigger" "new_git_build_trigger" {
  count       = "${length(local.local-triggers)}"
  description = "Trigger Git repository ${local.local-repository} / ${lookup(local.local-triggers[count.index], "branch", "${var.branch}")}"

  trigger_template {
    project     = "${data.google_project.project.project_id}"
    branch_name = "${lookup(local.local-triggers[count.index], "branch", "${var.branch}")}"
    repo_name   = "${local.local-repository}"
  }

  substitutions {
    _TAG             = "${lookup(local.local-triggers[count.index], "branch", "${var.branch}")}"
    _COMPUTE_ZONE = "${local.local-zone}"
    _CLUSTER = "${data.local_file.cluster-name.content}"
  }

  filename = "cloudbuild.yaml"

  # Google Git repository has been created.
  depends_on = [
    "google_sourcerepo_repository.new_git_repository",
  ]
}
