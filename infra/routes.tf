module "api_gateway_health_route" {
  source     = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/API-Gateway-Routes?ref=main"
  depends_on = [module.catalog_api]

  api_id          = data.terraform_remote_state.infra.outputs.api_gateway_id
  gwapi_route_key = "GET /health"
  alb_proxy_id    = aws_apigatewayv2_integration.alb_proxy.id
}


# module "api_gateway_products_routes" {
#   source     = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/API-Gateway-Routes?ref=main"
#   depends_on = [module.catalog_api]
#   api_id          = data.terraform_remote_state.infra.outputs.api_gateway_id
#   # vpc_link_id       = data.terraform_remote_state.infra.outputs.api_gateway_vpc_link_id
#   # alb_listener_arn  = module.ALB.listener_arn
#   gwapi_route_key = "ANY /v1/products/{proxy+}"
#   alb_proxy_id    = aws_apigatewayv2_integration.alb_proxy.id

# }

# module "api_gateway_categories_routes" {
#   source     = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/API-Gateway-Routes?ref=main"
#   depends_on = [module.catalog_api]
#   api_id          = data.terraform_remote_state.infra.outputs.api_gateway_id
#   # vpc_link_id       = data.terraform_remote_state.infra.outputs.api_gateway_vpc_link_id
#   # alb_listener_arn  = module.ALB.listener_arn
#   gwapi_route_key = "ANY /v1/categories/{proxy+}"
#   alb_proxy_id    = aws_apigatewayv2_integration.alb_proxy.id
#}

module "api_gateway_v1_proxy_route" {
  source          = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/API-Gateway-Routes?ref=main"
  depends_on      = [module.catalog_api]
  api_id          = data.terraform_remote_state.infra.outputs.api_gateway_id
  gwapi_route_key = "ANY /v1/{proxy+}"
  alb_proxy_id    = aws_apigatewayv2_integration.alb_proxy.id
}

