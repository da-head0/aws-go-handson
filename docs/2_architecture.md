```
flowchart TD
    A[soure code]|serverless.yml| B(serverless deploy)
    B --> C{CloudFormation}
    C -->|lambda function| D[endpoint] --> E[User]
```