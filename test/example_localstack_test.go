package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestBucket(t *testing.T) {
	var terraformOpts = &terraform.Options{
			TerraformDir: "./example",
	}
    defer terraform.Destroy(t, terraformOpts)
	terraform.InitAndApply(t, terraformOpts)
    bucketUrl := terraform.Output(t, terraformOpts, "localstack_bucket_url")
    assert.Equal(t, "example-localstack-bucket.s3.amazonaws.com", bucketUrl)
}
