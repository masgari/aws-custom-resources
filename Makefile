.PHONY: clean build test

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/custom-resource cmd/custom-resource/main.go

test:
	go test -short -timeout 30s -v ./... -coverprofile=/tmp/custom-resource-test-coverage.txt -race	

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy:
	sls deploy -c deployments/serverless.yml --verbose

remove:
	sls remove -c deployments/serverless.yml --verbose
