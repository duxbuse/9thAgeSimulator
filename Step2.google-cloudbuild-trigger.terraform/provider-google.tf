provider "google" {
  project     = "${local.local-project}"
  credentials = "${file("rosy-element-228200-a5e4b2b2ac79.json")}"
  region      = "${local.local-region}"
  zone        = "${local.local-zone}"
}