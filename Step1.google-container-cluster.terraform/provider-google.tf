provider "google" {
  project     = "${var.project}"
  credentials = "${file("diceroller-220503-8497483a16e9.json")}"
  region      = "${var.region}"
  zone        = "${var.zone}"
}