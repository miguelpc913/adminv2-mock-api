
# resource "random_string" "suffix" {
#   length  = 8
#   special = false
#   lower   = true
#   upper   = false
# }


resource "aws_ecs_task_definition" "task_keycloak" {
  family = "admin-mock-api"
  container_definitions = jsonencode([
    {
      name      = var.task_name
      container_name = var.task_name
      image     = var.ecr_image
      cpu       = 1
      memory    = 1024
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
    }
  ])

}


