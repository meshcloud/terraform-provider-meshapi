# BUILD
go build

# MOVE
mv terraform-provider-meshstack ~/.terraform.d/plugins/registry.terraform.io/hashicorp/meshstack/1.0.0/darwin_amd64/terraform-provider-meshstack

# CLEAN
rm -rf .terraform/*
rm -f .terraform.lock.hcl
rm -f terraform.tfstate