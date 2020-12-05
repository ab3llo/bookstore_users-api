# bookstore_users-api

### MySQL Database Setup

To run a MySQL database in docker run the following command

```
  docker run --name user-db -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -d mysql:latest
```