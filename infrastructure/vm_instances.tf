terraform {
  required_version = ">= 0.13"
  required_providers {
    google = "3.76.0"
  }
}

provider "google" {
  credentials = file(var.credentials_file)
  project     = var.project
  region      = var.region
  zone        = var.zone
}

resource "google_compute_network" "vpc_network" {
  name = "minecraft-network"
}

data "google_compute_image" "base_image" {
  family  = "ubuntu-2004-lts"
  project = "ubuntu-os-cloud"
}

resource "google_compute_instance" "vm_instance" {
  name         = "minecraft-abekoh"
  machine_type = "e2-medium"
  tags         = ["minecraft"]

  boot_disk {
    auto_delete = false
    device_name = "minecraft-abekoh-disk"
    mode        = "READ_WRITE"

    initialize_params {
      image = data.google_compute_image.base_image.self_link
      size  = 30
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.name
    access_config {
    }
  }

  scheduling {
    preemptible = true
    automatic_restart = false
  }
}