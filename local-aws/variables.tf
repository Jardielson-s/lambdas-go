variable "region" {
  type        = string
  default     = "us-east-1"
  description = "Here is aws region"
}
variable "bucket_name" {
  type        = string
  description = "Here is bucket name"
}

variable "queue_name" {
  type        = string
  description = "Here is queue name"
}

variable "env" {
  type        = string
  description = "Here is env"
}


variable "lambdas" {

  type = set(string)

  default = ["process_csv"]
}
