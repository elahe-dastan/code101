sqlboiler is a database first ORM, we have the passenger table of userauth and now we want to do CRUD operations on it.

1. install sqlboiler
```shell script
GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
```
NOTE: sqlboiler will be placed at `$HOME/go/bin` which was not included in my $PATH variable, it confused me for minutes
till I found out that I should add `export PATH=$HOME/go/bin:$PATH` in my .bashrc.

2. install the driver we want, our database is mysql, so we need mysql driver
```shell script
GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
```

3. I wanted to do the task using passenger table f userauth but I couldn't access it even using username and password of
the app from okd, so I create the table with the same columns and another table which has a foreign key to this one.
```shell script
docker pull mysql
docker run -e MYSQL_ROOT_PASSWORD=parham -p 3306:3306 -d mysql
docker exec -it [container-name] /bin/bash
```
inside the docker container
```shell script
mysql -uroot -p
```
inside mysql terminal
```shell script
CREATE DATABASE parham;
```
