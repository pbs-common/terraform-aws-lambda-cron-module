module "lambda_cron" {
  source = "../.."

  handler  = "bootstrap"
  filename = "../artifacts/handler.zip"
  runtime  = "provided.al2023"

  environment  = var.environment
  product      = var.product
  repo         = var.repo
  owner        = var.owner
  organization = var.organization
}
