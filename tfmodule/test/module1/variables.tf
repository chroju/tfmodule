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
