variable "username" {
  default = "admin"
}
variable "password" {
  default = "adminpassword1234567890"
}
variable "project" {
  default = "diceroller-220503"
}
resource "local_file" "project" {
  content  = "${var.project}"
  filename = "./../terraform-data/project.tfdata"
}

variable "region" {
  default = "us-central1"
}
resource "local_file" "region" {
  content  = "${var.region}"
  filename = "./../terraform-data/region.tfdata"
}

variable "zone" {
  default = "us-central1-c"
}
resource "local_file" "zone" {
  content  = "${var.zone}"
  filename = "./../terraform-data/zone.tfdata"
}
variable "deployment-name" {
  default = "diceroller"
}
resource "local_file" "deployment-name" {
  content  = "${var.deployment-name}"
  filename = "./../terraform-data/deployment-name.tfdata"
}
