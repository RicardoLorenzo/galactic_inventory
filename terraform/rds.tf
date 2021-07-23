# Security Group for RDS
resource "aws_security_group" "eks_ricardo_rds_sg" {
  name        = "db"
  description = "security group for webservers"
  vpc_id      = aws_vpc.eks_ricardo_vpc.id


  # Allowing traffic only for MySQL and that too from same VPC only.
  ingress {
    description = "MYSQL"
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = aws_subnet.eks_ricardo_vpc_internal_subnet[*].cidr_block
  }


  # Allowing all outbound traffic
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "rds sg"
  }
}

resource "aws_db_instance" "eks_ricardo_db" {
  depends_on        = [aws_security_group.eks_ricardo_rds_sg]
  allocated_storage = 20
  storage_type      = "gp2"

  # Using MYSQL engine for DB
  engine = "mysql"
  engine_version         = "5.7.30"

  # Defining the Security Group Created
  vpc_security_group_ids = [aws_security_group.eks_ricardo_rds_sg.id]

  instance_class         = "db.t2.micro"

  # Giving Credentials
  name                 = "galacticdb"
  username             = "yoda"
  password             = "theforce"
  parameter_group_name = "default.mysql5.7"

  # Making the RDS/ DB publicly accessible so that end point can be used
  publicly_accessible = true
  # Setting this true so that there will be no problem while destroying the Infrastructure as it won't create snapshot
  skip_final_snapshot = true


  tags = {
    Name = "galacticdb"
  }
}
