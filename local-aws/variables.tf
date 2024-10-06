variable "bucket_name" {
  type        = string
  description = "Here is bucket name"
}

variable "env" {
  type        = string
  description = "Here is env"
}


variable "lambdas" {

  type = set(string)

  default = ["process_csv"]
}
