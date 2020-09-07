# Golang serverless example

A demo stack with two lambdas (created from serverless.com template).

See usage of custom resource `ApiKeyValue` in [serverless.yml](serverless.yml#65) file:
```yaml
  DevApiKeyValue:
    Type: Custom::ApiKeyValue
    Properties:
      ServiceToken: ${cf:custom-resources-dev.ServiceToken}
      ApiKeyId: !Ref DevApiKey
      SsmParamPath: /api/keys/${self:provider.stage}
```
Above resource reads the value of `DevApiKey` and store it in `/api/keys/dev` SSM secure parameter.

## Endpoints

First endpoint `/dev/hello` is public, second `/dev/world` is private. To invoke it pass `x-api-key` header.


## Building and deployment

Install [serverless](https://www.serverless.com/framework/docs/getting-started/) and `make` (`brew install make` on MacOS) and then run:
```shell
make clean build deploy
```