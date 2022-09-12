TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=gregorystotts.com
NAMESPACE=com
NAME=insightcloudsec
BINARY=terraform-provider-${NAME}
VERSION=0.3.3
OS_ARCH=darwin_arm64

default: install

build:
	go build -o ${BINARY}

install: build
	touch tests/.terraform && rm -r tests/.terraform
	touch tests/.terraform.lock.hcl && rm -r tests/.terraform.lock.hcl
	touch tests/*.tfstate && rm -r tests/*.tfstate
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4
	
testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
