# **Sample Leon Service**

#### **Tools used**
    - Golang version 1.16
    - Docker
    - Docker compose
    - Postman
    
#### **Docker Compose Commends**
    - docker-compose up --build -d | Build and run the service | port:8080
    
#### **Run without Docker**
    * Note : you have to adjust the database config to your local database, you execute the init.sql file
             
    - go run main.go | run the service | port:8081
    - go test ./.. | run integration test 