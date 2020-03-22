tfmodule
========

![test](https://github.com/chroju/tfmodule/workflows/test/badge.svg?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/chroju/tfmodule/badge.svg?branch=master)](https://coveralls.io/github/chroju/tfmodule?branch=master)

tfmodule is a CLI tool for managing Terraform modules.

## Usage

tfmodule has sub-commands `template` , `analyze` .

**NOTE: in progress status, so you can use only `template` sub-command.**

### template

`template` parses the Terraform module configuration in the given file path, and outputs a template for the module.

```bash
$ tfmodule template ../modules/hoge
module "hoge" {
  source = "../modules/hoge"

  // variable discription
  // type: number
  instance_counts = 2
  // variable discription
  // type: string
  instance_comment = "" // no default value
}

# You can replace module name in the template with -n or --name option.
# Use the --minimum option, ignore variables with default values.
$ tfmodule template ../modules/hoge -n fuga --minimum
module "fuga" {
  source = "../modules/hoge"

  // variable discription 
  // type: string
  instance_comment = "" // no default value
}
```

### analyze

`analyze` analyzes a Terraform module configuration and outputs the description about its internal structure.

```bash
$ tfmodule analyze "../modules/hoge"
resources:
aws_s3_bucket.s3_bucket
aws_s3_bucket_policy.s3_bucket_policy

outputs:
bucket_arn: aws_s3_bucket.s3_bucket.arn
```

## Author

chroju <chor.chroju@gmail.com>
