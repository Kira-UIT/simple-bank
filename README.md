# **1. Set up**
* Config Makefile

Change **username** and **password**
```
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=<username> -e  POSTGRES_PASSWORD=<password> -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=<username> --owner=<username> simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://<username>:<password>@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://<username>:<password>@localhost:5432/simple_bank?sslmode=disable" --verbose down
```

* Docker
```
docker pull postgres@12-alpine
docker pull kjconroy/sqlc
```

* Run Makefle

```
make postgres
make createdb
make migrateup
make sqlc
```