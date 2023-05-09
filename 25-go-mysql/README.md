## Connect Go Applications to MySQL Databases

### Learning Objective

Connect a Go application to a MySQL database engine.
Create databases, create tables, insert data rows,
and select data rows from the MySQL engine.
Additionally, we will turn data rows into custom Go structs.

### Commands

Create Go module with MySQL driver:

```
go mod init cbtmysql
go get -u github.com/go-sql-driver/mysql 
```

Install Docker Engine on Ubuntu Linux:

```
sudo apt update && sudo apt-get install docker.io --yes
```

Create a MySQL container using Docker:

```
docker run --detach \
  --restart=unless-stopped \
  --publish=3306:3306 \
  --env MYSQL_ROOT_PASSWORD=mysqlpassword123 \
  docker.io/mysql:latest
```

Get interactive MySQL CLI in MySQL server container:

```
docker exec -it <containerID> bash

# Then run ...

mysql --password
```