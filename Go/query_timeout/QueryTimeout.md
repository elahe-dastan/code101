# MySQL Query Timeout
I once wondered if I can put a timeout constraint on my SQL queries and YES, here is the process if you want to try it too.

# Table of Content
1. Setting Up a MySQL Database in Docker
2. Fill Database with Data
3. Set Constraint on the Query
4. Set Constraint on the Database

## Setting Up a MySQL Database in Docker
First of all, we need a mySQL database to put our constraint on. I prefer to set up a database in docker, so we don't care
what your os is.
```shell script
docker pull mysql
docker run --name [container_name] -p [open_port]:3306 -e MYSQL_ROOT_PASSWORD=[root_password] -d mysql 
docker exec -it [container_name] /bin/bash # get the bash of container
```
inside the bash of the container, run below commands
```shell script
mysql -uroot -p # connect to mysql server
CREATE DATABASE [database_name];
``` 

## Fill Database with Data
You can manually create a table inside the database, and fill it in many ways but preferred to use `migrations` to create
the table and write a simple golang code to fill it.
```shell script
brew install golang-migrate
cd [specific_directory]
migrate create -ext sql [name]
```
inside the file with the extension up.sql I just wrote
```mysql
CREATE TABLE [table_name]
(
    id int PRIMARY KEY,
    [column_name] [column_type]
);

```
It's a good idea to have even more columns.

Below code run migration to create table 
```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:[root_password]@tcp(127.0.0.1:[container_port])/[dbName]")
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil{
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://[migrations_directory]",
		"[dbName]",
		driver,
	)
	if err != nil{
		log.Fatal(err)
	}

	if err := m.Up(); err != nil{
		log.Fatal(err)
	}
}
```
fill the table with data
```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
    	db, err := sql.Open("mysql", "root:[root_password]@tcp(127.0.0.1:[container_port])/[dbName]")
    	if err != nil {
    		panic(err.Error())
    	}

	defer db.Close()

	//fill the table with data
	for i := 0; i < 5000; i++ {
		_, err := db.Exec("INSERT INTO [table_name] VALUES (?, [value]);", i)
		if err != nil{
			log.Fatal(err)
		}
	}

}
```

I decided to put 5000 records to the table, but you can do what you want. 

## Set Constraint on the Query
If you are coding in golang like me one the approaches I can suggest is using contexts
```go
package main

import (
	"context"
	"database/sql"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:[root_password]@tcp(127.0.0.1:[container_port])/[dbName]")
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

    // execution time
	ctx, cancel := context.WithTimeout(context.Background(), [time])
	defer cancel()

    _, err = db.QueryContext(ctx,"SELECT * FROM [table_name];")
	if err != nil{
		log.Fatal(err)
	}
}
```
If the query takes longer than the time you set, this error will be returned.
```shell script
context deadline exceeded
```
Another way is to specify the constraint inside the query:<br/>
After getting the bash of the container and connecting to mysql server run below query
```shell script
SELECT /*+ MAX_EXECUTION_TIME([time]) */ * FROM [table_name];
``` 
If the time you set is not enough, you'll get this error
```shell script
ERROR 3024 (HY000): Query execution was interrupted, maximum statement execution time exceeded
```

Now let's try the same query inside our golang code. What do we expect??!! I think it's reasonable if it returns an error 
and no rows.  
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:[root_password]@tcp(127.0.0.1:[container_port])/[dbName]")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT /*+ MAX_EXECUTION_TIME([time]) */ * from [table_name];")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name);err != nil{
			log.Fatal(err)
		}
		fmt.Println(id)
	}
}
```
Surprisingly, there were no errors 
```shell script
0
1
.
.
.
2826
2827
```
As you can see it just gave us as many records as it could in the specified time and didn't return any errors :neutral_face::neutral_face::neutral_face:.
WHY? That's because whey you write a query in mysql client it gives you the response for example in our case the whole
table at the same time OR an error but when you use libraries in your code to connect to database and write a query what 
it gives you back is a CURSOR, and then it uses the cursor to iterate over the records of the table you want and gives you 
the result, when the time limit exceeds it just finishes its job and doesn't get any error to return to you. Not knowing 
the point may cause bug in the behavioral of your code. 

## Set Constraint on the Database
We saw how to put constraint on the query but more commonly you like to put constraint on the database. You can easily use
below commands to do this.
```shell script
SET SESSION MAX_EXECUTION_TIME=[time];
SET GLOBAL MAX_EXECUTION_TIME=[time];
```
`SET SESSION MAX_EXECUTION_TIME=[time];` sets the constraint only on the specific session and 
`SET GLOBAL MAX_EXECUTION_TIME=[time];`sets the constraint on all the sessions.

**Note**: The results after setting these constraints are completely predictable, I just want to mention a point which may
confuse you.Suppose you have a session to your mysql server and want to put a `GLOBAL` constraint on it after setting the
constraint, it is not applied to the current session if you close the current connection and open a new one you see that 
everything works fine and as expected.

## Update
Now let's see how we can put timeout constraint on database query when we're using sqlx or gorm(both slqx and gorm are 
popular libraries for communication with database in golang)

## SQLX
Everything I said about SQL **remains exactly the same** for this part. So you can both put `/*+ MAX_EXECUTION_TIME([time]) */`
inside the query or use `QueryxContext`.

## Gorm
In gorm version 2 you can put timeout constraint on queries using `WithContext` but in Version 1 the only way you can do 
it is by using transactions 

migrate in terminal example<br/>
migrate -path "migrations dir" -database "driver://dsn" command
migrate -path . -database postgres://postgres:parham@172.17.0.2:5432/parham?sslmode=disable up