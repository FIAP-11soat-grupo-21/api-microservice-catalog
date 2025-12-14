data "terraform_remote_state" "infra" {
  backend = "s3"
  config = {
    bucket = "fiap-tc-terraform-846874-v2"
    key    = "tech-challenge-project/core/terraform.tfstate"
    region = "us-east-2"
  }
}