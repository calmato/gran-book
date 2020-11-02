.PHONY: init fmt plan apply destroy

create:
	mkdir -p ./env/${ENV}
	cp -r ./examples/ ./env/${ENV}/
	mv ./env/${ENV}/backend.tf.sample ./env/${ENV}/backend.tf
	mv ./env/${ENV}/terraform.tfvars.sample ./env/${ENV}/terraform.tfvars

init:
	cd ./env/${ENV}; terraform init

fmt:
	cd ./env/${ENV}; terraform fmt -check=true
	cd ./env/${ENV}; terraform validate

plan:
	cd ./env/${ENV}; terraform plan -parallelism=120

apply:
	cd ./env/${ENV}; terraform apply -auto-approve -parallelism=120

destroy:
	cd ./env/${ENV}; terraform destroy -auto-approve -parallelism=120
