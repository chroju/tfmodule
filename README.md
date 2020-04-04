tfmodule
========

![test](https://github.com/chroju/tfmodule/workflows/test/badge.svg?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/chroju/tfmodule/badge.svg?branch=master)](https://coveralls.io/github/chroju/tfmodule?branch=master)

tfmodule is a CLI tool for managing Terraform modules.

## Install

### Download binary

Download the latest binary from [here](https://github.com/chroju/tfmodule/releases) and put it in your `$PATH` directory.

### go get

If you have set up Go environment, you can also install `tfmodule` with `go get` command.

```
$ go get github.com/chroju/tfmodule
```

## Usage

tfmodule has sub-commands `template` , `analyze` .

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

output "hoge_instance_id" {
  value = module.hoge.instance_id
}
```

You can replace module name in the template with `-n` or `--name` option.

```
$ tfmodule template ../modules/hoge -n fuga
module "fuga" {
  source = "../modules/hoge"

  // variable discription 
  // type: string
  instance_comment = "" // no default value
}

output "fuga_instance_id" {
  value = module.fuga.instance_id
}
```

Use the `--no-default` option, ignore variables with default values, and `--no-outputs` option, ignore outputs.

```
$ tfmodule template ../modules/hoge --no-defaults --no-outputs
module "fuga" {
  source = "../modules/hoge"

  // variable discription 
  // type: string
  instance_comment = "" // no default value
}
```

`--minimum` has the same meaning as adding both `--no-outputs` and `--no-defaults` .

```
$ tfmodule template ../modules/hoge --minimum
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

## TODO

- [ ] Analyze the modules on the Internet as well.

## Author

chroju <chor.chroju@gmail.com>
