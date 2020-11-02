#################################################
# Common
#################################################
variable "location" {
  default = "asia-northeast1-a"
}

variable "tags" {
  default = {}
}

#################################################
# GKE Cluster
#################################################
variable "gke_cluster_name" {
  description = "GKE クラスタ名"
  default     = ""
}

variable "gke_cluster_description" {
  description = "GKE クラスタ説明"
  default     = ""
}

variable "gke_cluster_min_master_version" {
  description = "GKE クラスタ最低バージョン"
  default     = "1.14.10-gke.36"
}

variable "gke_cluster_ipv4_cidr" {
  description = "GKE PodのCIDR"
  default     = null
}

#################################################
# GKE Node
#################################################
variable "gke_node_configs" {
  description = "GKE ノード設定"
  type = list(object({
    name         = string # ノード名
    count        = number # ノード数
    preemptible  = bool   # プリエンプティブの利用
    machine_type = string # マシンタイプ e.g.) n1-standard-1, etc..
    disk_type    = string # ディスクタイプ e.g.) pd-standard or pd-ssd
    disk_size_gb = number # ディスクサイズ[GB] min: 10GB
  }))
}

#################################################
# GCE Global Address
#################################################
variable "create_global_address" {
  description = "グローバルアドレスの作成"
  type        = bool
  default     = false
}

variable "global_address_name" {
  description = "グローバルアドレス名"
  default     = ""
}

variable "global_address_description" {
  description = "グローバルアドレス説明"
  default     = ""
}

variable "global_address_ip_version" {
  description = "グローバルアドレスIPアドレスバージョン"
  default     = "IPV4"
}
