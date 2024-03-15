resource "aws_vpc" "primary" {
  cidr_block = "10.0.0.0/16"

  tags = {
    name = "Primary VPC"
  }
}

# Subnets

resource "aws_subnet" "public" {
  count = length(var.availability_zones)

  vpc_id                  = aws_vpc.primary.id
  cidr_block              = "10.0.${count.index}.0/24"
  availability_zone_id    = var.availability_zones[count.index]
  map_public_ip_on_launch = true
}

# Gateways & NAT

resource "aws_internet_gateway" "primary" {
  vpc_id = aws_vpc.primary.id
}

resource "aws_egress_only_internet_gateway" "primary" {
  vpc_id = aws_vpc.primary.id
}

# Routing

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.primary.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.primary.id
  }

  route {
    ipv6_cidr_block        = "::/0"
    egress_only_gateway_id = aws_egress_only_internet_gateway.primary.id
  }

  tags = {
    Name = "Public Route Table"
  }
}

resource "aws_route_table_association" "public" {
  count = length(var.availability_zones)

  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}


resource "aws_security_group" "primary" {
  name   = "primary-sg"
  vpc_id = aws_vpc.primary.id

  ingress {
    protocol    = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "Primary"
  }
}
