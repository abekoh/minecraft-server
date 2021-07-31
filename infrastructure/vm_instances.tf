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
      nat_ip = google_compute_address.static_ip.address
    }
  }

  scheduling {
    preemptible       = true
    automatic_restart = false
  }

  service_account {
    email  = var.vm_serviceaccount_email
    scopes = ["cloud-platform"]
  }

  metadata = {
    startup-script-url  = "gs://minecraft-abekoh-scripts/setup-minecraft-server.sh"
    shutdown-script-url = "gs://minecraft-abekoh-scripts/shutdown-minecraft-server.sh"
    enable-oslogin      = "TRUE"
  }

}


resource "google_compute_instance" "bot_instance" {
  name         = "minecraft-abekoh-discord-bot"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.base_image.self_link
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.name
  }

  service_account {
    email  = var.vm_serviceaccount_email
    scopes = ["cloud-platform"]
  }

  metadata = {
    startup-script-url = "gs://minecraft-abekoh-scripts/setup-discord-bot-server.sh"
    enable-oslogin     = "TRUE"
  }

}