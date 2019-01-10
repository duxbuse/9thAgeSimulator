locals {
  local-project = "${var.project ? var.project: data.local_file.project.content }"
  local-region = "${var.region ? var.region : data.local_file.region.content }"
  local-zone = "${var.zone ? var.zone : data.local_file.zone.content }"
  local-repository = "${var.repository ? var.repository : data.local_file.deployment-name.content }"
  local-triggers = [{ branch = "${var.branch}" }]
}
variable "project" {
  default = false
}variable "region" {
  default = false
}variable "zone" {
  default = false
}variable "repository" {
  description = "Name of mirror repository on GCP"
  default     = false
}variable "branch" {
  type = "string"
  description = "Branch to trigger on"
  default     =  "master"
}

data "local_file" "project"{
  filename = "./../terraform-data/project.tfdata"
}data "local_file" "region"{
  filename = "./../terraform-data/region.tfdata"
}data "local_file" "zone"{
  filename = "./../terraform-data/zone.tfdata"
}data "local_file" "deployment-name"{
  filename = "./../terraform-data/deployment-name.tfdata"
}