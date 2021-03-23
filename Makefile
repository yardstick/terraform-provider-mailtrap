TF_LOG=ERROR

test: build load init plan apply email



email:
	cd tf-test && \
		terraform output -json | python test.py


init:
	cd tf-test && \
		terraform init -var-file=test.tfvars

plan: 
	cd tf-test && \
		terraform plan -var-file=test.tfvars
apply:
	cd tf-test && \
		TF_LOG=$(TF_LOG) terraform apply -var-file=test.tfvars -auto-approve
		
destroy:
	cd tf-test && \
		TF_LOG=$(TF_LOG) terraform destroy -var-file=test.tfvars -auto-approve

build:
	go mod tidy
	go build .

load:
	cp terraform-provider-mailtrap ~/.terraform.d/plugins/darwin_amd64/terraform-provider-mailtrap

docs:
	go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs