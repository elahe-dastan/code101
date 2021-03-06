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
I use migrations to create the tables
```shell script
migrate create -ext sql passenger_table
```

```sql
CREATE TABLE passenger
(
    id int primary key,
    created_at timestamp,
    updated_at timestamp,
    email varchar(255),
    password varchar(255),
    fullname varchar(255),
    phone varchar(255),
    cellphone varchar(255),
    is_cellphone_verified tinyint(1) not null,
    photo_url varchar(255),
    is_email_verified tinyint(1) not null,
    locale varchar(255),
    referral_code varchar(255),
    registeration_ip varchar(255),
    is_registered_with_google tinyint(1) not null,
    email_verification_code varchar(255) not null,
    cellphone_verification_code varchar(255) not null,
    is_blocked tinyint(1) not null,
    adjust_fingerprint varchar(255) not null,
    comapny_id int(11) not null
);
```

```shell script
migrate create -ext sql ride_table
```

```sql
CREATE TABLE ride
(
    id          int PRIMARY KEY,
    passenger_id int,
    FOREIGN KEY (passenger_id) REFERENCES passenger (id)
);
```

```shell script
migrate -path . -database mysql://root:parham@tcp"(127.0.0.1:3306)"/parham up
```
for postgres:
```shell
migrate -path . -database postgres://postgres:postgres@127.0.0.1:5432/sotoon?sslmode="disable" up
```

Now our database is ready, and we want to generate the models using sqlboiler<br/>

4. install sqlboiler as dependency
```shell script
go mod init
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/null/v8
```

5. sqlboiler needs access to your database to generate the models, and to get the configurations it needs a configuration 
file called sqlboiler.toml (some other extension are ok too)
```text
[mysql]
  dbname  = "dbname"
  host    = "localhost"
  port    = 3306
  user    = "dbusername"
  pass    = "dbpassword"
  sslmode = "false"
```

6. generate models
```shell script
sqlboiler mysql
```
7. to check if everything is ok we should run the test but to run them we need mysql-client
```shell script
sudo apt install mysql-client
go test ./models
```

NOTE: **ANY** capital letter in the name of the table or columns will cause trouble 

8. CRUD operations "SELECT, INSERT, UPDATE, DELETE"

NOTE: For MySQL if using the github.com/go-sql-driver/mysql driver which is out case we should please time.Time parsing 
when making MySQL database connection.