resource "google_compute_network" "vpc_network" {
  name = "minecraft-network"
}

resource "google_compute_address" "static_ip" {
  name = "minecraft-server-static-ip"
}

resource "google_compute_firewall" "allow_ssh" {
  name    = "allow-ssh"
  network = google_compute_network.vpc_network.name
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }
  priority = 9000
}

resource "google_compute_firewall" "allow_icmp" {
  name    = "allow-icmp"
  network = google_compute_network.vpc_network.name
  allow {
    protocol = "icmp"
  }
  priority = 8000
}

resource "google_compute_firewall" "allow_minecraft" {
  name    = "allow-minecraft"
  network = google_compute_network.vpc_network.name
  allow {
    protocol = "tcp"
    ports    = ["25565"]
  }
  target_tags = ["minecraft-server"]
  priority    = 1000
}

resource "google_compute_firewall" "allow_dynmap" {
  name    = "allow-minecraft"
  network = google_compute_network.vpc_network.name
  allow {
    protocol = "tcp"
    ports    = ["8123"]
  }
  target_tags = ["minecraft-server"]
  priority    = 2000
}
