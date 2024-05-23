# Admin Mock Stack [IaC]
![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)

New prove of concepts implementing Terraform tfstate storage to S3 bucket and locked state by using DynamoDB.

The service deployed it was done by AWS console and not by Terraform.
## AWS Services deployed:

* [x] RDS
* [x] Task Definition
## Explanation:

You can update the stack from your localhost because we have set the remote state at S3 bucket.

* S3 bucket: [remotestate-vcetfh20](arn:aws:s3:::remotestate-vcetfh20)
* S3 object location: [/admin-mock-api](https://remotestate-vcetfh20.s3.eu-west-1.amazonaws.com/admin-mock-api/)

```
cd terraform
terraform init
terraform plan
terraform apply
```
#### Requeriments:

* Terraform
* AWS dev account set on your local machine

### Maintainer Disclaimer

Further information you can open an issue to [castrillo-clorian](https://github.com/castrillo-clorian)