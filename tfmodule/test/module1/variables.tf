variable "no_default" {
  type        = string
  description = "no default description"
}

variable "object_type" {
  type = object(
    {
      name  = string,
      count = number,
    }
  )
  default = {
    name  = "default",
    count = 1
  }
  description = "object type description"
}
module "" {
  source = ""

  // no default description
  // type:  string
  no_default = ""

  // object type description
  // type:  object({ name = string, count = number, })
  object_type = {
    name  = "default",
    count = 1
  }

}
