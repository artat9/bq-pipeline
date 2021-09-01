.PHONY: setup build deploy
TARGET = dev

setup:
	export GO111MODULE=on
	go mod vendor
	go get github.com/aws/aws-lambda-go/cmd/build-lambda-zip

abi:
	abigen --abi aurora-core/abi/contracts/PostManager.sol/PostManager.json --pkg contracts --out lib/functions/lib/contracts/postmanager.go	

build: 
	export GO111MODULE=on
	for module_dir in $$(ls lib/functions | grep -v lib| grep -v voucher|grep -v assets); do\
	  echo  "building start... $${module_dir}";\
	  	cd lib/functions/$${module_dir};\
			pwd;\
			mkdir -p bin;\
			env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main main.go || exit 1;\
			cp -R ../assets . ;\
			rm bin/main.zip ;\
			zip bin/main.zip bin/main assets/gltf/normal/* assets/gltf/ultrarare/* assets/arweave/arweave.json;\
			cd ../../..;\
			echo  "building finished. $${module_dir}";\
	done
deploy: 
	cdk deploy -c target=dev --all --require-approval never

deploy_env: 
	cdk deploy -c target=$${TARGET} --all --require-approval never