# Event Management REST API

This project is a RESTful API for managing events and user registrations. It is built with Go (Golang) and utilizes the Gin framework for routing, along with SQL for database management.

## Features

- User authentication (signup/login)
- Event management (create, update, delete, view)
- User registration for events
- Cancel registration for events
- Secure API endpoints with JWT authentication

## Technologies Used

- **Programming Language**: Go (Golang)
- **Web Framework**: Gin
- **Database**: SQL (MySQL/PostgreSQL compatible)
- **Authentication**: JWT

## Project Structure

```plaintext
.
├── db/                  # Database connection and management
├── middleware/          # Middleware (e.g., authentication)
├── models/              # Data models and database interaction logic
├── routes/              # Route definitions
├── utils/               # Utility functions (e.g., password hashing)
├── main.go              # Entry point of the application
└── README.md            # Project documentation
```

## Installation and Setup

### Prerequisites

- Go (version 1.19 or later)
- MySQL or PostgreSQL database

### Steps to Run the Project

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/event-management-api.git
   cd event-management-api
   ```

2. **Set Up Environment Variables**: Create a `.env` file in the root directory with the following content:

   ```env
   DB_USER=<your-database-username>
   DB_PASSWORD=<your-database-password>
   DB_NAME=<your-database-name>
   DB_HOST=localhost
   DB_PORT=3306
   JWT_SECRET=<your-secret-key>
   ```

3. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

4. **Run Database Migrations**: Set up your database schema manually or with a migration tool.

5. **Start the Server**:

   ```bash
   go run main.go
   ```

6. **Test the Endpoints**: Use a tool like Postman or curl to interact with the API.

## API Endpoints

### Authentication

| Method | Endpoint  | Description |
| ------ | --------- | ----------- |
| POST   | `/signup` | User signup |
| POST   | `/login`  | User login  |

### Event Management

| Method | Endpoint      | Description          |
| ------ | ------------- | -------------------- |
| GET    | `/events`     | Get all events       |
| GET    | `/events/:id` | Get a specific event |
| POST   | `/events`     | Create a new event   |
| PUT    | `/events/:id` | Update an event      |
| DELETE | `/events/:id` | Delete an event      |

### User Registration

| Method | Endpoint               | Description               |
| ------ | ---------------------- | ------------------------- |
| POST   | `/events/:id/register` | Register for an event     |
| DELETE | `/events/:id/register` | Cancel event registration |

## Example HTTP Requests

### Register for an Event

```http
POST http://localhost:8080/events/1/register
Authorization: Bearer YOUR_JWT_TOKEN
Content-Type: application/json
```

### Cancel Event Registration

```http
DELETE http://localhost:8080/events/1/register
Authorization: Bearer YOUR_JWT_TOKEN
Content-Type: application/json
```

## Contributing

1. Fork the repository.
2. Create a new branch for your feature.
3. Commit your changes.
4. Push to your fork and submit a pull request.
