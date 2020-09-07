# AWS Custom Resources

A Lambda for creating custom resources in CloudFormation.

## Supported Custom Resource Types

* `Custom::ApiKeyValue`: Store value of API Key in a secure SSM parameter.
```yaml
  DevApiKeyValue:
    Type: Custom::ApiKeyValue
    Properties:
      ServiceToken: ${cf:custom-resources-dev.ServiceToken}
      ApiKeyId: !Ref DevApiKey
      SsmParamPath: /api/keys/${self:provider.stage}
      Overwrite: true # default is true
```

## Building and deployment

Install [serverless](https://www.serverless.com/framework/docs/getting-started/) and `make` (`brew install make` on MacOS) and then run:
```shell
make clean build test deploy
```