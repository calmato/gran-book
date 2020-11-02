resource "google_compute_global_address" "this" {
  count = var.create_global_address ? 1 : 0

  name        = var.global_address_name
  description = var.global_address_description

  ip_version = var.global_address_ip_version
}
