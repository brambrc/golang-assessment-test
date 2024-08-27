# Mezink Golang Assessment Project

This project is a simple Golang application that includes a RESTful API with PostgreSQL as the database. The project is containerized using Docker and Docker Compose for easy setup and deployment.

## Getting Started

### Prerequisites

Make sure you have the following installed on your local machine:

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

### Cloning the Repository

Clone the repository to your local machine using Git:

```bash
git clone https://github.com/your-username/mezink-golang-assessment.git
cd mezink-golang-assessment
```

### Running the Project

1. Ensure you are in the root directory of the project where the docker-compose.yml file is located.
2. Build and run the Docker containers:
```bash
docker-compose up --build
```
This command will:
- Build the Go application.
- Set up and start a PostgreSQL database.
- Run database migrations.
- Start the Go application on port 8080.

Access the application:
3. The application will be running at http://localhost:8080.

## API Documentation
### 1. Fetch Records (POST /api/records)

This API endpoint inserts a sample record into the database and fetches records based on the provided filters.

#### Request Payload:
```json
{
    "startDate": "YYYY-MM-DD",
    "endDate": "YYYY-MM-DD",
    "minCount": 100,
    "maxCount": 300
}
```
#### Explanation:
- startDate: Start date for filtering records based on the createdAt timestamp.
- endDate: End date for filtering records based on the createdAt timestamp.
- minCount: Minimum sum of the marks array in the record.
- maxCount: Maximum sum of the marks array in the record.

#### Possible Responses:
- Success (200 OK):
```json 
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "id": 1,
            "createdAt": "2024-08-27T12:00:00Z",
            "totalMarks": 250
        },
        {
            "id": 2,
            "createdAt": "2024-08-27T12:01:00Z",
            "totalMarks": 300
        }
    ]
}
```
- Error (400 Bad Request or 500 Internal Server Error):
 ```json 
{
  "code": 1,
  "msg": "Error message"
}
```

### 2. Fetch All Records (GET /api/fetch-table)
This API endpoint fetches all records from the database.
#### Possible Responses:
- Success (200 OK):
```json 
{
  "code": 0,
  "msg": "Success",
  "records": [
    {
      "id": 4,
      "createdAt": "2024-08-27T10:33:48.590197Z",
      "totalMarks": 250
    },
    {
      "id": 14,
      "createdAt": "2024-08-27T12:32:59.63759Z",
      "totalMarks": 408
    },
    {
      "id": 3,
      "createdAt": "2024-08-27T10:32:23.933847Z",
      "totalMarks": 250
    },
    {
      "id": 13,
      "createdAt": "2024-08-27T12:16:14.212694Z",
      "totalMarks": 442
    }
  ]
}
```
- Error (400 Bad Request or 500 Internal Server Error):
 ```json 
{
  "code": 1,
  "msg": "Error message"
}
```

## Notes
- Ensure that your Docker environment is correctly configured and running before attempting to start the service.
- The database credentials and other sensitive information are stored in the environment variables or the docker-compose.yml file.