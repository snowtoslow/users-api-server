user-api-server is a REST api service which provides the functionality 
to Create,Update, Partially Update, GetByID and GetAll with sorting, filter and pagination.

The server is running on port :8080. Default server config file is located in 
./config directory. 

The default config file has format: 

```json
{
  "app_port": "8080",
  "databaseName": "test.db",
  "secretKey": "banana-sausage",
  "issuer": "Vladimir" 
}
```
It could be overwritten using -config flag.Like -config="pathToFile"

Also the repository contains a postman collection which could be used for end-points testing.




```
How to run:
    1. git clone 
    2. cd users-api-server
    3. go run main.go (without config file) or go run main.go -config=*pathToMyMagicFile*
   
```
