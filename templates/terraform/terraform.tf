provider "google" {
  project = "{{.ProjectName}}"
  region  = "{{.ProjectRegion}}"
}

resource "google_storage_bucket" "primarybucket" {
  name          = "{{.ProjectName}}-uploads"
  location      = "EU"
}

resource "google_secret_manager_secret" "my_secret" {
  provider = google-beta
  secret_id = "my-secret"
}

resource "google_secret_manager_secret_version" "my_secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.my_secret.id
  secret_data = "my-secret-value"
}

resource "google_compute_instance" "example" {
  ...
  metadata_startup_script = google_secret_manager_secret_version.my_secret_version.secret_data
  ...
}