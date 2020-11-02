locals {
  gke_node_configs = {
    for c in var.gke_node_configs : c.name => c
  }
}
