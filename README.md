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
## clone repo and download
```go
git clone git@github.com:obh/bug_example.git
go mod download
```

##Set env variables
```
export MYSQL_USER=<your username>
export MYSQL_PASSWORD=<your password"
export MYSQL_DB=testDb
```

## Start service
go run main.go

##Run test
```bash
curl --request POST \
  --url http://localhost:7000/v1/request \
  --header 'Content-Type: application/json' \
  --data '{
	"MID": 27
}'
```
