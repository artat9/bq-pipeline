.PHONY: setup build deploy
TARGET = v1dev

setup:
	export GO111MODULE=on
	go mod vendor
	go get github.com/aws/aws-lambda-go/cmd/build-lambda-zip
  git clone https://github.com/bridges-inc/kaleido-core.git
	cd kaleido-core && yarn && npx hardhat compile

abi:
	abigen --abi kaleido-core/abi/contracts/AdManager.sol/AdManager.json --pkg contracts --out pkg/contracts/adm/admanager.go	--alias bid=bid1, _bid=bid2
	abigen --abi kaleido-core/abi/contracts/token/DistributionRight.sol/DistributionRight.json --pkg contracts --out pkg/contracts/right/right.go
	abigen --abi kaleido-core/abi/contracts/peripheries/MediaRegistry.sol/MediaRegistry.json --pkg contracts --out pkg/contracts/reg/mediaregistry.go	--alias bid=bid1, _bid=bid2
	abigen --abi kaleido-core/abi/contracts/interfaces/IAdManager.sol/IAdManager.json --pkg contracts --out pkg/contracts/iadmanager/iadmanager.go	--alias bid=bid1, _bid=bid2


build:
	export GO111MODULE=on
	for module_dir in $$(ls cmd | grep -v arweave); do\
	  echo  "building start... $${module_dir}";\
	  	cd cmd/$${module_dir};\
			pwd;\
			mkdir -p bin;\
			env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main main.go || exit 1;\
			cp -R assets . ;\
			rm bin/main.zip ;\
			zip bin/main.zip bin/main assets/arweave/arweave.json;\
			cd ../..;\
			echo  "building finished. $${module_dir}";\
	done
deploy:
	cdk deploy -c target=v1dev --all --require-approval never

deploy_env:
	cdk deploy -c target=$${TARGET} --all --require-approval never

deploy_gcp:
  cd lib/terraform; terraform apply -var-file="dev.tfvars" -state=./local/terraform.tfstate -auto-approve; cd -
