provider "google" {
  project     = "${var.project}"
  credentials = "${file("rosy-element-228200-a5e4b2b2ac79.json")}"
  region      = "${var.region}"
  zone        = "${var.zone}"
}