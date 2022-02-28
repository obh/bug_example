##Create database
`create database testDb;`

##Create table
```
CREATE TABLE `RequestLog` (
  `id` int NOT NULL AUTO_INCREMENT,
  `merchantId` int DEFAULT NULL,
  `request` varchar(4000) DEFAULT NULL,
  `response` varchar(4000) DEFAULT NULL,
  `addedOn` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);
```
##Set env variables
```
export MYSQL_USER=<your username>
export MYSQL_PASSWORD=<your password"
export MYSQL_DB=testDb
```

## Start service
go run main.go
