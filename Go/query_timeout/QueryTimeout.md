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
## migrations
the files should be created automatically

to check if the timeout constraint works or not se should have a select query that takes long, so I have to fill the table 
with data. There are plenty of ways to do this and you can choose the one which is the easiest for you, I tried writing 
a small golang code.

## set constraint
### in query
SELECT /*+ MAX_EXECUTION_TIME(1) */ * FROM parham;

in the mysql client whether you put a constraint in the query or a constraint on the mysql server, if your query takes 
longer then it will just return an error and no rows but if you use go-sql-driver/mysql library to query the database then
if you want to put constraint in the query you can specify it in the query or use context. If you use context then it 
will exceed time doesn't give you back anything but if you put constraint in the query then it won't return any error 
and returns as many rows as it could fetch within the max execution time.
# note
One scenario is this, you want to get all the table and do something, you use constraint in the query, you get one third 
of the rows and cause there are no errors you think you have fetched all the table. So if you want to put a constraint
use context. BUT what if there is a constraint on the mysql server itself??
Usually the max execution time of a mysql server is too high that no one ever hits it but if the were a tight constrait
there are two ways, first, each query should have context with the same timeout, second 

 
