resource "google_bigquery_dataset" "sample" {
  dataset_id  = var.dataset_id
  description = "sample dataset"
  location    = var.dataset_location
  labels = {
    env = var.env
  }
}

resource "google_bigquery_table" "sample" {
  dataset_id = google_bigquery_dataset.sample.dataset_id
  table_id   = "sample_terraform_user_table"
  schema     = "${file("bigquery_schema.json")}"
  labels = {
    env = var.env
  }
  time_partitioning {
    type  = "DAY"
  }
}


# module "bigquery" {
#   source  = "terraform-google-modules/bigquery/google"
#   version = "~> 4.4"

#   dataset_id                  = var.dataset_id
#   dataset_name                = "sample"
#   description                 = "sample dataset"
#   project_id                  = var.project
#   location                    = var.dataset_location

#   tables = [
#   {
#     table_id           = "User",
#     schema             =  "${file("bigquery_schema.json")}",
#     time_partitioning  = {
#       type                     = "DAY",
#       field                    = null,
#       require_partition_filter = false,
#       expiration_ms            = null,
#     },
#     range_partitioning = null,
#     expiration_time = null,
#     clustering         = [],
#     labels          = {
#       env      = var.env
#     },
#   },
#   ]
# }
