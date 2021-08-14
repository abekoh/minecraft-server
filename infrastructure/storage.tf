resource "google_storage_bucket" "backup" {
  name          = "minecraft-abekoh-backup"
  location      = "US-CENTRAL1"
  storage_class = "REGIONAL"
  lifecycle_rule {
    condition {
      age = 30
    }
    action {
      type = "Delete"
    }
  }
}