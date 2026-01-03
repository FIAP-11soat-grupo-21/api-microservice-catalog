application_name = "catalog-api"
image_name       = "GHCR_IMAGE_TAG"
image_port       = 8080
app_path_pattern = ["/products*", "/products/*", "/categories*", "/categories/*"]

# =======================================================
# Configurações do ECS Service
# =======================================================
container_environment_variables = {
  GO_ENV : "production"
  API_PORT : "8080"
  API_HOST : "0.0.0.0"
  AWS_REGION : "us-east-2"

  DB_RUN_MIGRATIONS : "true"
  DB_NAME : "postgres"
  DB_PORT : "5432"
  DB_USERNAME : "adminuser"
  AWS_REGION : "us-east-2"
  AWS_S3_BUCKET_NAME : "product-photo-fiap-tech-challenge-catalog"
  AWS_S3_PRESIGN_EXPIRATION : "5m"
  AWS_S3_ENDPOINT : ""
  # API_UPLOAD_URL será preenchida automaticamente pelo Terraform no main-2.tf usando o output do módulo S3
  # Exemplo de valor: https://product-photo-fiap-tech-challenge-catalog.s3.us-east-2.amazonaws.com
}
container_secrets = {}
health_check_path = "/health"
task_role_policy_arns = [
  "arn:aws:iam::aws:policy/AmazonS3FullAccess",
  "arn:aws:iam::aws:policy/AmazonRDSFullAccess",
  "arn:aws:iam::aws:policy/SecretsManagerReadWrite"
]

# =======================================================
# Configurações do API away
# =======================================================
apigw_integration_type       = "HTTP_PROXY"
apigw_integration_method     = "ANY"
apigw_payload_format_version = "1.0"
apigw_connection_type        = "VPC_LINK"

authorization_name = "CognitoAuthorizer"

api_endpoints = {
  get_category = {
    route_key  = "GET /categories/{id}"
    restricted = false
  },
  get_all_categories = {
    route_key  = "GET /categories"
    restricted = false
  },
  put_category = {
    route_key  = "PUT /categories/{id}"
    restricted = false
  },
  post_category = {
    route_key  = "POST /categories"
    restricted = false
  },
  delete_category = {
    route_key  = "DELETE /categories/{id}"
    restricted = false
  },
  get_product = {
    route_key  = "GET /products/{id}"
    restricted = false
  },
  get_all_products = {
    route_key  = "GET /products"
    restricted = false
  },
  put_products = {
    route_key  = "PUT /products/{id}"
    restricted = false
  },
  post_products = {
    route_key  = "POST /products"
    restricted = false
  },
  delete_product = {
    route_key  = "DELETE /products/{id}"
    restricted = false
  },
  patch_product_image = {
    route_key  = "PATCH /products/{id}/images"
    restricted = false
  },
  delete_product_image = {
    route_key  = "DELETE /products/{id}/images/{image_file_name}"
    restricted = false
  },
  get_product_images = {
    route_key  = "GET /products/{id}/images"
    restricted = false
  },
  health = {
    route_key  = "GET /health"
    restricted = false
  },
  v1_health = {
    route_key  = "GET /v1/health"
    restricted = false
  }
}