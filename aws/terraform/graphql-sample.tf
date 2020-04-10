provider "aws" {
  profile    = "default" // ~/.aws/credentials
  region     = "ap-northeast-1"
}

resource "aws_instance" "graphql-sample" {
  ami           = "ami-08857a61019dd36ca" // Todo: amiの自動取得(値を直接指定しない)
  instance_type = "t2.micro"
  key_name = "graphql-sample"
  security_groups = ["web-server", "allow-outgoing", "allow-ssh"]
}
