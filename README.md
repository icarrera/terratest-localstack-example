## terratest localstack example
PoC for running terratest against localstack

### why
- spinning up AWS infrastructure for all tests takes really really long

### requirements 
- docker
- terraform 

### start localstack in docker
`./localstack.sh`

### run terratest against localstack
`go test test/localstack_example_test.go`

### resources & inspo
- https://terratest.gruntwork.io/
- https://github.com/localstack/localstack
- https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/custom-service-endpoints#localstack
- https://spin.atomicobject.com/2020/02/03/localstack-terraform-circleci/ 
- https://anthony-f-tannous.medium.com/set-up-terraform-cloud-to-work-with-aws-localstack-473d4175596f  