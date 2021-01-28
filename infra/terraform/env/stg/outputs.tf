output "gke_cluster" {
  value = module.gke.container_cluster
}

output "gke_node" {
  value = module.gke.container_node
}
