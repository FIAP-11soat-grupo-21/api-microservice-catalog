# module "GetCatalogAPIRoute" {
#   source     = "git::https://github.com/FIAP-11soat-grupo-21/infra-core.git//modules/API-Gateway-Routes?ref=main"
#   depends_on = [module.catalog_api]

#   api_id       = data.terraform_remote_state.infra.outputs.api_gateway_id
#   alb_proxy_id = aws_apigatewayv2_integration.alb_proxy.id

#   endpoints = var.api_endpoints
# }

