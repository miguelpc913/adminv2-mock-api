variable "aws_region" {
  type        = string
  description = "variable of the region selected"
}

# DAtabase 

variable "db_name" {
  type        = string
  description = "db name"
}

variable "key_mag" {
  type        = string
  description = "key management for RDS"
}

variable "db_pass" {
  type        = string
  description = "pass for RDS"
}

# task definition ECS

variable "ecr_image" {
  type        = string
  description = "repository image"
}

variable "health_check" {
  type = map(string)
  default = {
    "timeout"             = "5"
    "interval"            = "20"
    "path"                = "/admin/manage/health"
    "port"                = "8080"
    "unhealthy_threshold" = "2"
    "healthy_threshold"   = "4"
  }
}