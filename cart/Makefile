default: build

package-name=cart
service-name=cart-service
artifacts-dir=./.artifacts
bin-dir=./.bin

.PHONY: gen
gen:
	go generate ./...

clean_build:
	echo "Cleaning build directory..." && \
	rm -rf $(bin-dir) && \
    mkdir -p $(bin-dir)

build: clean_build
	go build -o $(bin-dir)/$(service-name)

build_linux: clean_build
	GOOS=linux GOARCH=amd64 go build -o $(bin-dir)/$(service-name)

run:
	go run main.go
	
test:
	mkdir -p $(artifacts-dir) && \
	gotestsum --format short-verbose --junitfile $(artifacts-dir)/junit-report.xml \
	./... github.com/Ralphbaer/hash/common... \
	-coverprofile=$(artifacts-dir)/coverage.out

cover-html:
	go tool cover -html=$(artifacts-dir)/coverage.out -o $(artifacts-dir)/coverage.html

cover: test cover-html
	open $(artifacts-dir)/coverage.html

clean:
	rm -rf $(bin-dir)
	rm -rf $(artifacts-dir)

image: build_linux
	docker build -t hash/$(service-name):v1 .

container:
	docker run --env SERVER_ADDRESS=:9091 -it -p 127.0.0.1:9091:9091 hash/$(service-name):v1 bash

discount-image:
	docker pull hashorg/hash-mock-discount-service

discount-container:
	docker run -p 50051:50051 hashorg/hash-mock-discount-service
	
deploy-image:
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 742261599957.dkr.ecr.us-east-2.amazonaws.com/hash/$(service-name) | bash	
	docker tag hash/$(service-name):v1 742261599957.dkr.ecr.us-east-2.amazonaws.com/hash/$(service-name)
	docker push 742261599957.dkr.ecr.us-east-2.amazonaws.com/hash/$(service-name):latest

doc:
	swagger generate spec -o ./gen/swagger.yaml --scan-models \
	-c "github.com/Ralphbaer/hash/$(package-name)" \
	-c "github.com/Ralphbaer/hash/common" \
