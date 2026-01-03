resource "aws_s3_object" "default_product_image" {
  bucket = aws_s3_bucket.this.id

  key    = "default_product_image.webp"
  source = "../microservice/uploads/default_product_image.webp"

  content_type = "image/webp"

  etag = filemd5("../microservice/uploads/default_product_image.webp")

  depends_on = [
    aws_s3_bucket.this
  ]
}