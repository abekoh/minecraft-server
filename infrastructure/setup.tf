terraform {
  required_version = ">= 0.13"
  required_providers {
    google = "3.76.0"
  }
  backend "gcs" {
    bucket = "minecraft-abekoh-terraform"
  }
}

provider "google" {
  credentials = file(var.credentials_file)
  project     = var.project
  region      = var.region
  zone        = var.zone
}
