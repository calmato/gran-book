output "container_cluster" {
  value = google_container_cluster.this
}

output "container_node" {
  value = google_container_node_pool.this
}
