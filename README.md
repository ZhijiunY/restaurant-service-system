# Restaurant Service System
This project implements a basic restaurant service system using Go, Gin framework for the web server, PostgreSQL for the database, and Redis for session management and caching. It provides functionalities like user authentication, session management, menu browsing, and order submission.

## Getting Started
To get started with this project, follow the steps below:

### Prerequisites
- Go (version 1.13 or higher)
- PostgreSQL
- Redis
- A proper Go workspace set up

### Installation
1. Clone the repository:
```
git clone https://github.com/ZhijiunY/restaurant-service-system.git
```
2. Load environment variables:

The project uses github.com/joho/godotenv for loading environment variables. Make sure to create a .env file at the root of your project and define your PostgreSQL and Redis configurations:
```
DB_HOST=your_database_host
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
DB_PORT=your_database_port
REDIS_ADDR=your_redis_address
REDIS_PASSWORD=your_redis_password
REDIS_DB=your_redis_db
```
3. Install Go dependencies:
Navigate to the project directory and install the required Go modules:
```go 
go mod tidy
```
4. Database Migrations:
The project uses GORM for database operations. Run migrations to set up your database schema:
```
go run migrations/migrate.go
```
This will create the necessary tables in your PostgreSQL database.
5.Running the server:
To start the server, run:
```
go run main.go
```
This will start the Gin web server on port 8080.

### Usage
Once the server is running, you can access the following endpoints:
```
/ - Homepage
/menu - Displays the menu (authentication required)
/order - Page to submit a new order (authentication required)
/submit-order - Endpoint to submit a new order (authentication required)
/auth/getlogin - Login page
/auth/signup - Signup page
/auth/login - Endpoint for user login
/auth/logout - Endpoint for user logout
/auth/signup - Endpoint for user signup
```

### Authentication
The project uses sessions for authentication. A user needs to sign up and log in to access protected routes like viewing the menu and submitting orders.

### Session Management
Redis is used for storing session data, providing fast access and management of session information across requests.

### Static Files and Templates
The project serves static files and HTML templates for the frontend. The templates are located in the ./templates directory, and static files (e.g., CSS, JavaScript) are served from the ./static directory.

### Logging
Logging is commented out in the main.go file but can be easily enabled by uncommenting the setupLogging function and its call. This function configures Gin to log to both a file and stdout.

## Contribution
Feel free to fork the repository and submit pull requests to contribute to the project.

## License
This project is open-sourced under the MIT License. See the LICENSE file for more details.