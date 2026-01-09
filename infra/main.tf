module "catalog_api" {
  source     = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/ECS-Service?ref=main"
  depends_on = [aws_lb_listener.listener]

  cluster_id            = data.terraform_remote_state.infra.outputs.ecs_cluster_id
  ecs_security_group_id = data.terraform_remote_state.infra.outputs.ecs_security_group_id

  cloudwatch_log_group     = data.terraform_remote_state.infra.outputs.ecs_cloudwatch_log_group
  ecs_container_image      = var.image_name
  ecs_container_name       = var.application_name
  ecs_container_port       = var.image_port
  ecs_service_name         = var.application_name
  ecs_desired_count        = var.desired_count
  registry_credentials_arn = data.terraform_remote_state.infra.outputs.ecr_registry_credentials_arn

  ecs_container_environment_variables = merge(
    var.container_environment_variables,
    {
      DB_HOST        = data.terraform_remote_state.infra.outputs.rds_address,
      DB_USERNAME    = data.terraform_remote_state.infra.outputs.rds_postgres_db_username,
      API_UPLOAD_URL = module.s3_bucket.bucket_regional_domain_name != "" ? "https://${module.s3_bucket.bucket_regional_domain_name}" : ""
    }
  )

  ecs_container_secrets = merge(var.container_secrets,
    {
      DB_PASSWORD : data.terraform_remote_state.infra.outputs.rds_secret_arn
    }
  )

  private_subnet_ids      = data.terraform_remote_state.infra.outputs.private_subnet_ids
  task_execution_role_arn = data.terraform_remote_state.infra.outputs.ecs_task_execution_role_arn
  task_role_policy_arns   = var.task_role_policy_arns
  alb_target_group_arn    = aws_alb_target_group.target_group.arn
  alb_security_group_id   = data.terraform_remote_state.infra.outputs.alb_security_group_id

  project_common_tags = data.terraform_remote_state.infra.outputs.project_common_tags
}

module "GetCatalogAPIRoute" {
  source     = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/API-Gateway-Routes?ref=main"
  depends_on = [module.catalog_api, ]

  api_id       = data.terraform_remote_state.infra.outputs.api_gateway_id
  alb_proxy_id = aws_apigatewayv2_integration.alb_proxy.id

  endpoints = {
    get_category = {
      route_key           = "GET /categories/{id}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    get_all_categories = {
      route_key           = "GET /categories"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    create_category = {
      route_key           = "POST /categories"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    update_category = {
      route_key           = "PUT /categories/{id}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    delete_category = {
      route_key           = "DELETE /categories/{id}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    get_product = {
      route_key           = "GET /products/{id}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    get_all_product = {
      route_key           = "GET /products"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    create_product = {
      route_key           = "POST /products"
      restricted          = true
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    update_product = {
      route_key           = "PUT /products/{id}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    delete_product = {
      route_key           = "DELETE /products/{id}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    patch_product_image = {
      route_key           = "PATCH /products/{id}/images"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    delete_product_image = {
      route_key           = "DELETE /products/{id}/images/{image_file_name}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    get_product_images = {
      route_key           = "GET /products/{id}/images"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    health = {
      route_key           = "GET /health"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    # v1_health = {
    #   route_key           = "GET /v1/health"
    #   restricted          = false
    #   auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    # },
  }
}