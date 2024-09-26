## Project Portfolio - Web Server
This project is a web server running locally with a PostgreSQL database. Make sure you have installed PostgresSQL in your local device.

### Prerequisites
- [Go](https://golang.org/doc/install) 
- [PostgreSQL](https://www.postgresql.org/download/)

### Clone and Setup
1.  Clone the repository:
    ```bash
    git clone https://github.com/mariopandapotansimarmata/go-todo-api  
2.  Install Module:
    ```bash
    go mod tidy  
3.  Import database:\
    cd to folder project where you clone the porject
    ```bash
    psql -U username -d todo -f todo.sql 
    
### Run
1.  Run the webserver:
    ```bash
    go run main.go  


    