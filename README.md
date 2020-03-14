tfmodule
========

## Usage

### import

```bash
$ tfmodule import "../modules/hoge" -n "hoge"
module "hoge" {
  source = "../modules/hoge"

  // variable discription (type: number)
  instance_counts = 2
  // variable discription (type: string)
  instance_comment = "" // no default value
}

$ tfmodule import "../modules/hoge" -n "hoge" --required-only
module "hoge" {
  source = "../modules/hoge"

  // variable discription (type: string)
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
