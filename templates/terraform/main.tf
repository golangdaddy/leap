variable "project_id" {
  description = "The GCP project ID"
  default     = "{{.Config.ProjectID}}"
}

variable "project_name" {
  description = "The project name"
  default     = "{{.Config.ProjectName}}"
}

variable "project_region" {
  description = "The GCP region for resources"
  default     = "{{.Config.ProjectRegion}}"
}

# Define the Google provider
provider "google" {
  project = var.project_id
  region  = var.project_region
}

# App Engine Application - needed for Firestore
resource "google_app_engine_application" "app" {
  project     = var.project_id
  location_id = var.project_region

  lifecycle {
    create_before_destroy = true
  }
}

# Firestore Database
resource "google_firestore_database" "firestore" {
  project     = var.project_id
  name        = var.project_name
  location_id = var.project_region
  type        = "FIRESTORE_NATIVE"

  lifecycle {
    create_before_destroy = true
  }
}

{{range .Objects}}
# Firestore Index for "{{.Name}}" collection
resource "google_firestore_index" "{{lowercase .Name}}_index" {
  project    = var.project_id
  database   = google_firestore_database.firestore.name
  collection = "{{lowercase .Name}}"

  fields {
    field_path = "Meta.Moderation.Admins"
    order      = "ASCENDING"
  }
  fields {
    field_path = "Meta.Modified"
    order      = "DESCENDING"
  }
  fields {
    field_path = "__name__"
    order      = "DESCENDING"
  }

  query_scope = "COLLECTION"

  lifecycle {
    create_before_destroy = true
  }
}
{{end}}

# Google Cloud Storage Bucket
resource "google_storage_bucket" "bucket" {
  name          = var.project_name
  location      = var.project_region
  force_destroy = true

  lifecycle {
    create_before_destroy = true
  }
}

# Google Pub/Sub Topic
resource "google_pubsub_topic" "topic" {
  name = var.project_name
  labels = {
    project = var.project_name
  }

  lifecycle {
    create_before_destroy = true
  }
}
