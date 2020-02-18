# Go microservice with AWS Lambda and Amazon API Gateway

Build a highly scalable microservice with AWS Lambda and Amazon API Gateway using this CloudFormation template.

### Package CloudFormation template

`$ aws cloudformation package --template-file template.yaml --s3-prefix <BUCKET_KEY> --s3-bucket <BUCKET_NAME> --output-template-file packaged-template.yaml`

Packages the local artifacts (local paths) that your AWS CloudFormation template references. The command uploads local artifacts, such as source code for an AWS Lambda function or a Swagger file for an AWS API Gateway REST API, to an S3 bucket. The command returns a copy of your template, replacing references to local artifacts with the S3 location where the command uploaded the artifacts.

### Deploy CloudFormation template

`$ aws cloudformation deploy --template-file packaged-template.yaml --stack-name <STACK_NAME> --capabilities CAPABILITY_NAMED_IAM --parameter-overrides StageName=<STAGE_NAME>`

Deploys the specified AWS CloudFormation template by creating and then executing a change set. The command terminates after AWS CloudFormation executes the change set. If you want to view the change set before AWS CloudFormation executes it, use the --no-execute-changeset flag.