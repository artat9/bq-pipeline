# variable project_name {
#   type = string
#   description = "Project ID of GCP"
#   default = "os-tmp"
# }

# variable dataset_id {
#   type        = string
#   description = "Dataset ID of BigQuery"
#   default     = "sample-dataset-id"
# }

# variable env {
#   type = string
#   description = "Envrionment"
#   default = "DEV"
# }

# variable project_name {
#   type = string
#   description = "Project ID of GCP"
#   default = "os-tmp"
# }


## global variables
variable "project" {}
variable "env" {}
variable "region" {
  default = "us-central1"
}
variable "dataset_location" {
  default = "us-central1"
}
variable "dataset_id"{}
