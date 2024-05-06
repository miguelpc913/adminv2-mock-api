terraform {

  ####################
  ### remote state location
  ####################
  backend "s3" {
    bucket = "remotestate-vcetfh20"
    key            = "admin-mock-api/state/terraform.tfstate"
    region         = "eu-west-1"
    encrypt        = true
    kms_key_id     = "alias/terraform-bucket-key"
    dynamodb_table = "terraform-state"
  }
  #######################

  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "5.48.0"
    }
  }
  required_version = ">= 1.7.0"
}

provider "aws" {
  region = var.aws_region
  default_tags {
    tags = {
      Project = "AdminMockApi project"
      Folder  = "terraform"
      Version = "1.0.1"
    }
  }
}
