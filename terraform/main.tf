terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }

  required_version = ">= 1.3.0"
}

provider "google" {
  credentials = file("../terraform-key.json")
  project     = var.project_id
  region      = var.region
}
