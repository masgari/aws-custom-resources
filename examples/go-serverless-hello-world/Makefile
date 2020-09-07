.PHONY: build clean deploy

API_GW_ID := set-api-gw-id

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy:
	sls deploy --verbose

remove:
	sls remove --verbose

invoke-hello:
	time -p curl -v https://$(API_GW_ID).execute-api.us-east-1.amazonaws.com/dev/hello

invoke-world:
	time -p curl -v -H "X-Api-Key: $$(aws ssm get-parameter --with-decryption --name /api/keys/dev | jq -r .Parameter.Value)" https://$(API_GW_ID).execute-api.us-east-1.amazonaws.com/dev/world
