variable "bucket_name" {
  default = "example-bucket"
}

resource "aws_s3_bucket" "example_bucket" {
  bucket = var.bucket_name
  acl = "private"
}

output "bucket_url" {
  value = aws_s3_bucket.example_bucket.bucket_domain_name
}
