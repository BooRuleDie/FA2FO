# Interview 1

### Question 1

You have been assigned to design a VPC architecture for a 2-tier application. The application needs to be highly available and scalable. How would you design the VPC architecture ?

### Answer 1

First, I'd create a VPC with 2 public and private subnets. I'd distribute those subnets across multiple availability zones for high availability, and I'd deploy a load balancer in the public subnets. Lastly, I'd deploy my instances in private subnets with an auto scaling group to achieve scalability.

---

### Question 2

Your organization has a VPC with multiple subnets. You want to restrict outbound internet access for resources in one subnet, but allow outbound internet access for resources in another subnet. How would you achieve this ?

### Answer 2

There are multiple ways to achieve that. Firstly, we can modify the route table associated with those subnets. If we want to restrict outbound traffic, we can remove the route that points to the Internet Gateway. Internet Gateway is the default outbound route for subnets, so for the other subnet we don't need any configuration.

Or we can achieve a similar behavior with Network ACLs, they help us control the outbound and inbound traffic to subnets. We can restrict particular subnets' NACL to restrict particular ports, IP addresses etc.

---

### Question 3

You have a VPC with a public subnet and a private subnet. Instances in the private subnet need to access the internet for software updates. How would you allow internet access for instances in the private subnet ?

### Answer 3

NAT Gateway is used for that purpose. I'd deploy a NAT Gateway inside the public subnet so it would have access to the Internet Gateway for outbound communication (in this case software updates). I'd then connect that NAT Gateway with the private subnet so whenever instances in this subnet need to communicate with the outside world, they can use the NAT Gateway.

---

### Question 4

You have launched EC2 instances in your VPC, and you want them to communicate with each other using private IP addresses. What steps would you take to enable this communication ?

### Answer 4

By default, as long as they're in the same VPC and Security Groups of instances don't block the communication, there's no need for additional setup for them to talk with each other. You can verify this with SSH - use the same SSH key for all private EC2 instances and after you've connected to one from a bastion host, try to connect to others from that bastion host. You'll see that you are able to do it without any additional configuration.

---

### Question 5

You want to implement strict network access control for your VPC resources. How would you achieve this?

### Answer 5

For network security, there are two components that can help: Network ACLs and Security Groups. The first one provides subnet-level protection and the other one provides instance-level protection. They act like firewalls, allowing us to restrict inbound and outbound traffic. The main difference is their protection level.

---

### Question 6

Your organization requires an isolated environment within the VPC for running sensitive workloads. How would you set up this isolated environment ?

### Answer 6

After the VPC is created, we can create a private subnet. By default, private subnets are not attached to the Internet Gateway, so no one can access these instances from outside, and these instances can't access any resources outside. If the second statement is a problem, it can be solved with a NAT Gateway/Instance. If the first statement is a problem, we can deploy a bastion/jump host in the public subnet, and then we can connect to those isolated private instances through the bastion host.

---

### Question 7

Your application needs to access AWS services, such as S3 securely within your VPC. How would you achieve this ?

### Answer 7

To solve this requirement, we can use VPC Endpoints, which come in two types: Gateway Endpoints (for S3 and DynamoDB) and Interface Endpoints (for other AWS services). VPC Endpoints allow instances inside the VPC to communicate securely with AWS Services without requiring an Internet Gateway or NAT Gateway. The traffic stays within AWS's network, providing better security and reduced latency. Additionally, there are no data transfer costs for using VPC Endpoints within the same region.

---

### Question 8

What is the difference between NACL and Subnet? Explain with a use case?

### Answer 8

NACL (Network Access Control List) and Subnet are two different but related concepts. A subnet is considered a part of the VPC. A VPC is the dedicated private network in AWS, and a subnet is a part of that dedicated network. It can be either public or private, which means it can have internet access or can be completely isolated.

NACLs provide security for subnets. They restrict outbound and inbound traffic to a particular subnet by using ports, protocols, and IP address rules.

---

### Question 9

What is the difference between IAM users, groups, roles and policies ?

### Answer 9

IAM users are individual identities that represent people or applications that need access to AWS resources. Each user has their own credentials.

IAM groups are collections of IAM users. Groups make it easier to manage permissions - instead of attaching policies to individual users, you can attach them to a group and all users in that group inherit those permissions.

IAM roles are similar to users but are meant to be assumed by AWS services, applications, or users who need temporary access. Unlike users, roles don't have permanent credentials - they provide temporary security credentials.

IAM policies are documents that define permissions - what actions are allowed or denied on what AWS resources. Policies can be attached to users, groups, or roles. They're written in JSON format and specify exactly what level of access is granted.

