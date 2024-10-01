terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 3.5"
    }
  }

  required_version = ">= 0.12"
}

provider "google" {
  credentials = file("<YOUR-CREDENTIALS-FILE>.json")
  project     = var.project_id
  region      = var.region
}
