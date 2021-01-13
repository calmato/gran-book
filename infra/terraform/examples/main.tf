provider "google" {
  project = var.project_id
  region  = local.region
}

locals {
  region   = "asia-northeast1"
  location = "asia-northeast1-a"
}

module "gke" {
  source = "./../../modules/gke"

  location = local.location

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

module "mysql" {
  source = "./../../modules/sql"

  region = local.region

  #################################################
  # Cloud SQL - Instance
  #################################################
  sql_instance_name          = "xxxxxx-stg-mysql"
  sql_instance_root_password = var.sql_instance_root_password

  sql_instance_database_version = "MYSQL_8_0"
  sql_instance_type             = "db-f1-micro"

  #################################################
  # Cloud SQL - Network
  #################################################
  sql_availability_type = "ZONAL" # ZONAL / REGIONAL
  sql_ipv4_enabled      = false
  sql_private_network   = ""

  #################################################
  # Cloud SQL - Network
  #################################################
  sql_disk_type       = "PD_SSD" # PD_SSD / PD_HDD
  sql_disk_autoresize = false
  sql_disk_size       = 10

  #################################################
  # Cloud SQL - Network
  #################################################
  sql_backup_enabled    = false
  sql_backup_start_time = "" # format: HH:mm
}

