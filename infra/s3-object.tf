resource "aws_s3_object" "default_product_image" {
  bucket = module.s3_bucket.bucket_name

  key    = "default_product_image.webp"
  source = "../microservice/uploads/default_product_image.webp"

  content_type = "image/webp"

  etag = filemd5("../microservice/uploads/default_product_image.webp")

  depends_on = [
    module.s3_bucket
  ]
}