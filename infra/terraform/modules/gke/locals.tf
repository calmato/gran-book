locals {
  gke_node_configs = {
    for c in var.gke_node_configs : c.name => {
      name         = c.name
      count        = c.count
      preemptible  = c.preemptible
      machine_type = c.machine_type
      disk_type    = c.disk_type
      disk_size_gb = c.disk_size_gb
      oauth_scopes = flatten([
        [
          "https://www.googleapis.com/auth/cloud-platform",
        ],
        c.monitoring_enabled ? [
          "https://www.googleapis.com/auth/logging.write",
          "https://www.googleapis.com/auth/monitoring",
        ] : [],
      ])
    }
  }
}
