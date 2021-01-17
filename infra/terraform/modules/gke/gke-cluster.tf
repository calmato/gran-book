#################################################
# GKE Cluster
#################################################
resource "google_container_cluster" "this" {
  provider = google-beta

  name        = var.gke_cluster_name
  description = var.gke_cluster_description

  location = var.location

  min_master_version = var.gke_cluster_min_master_version
  cluster_ipv4_cidr  = var.gke_cluster_ipv4_cidr

  remove_default_node_pool = true
  initial_node_count       = 1

  addons_config {
    istio_config {
      disabled = var.gke_cluster_istio_disabled
    }
  }
}
