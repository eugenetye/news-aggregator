resource "google_container_cluster" "news-cluster" {
  name     = "news-cluster"
  location = var.region

  enable_autopilot = true

  networking_mode = "VPC_NATIVE"
  network    = google_compute_network.vpc.name
  subnetwork = google_compute_subnetwork.subnet.name
}
