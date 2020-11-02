provider "google" {
  project = var.project_id
  region  = "asia-northeast1"
}

module "this" {
  source = "./../../modules/gke"

  location = "asia-northeast1-a"

  #################################################
  # GKE Cluster
  #################################################
  gke_cluster_name        = "xxxxxx-cluster"
  gke_cluster_description = "xxxxxx application cluster for staging"

  gke_cluster_min_master_version = "1.15.11-gke.13"

  #################################################
  # GKE Node
  #################################################
  gke_node_configs = [
    {
      name         = "xxxxxx-node"
      count        = 1
      preemptible  = false
      machine_type = "g1-small"
      disk_type    = "pd-standard"
      disk_size_gb = 10
    },
    {
      name         = "xxxxxx-spot-node"
      count        = 1
      preemptible  = true
      machine_type = "g1-small"
      disk_type    = "pd-standard"
      disk_size_gb = 10
    },
  ]

  #################################################
  # GCE Global Address
  #################################################
  create_global_address = true

  global_address_name = "xxxxxx-ip-address"
}
