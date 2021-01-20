# MySQL Query Timeout
I once wondered if I can put a timeout constraint on my SQL queries and YES, here is the process if you want to try it too.

## MySQL database
First of all, we need a mySQL database to put our constraint on.
```shell script
docker pull mysql #pull mysql image 
docker run --name [container_name] -p [open_port]:3306 -e MYSQL_ROOT_PASSWORD=[root_password] -d mysql 
docker exec -it [container_name] /bin/bash # get the bash of container
```
inside the bash of the container, run below commands
```shell script
mysql -uroot -p # connect to mysql server
CREATE DATABASE [database_name];
``` 

## Fill Database with Data
to check if the timeout constraint works or not se should have a select query that takes long, so I have to fill the table 
with data. There are plenty of ways to do this and you can choose the one which is the easiest for you, I tried writing 
a small golang code.
 
