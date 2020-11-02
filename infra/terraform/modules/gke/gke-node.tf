#################################################
# GKE Node
#################################################
resource "google_container_node_pool" "this" {
  for_each = local.gke_node_configs

  name = each.key

  cluster  = google_container_cluster.this.name
  location = var.location

  node_count = each.value.count

  management {
    auto_repair = true
  }

  node_config {
    preemptible  = each.value.preemptible
    machine_type = each.value.machine_type

    disk_type    = each.value.disk_type
    disk_size_gb = each.value.disk_size_gb

    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}
