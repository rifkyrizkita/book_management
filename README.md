Book Management API

This project is a Book Management API built with Golang and Fiber, designed to allow users to manage books, borrowing, and categories. It provides endpoints for borrowing books, fetching borrowed books, and more.
Technologies Used

    Go (Golang): Backend language used for building the API.
    Fiber: Web framework for Go that provides fast HTTP routing and middleware.
    GORM: Object Relational Mapping (ORM) for Go to interact with the database.
    MySQL: Database for storing book and user data.
    Docker: Containerization for easy deployment and running of the application.

Setup and Installation
Prerequisites

Make sure you have the following installed:

    Docker
    Docker Compose
    Go 1.21+ (for local development)

Clone the Repository

git clone https://github.com/yourusername/book-management-api.git
cd book-management-api

Using Docker (Recommended)

This project is containerized using Docker. To build and run the app using Docker Compose, follow these steps:

    Build and start the application:

    docker-compose up --build

    This will start both the go-api and mysql services as defined in the docker-compose.yml.

    Access the application: The API will be running on http://localhost:8080.


Without Docker (Local Development)

    Install dependencies: If you want to run the project without Docker, you can follow these steps:
        Install Go (version 1.21+)
        Install MySQL (local instance or use a service like MySQL Workbench or XAMPP)

    Install Go modules:

go mod tidy

Run the application:

    go run main.go

    This will start the application at http://localhost:8080.

Environment Variables

The following environment variables must be set for the application:

    DB_HOST: Host for the MySQL database
    DB_PORT: Port of the MySQL database 
    DB_USER: MySQL username 
    DB_PASSWORD: MySQL password 
    DB_NAME: The name of the database 
