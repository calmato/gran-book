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
  gke_cluster_name        = "gran-book-stg-cluster"
  gke_cluster_description = "gran-book-stg application cluster for staging"

  gke_cluster_min_master_version = "1.18.12-gke.1201"

  #################################################
  # GKE Node
  #################################################
  gke_node_configs = [
    {
      name         = "gran-book-stg-node"
      count        = 1
      preemptible  = false
      machine_type = "e2-micro"
      disk_type    = "pd-standard"
      disk_size_gb = 8
    },
    {
      name         = "gran-book-stg-spot-node"
      count        = 1
      preemptible  = true
      machine_type = "e2-medium"
      disk_type    = "pd-standard"
      disk_size_gb = 16
    },
  ]

  #################################################
  # GCE Global Address
  #################################################
  create_global_address = true

  global_address_name = "gran-book-stg-ip-address"
}
