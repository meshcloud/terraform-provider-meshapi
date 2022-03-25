# BUILD
go build

# MOVE
mv terraform-provider-meshapi ~/.terraform.d/plugins/registry.terraform.io/hashicorp/meshapi/1.0.0/darwin_amd64/terraform-provider-meshapi

# CLEAN
rm -rf .terraform/*
rm -f .terraform.lock.hcl
rm -f terraform.tfstate