# Create a Firestore database in Native mode
resource "google_firestore_database" "default" {
  provider = google
  project  = var.project_id
  type     = "CLOUD_FIRESTORE_NATIVE"
  location = var.region
}

# Create a Google Cloud Storage bucket
resource "google_storage_bucket" "default" {
  name     = "${var.project_id}-bucket"
  location = var.region

  storage_class = "STANDARD"
  lifecycle_rule {
    condition {
      age = 365
    }
    action {
      type = "Delete"
    }
  }
}
