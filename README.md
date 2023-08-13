# Todo API

Todo API is a small todo api backend built in Go

## Try live version

### To create a new item 
Use the endpoint using POST request
```bash
ec2-18-144-22-131.us-west-1.compute.amazonaws.com:8000/api/create
```
Pass in json

```json
{
	"name":"Buy bread"
}
```

### To create get all items
Use the endpoint using GET request
```bash
ec2-18-144-22-131.us-west-1.compute.amazonaws.com:8000/api/todos
```

### To update an item 
Use the endpoint using PUT request
```bash
ec2-18-144-22-131.us-west-1.compute.amazonaws.com:8000/api/update
```
Pass in json

```json
{
	"id": 1,
	"name": "Buy bread",
	"finished": true
}
```


### To delete an item 
Use the endpoint using DELETE request 
```bash
ec2-18-144-22-131.us-west-1.compute.amazonaws.com:8000/api/delete
```
Pass in json

```json
{
	"ids": [1]
}
```

## Locally

### Migrations

Will need to run so we have a db to use
```bash
docker compose up postgresdb
```

install goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
or on mac 
```bash
brew install goose
```

cd into migration folder 
```bash
cd internal/repositories/db/migrations
```
Run:
```bash
 goose postgres "user=<user> password=<password> dbname=<db name> sslmode=disable" up
```

You can clone and run locally
Just edit the .env file
and make sure docker is installed and run 

```bash
docker compose up
```