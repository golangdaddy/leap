provider "google" {
  project = "{{.PROJECT_ID}}"
  region  = "{{.REGION}}"
}

resource "google_storage_bucket" "bucket1" {
  name          = "bucket1-name"
  location      = "US"
}

resource "google_storage_bucket" "bucket2" {
  name          = "bucket2-name"
  location      = "US"
}
