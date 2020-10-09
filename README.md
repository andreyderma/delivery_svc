## How to run locally
Requirements to run locally (Without docker):
- Golang version: go1.13.6
- MySQL 5.6
 
Create .env file with some environment variables:
```
MYSQL_USERNAME=your_username
MYSQL_PASSWORD=your_password
MYSQL_HOST=localhost
MYSQL_PORT=3306
DATABASE_NAME=delivery_app
APP_PORT=:5000
ENV=dev
MAP_API_KEY=your_google_map_api_key
```

1. Create database `delivery_app` and import database in `db/init_db.sql`
2. Run using make
```
make all
```
Application service ready on port 5000, I also provide swagger interface accessible at 
`http://localhost:5000/v1/delivery_svc/swagger/index.html`

To run unittest separately:
```
go test -v -count=1 ./... -coverprofile=c.out
```

## How to run using docker container
You need to provide google api key and run start.sh as follow
```
./start your_google_map_api_key

example:

./start AIzaSyDUYcUu5xG8-R5_X5A0eSWwBXmE-WRONG-KEY
```

`start.sh` will execute `docker-compose` command and will run 2 services, MySql service and Application service.
If there are no errors during start, the service can be access on port 8080.
`http://localhost:8080/v1/delivery_svc/swagger/index.html`


### Notes
I put some environment variables in `docker-compose.yml`. 
In production, we need to store in centralized way, either using consul KV, 
AWS SSM parameter store or other 3rd party service for storing service configuration and metadata.