module "s3_bucket" {
  source = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/S3?ref=main"

  bucket_name         = "product-photo-fiap-tech-challenge-catalog"
  enable_versioning   = true
  enable_encryption   = true
  project_common_tags  = { Project = "catalog" }
}