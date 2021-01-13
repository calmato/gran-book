##################################################
# Cloud SQL - Instance
##################################################
resource "google_sql_database_instance" "this" {
  name                 = var.sql_instance_name
  master_instance_name = var.sql_instance_name
  region               = var.location

  database_version = var.sql_instance_database_version
  root_password    = var.sql_instance_root_password

  deletion_protection = false

  settings {
    tier              = var.sql_instance_type
    availability_type = var.sql_availability_type

    disk_type       = var.sql_disk_type
    disk_autoresize = var.sql_disk_autoresize
    disk_size       = var.sql_disk_size

    ip_configuration {
      ipv4_enabled = var.sql_ipv4_enabled
      private_network = var.sql_ipv4_enabled ? var.sql_private_network : null
    }

    backup_configuration {
      enabled    = var.sql_backup_enabled
      start_time = var.sql_backup_enabled ? var.sql_backup_start_time : null
    }
  }
}
