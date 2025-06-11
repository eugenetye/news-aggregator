output "cluster_name" {
  value = google_container_cluster.news-cluster.name
}

output "location" {
  value = var.region
}
