provider "google" {
  project     = "${local.local-project}"
  credentials = "${file("diceroller-220503-8497483a16e9.json")}"
  region      = "${local.local-region}"
  zone        = "${local.local-zone}"
}