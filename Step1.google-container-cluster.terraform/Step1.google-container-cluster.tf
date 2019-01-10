resource "google_container_cluster" "primary" {
  name                     = "diceroller-cluster"
  remove_default_node_pool = true
  initial_node_count = 3

  lifecycle{
    create_before_destroy = true
  }

  master_auth {
    username = "${var.username}"
    password = "${var.password}"
  }
}

resource "google_container_node_pool" "primary_pool" {
  name       = "primary-pool"
  cluster    = "${google_container_cluster.primary.name}"
  node_count = "3"

  lifecycle { 
    create_before_destroy = true
  }

  node_config {
    machine_type = "f1-micro"
    preemptible  = true
  }

  management {
    auto_repair  = true
    auto_upgrade = true
  }

  autoscaling {
    min_node_count = 3
    max_node_count = 5
  }
}

resource "local_file" "cluster-name" {
  # outputting the cluster name so that Step 2 can read it.
  content  = "${google_container_cluster.primary.name}"
  filename = "./../terraform-data/cluster-name.tfdata"
}