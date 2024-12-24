module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "19.13.1"

  cluster_name    = local.cluster_name
  cluster_version = "1.31"

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets
  cluster_endpoint_public_access = true

  cluster_encryption_policy_name = "EKSGoMicro"

  eks_managed_node_group_defaults = {
    ami_type = "AL2_x86_64"
  }

  eks_managed_node_groups = {
    one = {
      name = "node-group-1"

      instance_types = ["t3.large"]

      min_size     = 1
      max_size     = 3
      desired_size = 1
    }
  }
}

resource "aws_iam_policy" "cloudwatch-agent-server-policy" {
  name        = "CloudWatchAgentServerPolicy"
  path        = "/"
  description = "IAM policy for CloudWatch Agent on EKS nodes"

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams"
        ],
        "Resource": "arn:aws:logs:*:*:*"
      }
    ]
  })
}

output "cloudwatch_agent_server_policy_arn" {
  value = aws_iam_policy.cloudwatch-agent-server-policy.arn
  description = "CloudWatch Agent Server Policy ARN"
}
