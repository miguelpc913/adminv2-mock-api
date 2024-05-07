data "aws_iam_role" "ecs_role" {
  name = "ecsServiceRole"
}

data "aws_vpc" "dev_vpc" {
  id = "vpc-183c2f7a"
}

data "aws_ecs_cluster" "ecs-qa" {
  cluster_name = "qa"
}

data "aws_lb" "microservices" {
  name = "QA-Microservices"
}

data "aws_ecs_task_definition" "admin_task" {
  task_definition = "admin-mock-api"
}