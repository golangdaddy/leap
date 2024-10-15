variable "project_id" {
  description = "The GCP project ID"
  default     = "npg-generic"
}

variable "project_name" {
  description = "The project name"
  default     = "alexfirstproject"
}

variable "region" {
  description = "The GCP region for resources"
  default     = "europe-west2"  # Default to the London region
}

# Define the Google provider
provider "google" {
  project = var.project_id
  region  = var.region
}

# App Engine Application - needed for Firestore
resource "google_app_engine_application" "app" {
  project     = var.project_id
  location_id = var.region
}

# Firestore Database
resource "google_firestore_database" "firestore" {
  project     = var.project_id
  location_id = var.region
  type        = "FIRESTORE_NATIVE"
}

{{range .Objects}}
# Firestore Index for "{{.ClassName}}" collection
resource "google_firestore_index" "{{.Name}}_index" {
  project  = var.project_id
  database = google_firestore_database.firestore.name
  collection = "{{lowercase .Name}}"

  fields {
    field_path = "Meta.Moderation.Admins"
    order      = "ARRAY"
  }
  fields {
    field_path = "Meta.Modified"
    order      = "DESCENDING"
  }
  fields {
    field_path = "__name__"
    order      = "DESCENDING"
  }

  # Specifies that this is a composite index for sorting
  query_scope = "COLLECTION"
}
{{end}}

# Google Cloud Storage Bucket
resource "google_storage_bucket" "bucket" {
  name          = var.project_name
  location      = var.region
  force_destroy = true
}

# Google Pub/Sub Topic
resource "google_pubsub_topic" "topic" {
  name = var.project_name
  labels = {
    project = var.project_name
  }
}
