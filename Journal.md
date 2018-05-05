
# Space Team Project Journal

## Week 1 (4/14-4/21)

**April 10:**

+ Design and brainstroming the project work flow, structures and how all those personal databases.
can connect with a Go API and each database serve for one single feature of the API.  
+ Coding HTML, JavaScript for front end of Counter Burger order.
+ Coding Go and Redis for back end to connect to database.
+ Each master node will talk with the rest of the nodes to replicate the data. 

+ The order process:
 - Registry as a customer with names and email.
 - Login and build a burger for themselve.
 - Come to a specified Counter Burger Store, pick up and pay.

**Apr 15 :**

 + Webex to discuss the project design.
 + Discuss the sequence of customer placing an order.
 + Design User Interface and scope of project.

**Apr 17 :**

+ Decide technologies and languages for the the project.

**April 21:**

+ Webex to discuss the project design.
+ Tasks distributions of databases.

    Hoang Nguyen: Customer.
    Huy Huynh: Menu.
    Pratik Mehta: Payment.
    Vu Nguyen: Order.

_Challenge_

+ The project idea is brand new. All team member did not know this method is exist. 
+ Need to read alots to understand Redis, Go and AWS EC2 instances work together.

_Test_
 + Test AWS EC2 instances for all personal projects works and the port are opened for outside connection.

## Week 2 (4/22 - 4/28)

**April 23:**

+ Setup database Redis on AWS and coding Go API.

+ Set Access Redis database.
 
```client = redis.NewClient(&redis.Options{
      Addr:     "54.67.84.81:6379",
      Password: "", // No password
      DB:       0, // use default DB
	})
```
**April 24:**

_Done_

+ Finish the database design and implement for Customer, Menu and Order.
+ Test the connection and verify "Insert", "Retrieve".

**April 26:**

+ Setup partition tolerence demo in each members.
+ Access to each database to verify connection with Go.
+ First try to registry and place an order

**April 27:**

+ Implement web UI structure and logic of the order.

_Challenge_
  - Redis is simple to use but the logic is complitely new with "keys" concept.
  - The problem of customer place order links from Registration - Menu - Order. 
  
_Test_
  Run main.go to check insert data into database.

## Week 3 (4/28 - 5/4)

**April 30:**

Finish personal project to setup database and partition tolerence in each database.

**May 1:**
+ Webex to discuss the project design

Continue Go API to insert, retrieve and delete in Redis database 

**May 2:**

Finish Login/Signup page to test with GO API

**May 3:** 

 + Test the HTML pages for Login, Registration, Menu and Orders

 + Test Redis database with GO API. The API now can access all the databases and bind the front end to process registration, handel order.

 **May 4:** 
 
 + Test again the project UI and database handling with Go.

 + Test one cluster that its Redis server works.

 **A - All 5 nodes are up and running**

  * Master:

key $ssh -i "cmpe281-us-west-1.pem" 
ec2-user@ec2-54-67-84-81.us-west-1.compute.amazonaws.com

Last login: Sat May  5 03:59:03 2018 from 130.65.254.5

       __|  __|_  )
       _|  (     /   Amazon Linux AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-ami/2018.03-release-notes/

No packages needed for security; 3 packages available

Run "sudo yum update" to apply all updates.

[ec2-user@ip-10-0-0-85 ~]$ redis-cli

127.0.0.1:6379> 

 * Slave 1:

key $ssh -i "cmpe281-us-west-1.pem" 
ec2-user@ec2-54-193-65-181.us-west-1.compute.amazonaws.com

Last login: Sat May  5 04:34:09 2018 from 130.65.254.5

       __|  __|_  )
       _|  (     /   Amazon Linux AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-ami/2018.03-release-notes/

No packages needed for security; 3 packages available

Run "sudo yum update" to apply all updates.

[ec2-user@ip-10-0-0-47 ~]$ redis-cli

127.0.0.1:6379> 

* Slave 2:

key $ssh -i "cmpe281-us-west-1.pem" 

ec2-user@ec2-54-193-126-82.us-west-1.compute.amazonaws.com

Last login: Sat May  5 04:26:51 2018 from 130.65.254.5

       __|  __|_  )
       _|  (     /   Amazon Linux AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-ami/2018.03-release-notes/

No packages needed for security; 3 packages available

Run "sudo yum update" to apply all updates.

[ec2-user@ip-10-0-0-90 ~]$ redis-cli

127.0.0.1:6379>

* Slave 3:

key $ssh -i "cmpe281-us-west-1.pem"

ec2-user@ec2-13-56-213-228.us-west-1.compute.amazonaws.com

Last login: Sat May  5 04:26:48 2018 from 130.65.254.5

       __|  __|_  )
       _|  (     /   Amazon Linux AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-ami/2018.03-release-notes/

No packages needed for security; 3 packages available

Run "sudo yum update" to apply all updates.

[ec2-user@ip-10-0-0-240 ~]$ redis-cli

127.0.0.1:6379>

* Slave 4:

key $ssh -i "cmpe281-us-west-1.pem"

ec2-user@ec2-13-57-32-130.us-west-1.compute.amazonaws.com

Last login: Sat May  5 04:26:54 2018 from 130.65.254.5

       __|  __|_  )
       _|  (     /   Amazon Linux AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-ami/2018.03-release-notes/

No packages needed for security; 3 packages available

Run "sudo yum update" to apply all updates.

[ec2-user@ip-10-0-0-57 ~]$ redis-cli

127.0.0.1:6379>

