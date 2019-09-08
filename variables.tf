variable "test" {
    default = "test"
    description = "test description"
}

variable "hoge" {
    default = "test2"
    description = "test2 description"
}

resource "aws_instance" "name" {
  ami = "hoge"
}
