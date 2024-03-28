provider "google" {
  project = "{{.ProjectName}}"
  region  = "{{.ProjectRegion}}"
}

resource "google_storage_bucket" "primarybucket" {
  name          = "{{.ProjectName}}-uploads"
  location      = "EU"
}
