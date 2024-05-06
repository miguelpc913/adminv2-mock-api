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

variable "task_name" {
  type        = string
  description = "task name"
}

variable "ecr_image" {
  type        = string
  description = "repository image"
}
