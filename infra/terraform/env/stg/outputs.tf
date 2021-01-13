output "gke_cluster" {
  value = module.this.container_cluster
}

output "gke_node" {
  value = module.this.container_node
}
