Like always everything used to work fine **EXCEPT IN THE INTERVIEW SESSION** this is the story:<br/>
like million times before I wanted to create a simple mysql container using docker-compose

```dockerfile
version: '3'

services:
  mysql:
    image: mysql
    environment:
      MYSQL_ROOT_PASSWORD: 'raha'
    ports:
      - 3306:3306
```

I expected everything to work fine, created the container then tried to get its bash, and the first strange thing happened
though I had set MYSQL_ROOT_PASSWORD I could connect to mysql without the password, and I couldn't connect with it :||||| 
and when I wanted to run migrations I'd get 

```shell
error: Error 1130: Host '172.19.0.1' is not allowed to connect to this MySQL server
```

after investigation, I found that in MYSQL, each database user is defined with IP address in it, so you can have for example
a root user allowed to connect from localhost but not from other IP addresses. With a container, you never access to the 
database from 127.0.0.1, it could explain the problem.<br/>
Get the bash of container and connect to mysql then

```shell
SELECT host, user FROM mysql.user;
```

you'll see a table like this

```markdown
+------------+------------------+
| host       | user             |
+------------+------------------+
| %          | root             |
| 127.0.0.1  | root             |
| ::1        | root             |
| localhost  | mysql.sys        |
| localhost  | root             |
| localhost  | sonar            |
+------------+------------------+
```

it has to contain a line with your database user and '%' to work (% means "every IP address are allowed")<br/>
another strange thing: I tried to insert the line above, and it said ssl_cipher has no default value though the other lines
had NULL in ssl_cipher :|||<br/>

being a software engineer is really strange :|| magic happens everyday