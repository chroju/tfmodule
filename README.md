tfmodule
========

tfmodule is a CLI tool for management to Terraform modules.

## Usage

tfmodule has sub-commands `template` , `analyze` .

### template

```bash
$ tfmodule template "../modules/hoge"
module "hoge" {
  source = "../modules/hoge"

  // variable discription
  // type: number
  instance_counts = 2

  // variable discription
  // type: string
  instance_comment = "" // no default value
}

# You can replace module name with -n or --name option
# With --minimum option, ignore variables which is given the default value
$ tfmodule template "../modules/hoge" -n "fuga" --minimum
module "hoge" {
  source = "../modules/hoge"

  // variable discription 
  // type: string
  instance_comment = "" // no default value
}
```

## analyze

```bash
$ tfmodule analyze "../modules/hoge"
resources:
aws_s3_bucket.s3_bucket
aws_s3_bucket_policy.s3_bucket_policy
```
