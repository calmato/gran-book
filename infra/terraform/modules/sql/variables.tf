##################################################
# Common
##################################################
variable "location" {
  default = "asia-northeast1-a"
}

##################################################
# Cloud SQL - Instance
##################################################
variable "sql_instance_name" {
  description = "SQL インスタンス名"
  default     = ""
}

variable "sql_instance_database_version" {
  description = "SQL データベースバージョン"
  default     = "MYSQL_8_0" # e.g.) MYSQL_5_7, MYSQL_8_0, POSTGRES_12, POSTGRES_13
}

variable "sql_instance_root_password" {
  description = "SQL Rootパスワード"
  default     = ""
}

variable "sql_instance_type" {
  description = "SQL インスタンスタイプ"
  default     = "db-f1-micro"
}

variable "sql_availability_type" {
  description = "SQL 可溶性タイプ" # 高可用性: REGIONAL, 単一ゾーン: ZONAL
  default     = "ZONAL"
}

variable "sql_disk_type" {
  description = "SQL ディスクタイプ"
  default     = "PD_SSD" # PD_SSD, PD_HDD
}

variable "sql_disk_autoresize" {
  description = "SQL ディスク容量自動拡張"
  type        = bool
  default     = false
}

variable "sql_disk_size" {
  description = "SQL ディスク容量(GB)"
  type        = number
  default     = 10
}

variable "sql_ipv4_enabled" {
  description = "SQL プライベートIPv4利用"
  type        = bool
  default     = false
}

variable "sql_private_network" {
  description = "SQL 配置VPC"
  default     = "default"
}

variable "sql_backup_enabled" {
  description = "SQL バックアップ有効化"
  type        = bool
  default     = false
}

variable "sql_backup_start_time" {
  description = "SQL バックアップ開始時間"
  default     = "" # format: HH:mm
}
