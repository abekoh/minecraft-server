terraform {
  required_version = ">= 0.13"
  required_providers {
    google = "3.76.0"
  }
}

provider "google" {
  credentials = file(var.credentials_file)
  project = var.project
  region  = var.region
  zone    = var.zone
}

resource "google_compute_network" "vpc_network" {
  name = "minecraft-network"
}
