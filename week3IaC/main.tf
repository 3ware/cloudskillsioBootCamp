terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "eu-west-2"
}

module "webserver" {
  source = "./modules/ec2"

  servername = "terraformdemo"
  vmsize     = "t3.micro"
}
