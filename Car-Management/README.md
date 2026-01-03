# Car Management System

A RESTful API built with Go and PostgreSQL for managing cars and their engines.

## Features

- **Car Management**: Create, Read, Update, and Delete car records.
- **Engine Management**: Create, Read, Update, and Delete engine specifications.
- **Database Integration**: Persistent storage using PostgreSQL.
- **Dockerized**: Easy setup and deployment using Docker and Docker Compose.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd Car-Management
    ```

2.  **Start the services:**
    ```bash
    docker-compose up --build
    ```
    This command builds the application and starts both the Go server and the PostgreSQL database.

3.  **Access the API:**
    The server will be available at `http://localhost:8080`.

## API Endpoints

### Cars

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/cars` | Get all cars |
| `GET` | `/cars/{id}` | Get car by ID (UUID) |
| `GET` | `/cars/brand/{brand}` | Get cars by brand |
| `POST` | `/cars` | Create a new car |
| `PUT` | `/cars/{id}` | Update an existing car |
| `DELETE` | `/cars/{id}` | Delete a car |

**Example Car Payload (POST/PUT):**
```json
{
    "name": "Civic",
    "year": "2023",
    "brand": "Honda",
    "fuel_type": "Petrol",
    "engine": {
        "engine_id": "e1f86b1a-0873-4c19-bae2-fc60329d0140",
        "displacement": 2000,
        "no_of_cylinders": 4,
        "car_range": 600
    },
    "price": 25000
}
```

### Engines

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/engines` | Get all engines |
| `GET` | `/engines/{id}` | Get engine by ID (UUID) |
| `POST` | `/engines` | Create a new engine |
| `PUT` | `/engines/{id}` | Update an existing engine |
| `DELETE` | `/engines/{id}` | Delete an engine |

**Example Engine Payload (POST/PUT):**
```json
{
    "displacement": 2000,
    "no_of_cylinders": 4,
    "car_range": 600
}
```

## Environment Variables

The application uses the following environment variables (configured in `docker-compose.yml` and `.env`):

- `PORT`: The port the server listens on (default: `8080`).
- `DB_HOST`: The hostname of the PostgreSQL database (`db`).
- `DB_PORT`: The port of the PostgreSQL database (`5432`).
- `DB_USER`: The database user (`postgres`).
- `DB_PASSWORD`: The database password.
- `DB_NAME`: The database name.
- `SCHEMA_FILE`: Path to the SQL schema file.

## Development

To run the application locally without Docker:

1.  Ensure you have a PostgreSQL instance running.
2.  Update the `.env` file with your database credentials.
3.  Install dependencies:
    ```bash
    go mod download
    ```
4.  Run the application:
    ```bash
    go run main.go
    ```
