GO_OS := `go env GOOS`
GO_ARCH := `go env GOARCH`

build: tidy
	go build

tidy:
	go mod tidy
	go mod vendor

update-provider: build
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/hashicorp/meshapi/1.0.0/{{GO_OS}}_{{GO_ARCH}}/; mv terraform-provider-meshapi $_

tf-apply:
	terraform init && terraform apply -auto-approve

tf-clean-apply: clean-tf
	terraform init && terraform apply -auto-approve

clean-tf:
	rm -rf .terraform/*
	rm -f .terraform.lock.hcl
	rm -f terraform.tfstate