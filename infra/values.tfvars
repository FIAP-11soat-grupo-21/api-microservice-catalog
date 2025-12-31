application_name = "catalog-api"
image_name       = "GHCR_IMAGE_TAG"
image_port       = 8080

# =======================================================
# Configurações do ECS Service
# =======================================================
container_environment_variables = {
  GO_ENV : "production"
  API_PORT : "8080"
  API_HOST : "0.0.0.0"
  AWS_REGION : "us-east-2"
  API_UPLOAD_URL : "http://localhost:8080/uploads"
  DB_RUN_MIGRATIONS : "true"
  DB_NAME : "postgres"
  DB_PORT : "5432"
  DB_USERNAME : "adminuser"
  AWS_REGION : "us-east-2"
  AWS_S3_BUCKET_NAME : "product-photo-fiap-tech-challenge-catalog"
  AWS_S3_PRESIGN_EXPIRATION : "5m"
  AWS_S3_ENDPOINT : "http://minio:9000"
  # COLLECTOR_ID : "2456291815"
  # EXTERNAL_POS_ID : "tccaixafiapf1"
  # MERCADOPAGO_API_URL : "https://api.mercadopago.com"
  # ACCESS_TOKEN : "valor",
  # GOOGLE_PROJECT_ID : "fiap-tech-challenge",
}
container_secrets = {}
health_check_path = "/health"
task_role_policy_arns = [
  "arn:aws:iam::aws:policy/AmazonS3FullAccess",
  "arn:aws:iam::aws:policy/AmazonRDSFullAccess",
  "arn:aws:iam::aws:policy/SecretsManagerReadWrite"
]

# =======================================================
# Configurações do API Gateaway
# =======================================================
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
  }
}