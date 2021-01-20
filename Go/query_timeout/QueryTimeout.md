# MySQL Query Timeout
I once wondered if I can put a timeout constraint on my SQL queries and YES, here is the process if you want to try it too.

## MySQL database
First of all, we need a mySQL database to put our constraint on.
```shell script
docker pull mysql #pull mysql image 
docker run --name [container_name] -p [open_port]:3306 -e MYSQL_ROOT_PASSWORD=[root_password] -d mysql 
```