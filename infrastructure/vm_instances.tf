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

resource "google_compute_firewall" "allow_ssh" {
  name    = "allow-ssh"
  network = google_compute_network.vpc_network.name
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }
}

resource "google_compute_firewall" "allow_icmp" {
  name    = "allow-icmp"
  network = google_compute_network.vpc_network.name
  allow {
    protocol = "icmp"
  }
}

resource "google_compute_firewall" "allow_minecraft" {
  name    = "allow-minecraft"
  network = google_compute_network.vpc_network.name
  allow {
    protocol = "tcp"
    ports    = ["25565"]
  }
  target_tags = ["minecraft-server"]
}

resource "google_compute_disk" "game_data_disk" {
  name = "game-data-disk"
  type = "pd-standard"
  size = 30
}

data "google_compute_image" "base_image" {
  family  = "ubuntu-2004-lts"
  project = "ubuntu-os-cloud"
}

resource "google_compute_instance" "vm_instance" {
  name         = "minecraft-abekoh"
  machine_type = "e2-medium"
  tags         = ["minecraft-server"]

  boot_disk {
    initialize_params {
      image = data.google_compute_image.base_image.self_link
    }
  }

  attached_disk {
    source = google_compute_disk.game_data_disk.name
  }

  network_interface {
    network = google_compute_network.vpc_network.name
    access_config {
    }
  }

  scheduling {
    preemptible       = true
    automatic_restart = false
  }

  metadata = {
    startup-script-url = "gs://minecraft-abekoh-scripts/setup-minecraft-server.sh"
  }
}