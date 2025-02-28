# IAM (Identity and Access Management)

<img src="../img/IAM.png" alt="IAM Icon" width="256" height="256"/>

As you can guess from the name, this service is all about authentication and authorization. There are 4 important topics in IAM:

1. Users
2. Policies
3. Groups
4. Roles

### Users

Users are self-explanatory; you create accounts for people who need to authenticate themselves to AWS to perform particular activities.

### Policies

Policies are configurations that define what a particular user or group is allowed to do within AWS services. By using policies, you can grant permissions to particular users or restrict some users for security purposes.

There are two types of policies: AWS managed ones and custom ones. If you just need some basic configurations like `S3 Full Access` or `S3 Read Only Access`, you can use managed ones. However, if you need a more complex permission set, you might want to create a custom policy and use it for the service.

### Groups 

Handling too many users is cumbersome as you'd need to assign policies manually one by one for each user. Groups are useful in those cases. Instead of creating a user and assigning a policy individually, you can define a group like developer, devops, security, dbadmin... and assign a policy to the group. Thus, you don't need to assign each user to a policy but just assign a group to the user, and they'll have the intended permissions easily.

### Roles

Roles are not for actual users but mostly for services. If a program/service needs temporary access to an AWS service to perform a specific action, roles can be used in that case.