module "db" {
  source = "terraform-aws-modules/rds/aws"

  identifier = "adminmockapi-db"

  engine            = "mysql"
  engine_version    = "8.0.35"
  instance_class    = "db.t3.small"
  allocated_storage = 20

  db_name  = var.db_name
  username = "admin"
  # password = var.db_pass
  manage_master_user_password = true
  port     = "3306"

  iam_database_authentication_enabled = true

  vpc_security_group_ids = ["sg-90253ff2"]

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  # Enhanced Monitoring - see example for details on how to create the role
  # by yourself, in case you don't want to create it automatically
  monitoring_interval    = "30"
  monitoring_role_name   = "RDSMonitoringRole-admin"
  create_monitoring_role = true

  tags = {
    Owner       = "admin"
    Environment = "dev"
  }

  # DB subnet group
  create_db_subnet_group = false
  subnet_ids             = ["subnet-20e9b366", "subnet-76988114", "subnet-3cd1eb48", "subnet-0c0bf5631cf27e12f"]

  # DB parameter group
  family = "mysql8.0"

  # DB option group
  major_engine_version = "8.0"

  # Database Deletion Protection
  deletion_protection = false

  kms_key_id = var.key_mag

  # performance_insights_enabled          = true
  # performance_insights_retention_period = 7


  parameters = [
    {
      name  = "character_set_client"
      value = "utf8mb4"
    },
    {
      name  = "character_set_server"
      value = "utf8mb4"
    }
  ]

  options = [
    {
      option_name = "MARIADB_AUDIT_PLUGIN"

      option_settings = [
        {
          name  = "SERVER_AUDIT_EVENTS"
          value = "CONNECT"
        },
        {
          name  = "SERVER_AUDIT_FILE_ROTATIONS"
          value = "37"
        },
      ]
    },
  ]
}