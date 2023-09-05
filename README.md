# Golang Packages

1. compile daemon to watches your file and make changes just like Nodemon in nodejs
    go get github.com/githubnemo/CompileDaemon
    // To run compile daemon
    CompileDaemon -command="./go-gin-gorm"

2. .env file to load environment variables
    go get github.com/joho/godotenv

3. Gin Framework
    go get -u github.com/gin-gonic/gin

4. GORM ORM Library for Golang
    go get -u gorm.io/gorm

5. Database driver for postgres
    go get -u gorm.io/driver/postgres

6. Gorm Auto migrate command
    go run migrate/migrate.go

7. Download TablePlus to view your Postgres Custom Tables
    https://tableplus.com


# Database set up
Use Elephant SQL to set up your postgres database connection
Please refer to the gorm documentation for database migration , just follow the steps

https://gorm.io/docs/migration.html

please refer to my .env file to see how i set up my environment variable -- This is strictly for educational purpose. 
