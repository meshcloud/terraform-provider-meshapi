build:
	go build

tidy:
	go mod tidy
	go mod vendor

update-provider: build
	mv terraform-provider-meshapi ~/.terraform.d/plugins/registry.terraform.io/hashicorp/meshapi/1.0.0/darwin_amd64/terraform-provider-meshapi

tf-apply:
	terraform init && terraform apply -auto-approve

tf-clean-apply: clean-tf
	terraform init && terraform apply -auto-approve

clean-tf:
	rm -rf .terraform/*
	rm -f .terraform.lock.hcl
	rm -f terraform.tfstate