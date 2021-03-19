# remote

provider "aws" {
  region = "us-east-1"
}

module "example_bucket_remote" {
  source      = "./module"
  bucket_name = "example-bucket-remote"
}
