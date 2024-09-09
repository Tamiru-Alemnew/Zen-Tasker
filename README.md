Zen Tasker is a versatile service built with Go and MongoDB, offering seamless functionality for efficient data handling through a RESTful API. Designed with a clean architecture, it ensures modularity, scalability, and maintainability while leveraging MongoDB for robust data storage.
### Project Structure

### **Delivery**

```
Delivery/
├── main.go
├── controllers/
│   └── controller.go
└── routers/
    └── router.go
```

### **Domain**

```
Domain/
└── domain.go
```

### **Infrastructure**

```
Infrastructure/
├── auth_middleWare.go
├── jwt_service.go
└── password_service.go
```
### **Usecases**

```
Usecases/
├── task_usecases.go
└── user_usecases.go
```

### **Repositories**

```
Repositories/
├── task_repository.go
└── user_repository.go

```
### **test**

```

tests/
├── Delivery/
│   └── controllers_test.go
├── Infrastructure/
│   ├── jwt_services_test.go
│   └── middleware_test.go
├── Repositories/
│   ├── task_repository_test.go
│   └── user_repository_test.go
├── Usecases/
│   ├── task_usecases_test.go
│   └── user_usecases_test.go
└── mocks/
    ├── JWTService.go
    ├── PasswordService.go
    ├── TaskRepository.go
    ├── TaskUsecase.go
    ├── UserRepository.go
    └── UserUsecase.go
```

### **Others**

```
docs
└── api_documentation.md
coverage.out
go.mod
go.sum
main.go
README.md
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Tamiru-Alemnew/Zen-Tasker.git
   ```

2. Change to the project directory:

   ```bash
   cd Zen-Tasker
   ```

3. Install the Go dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

Before running the application, ensure you have a MongoDB Atlas connection URl and update the configuration in the `.env` file with your specific details:

1. **Create environment file**:

   ```bash
   cp example.env .env
   ```

2. **Update the `.env` file** with the following configurations:

   ```plaintext
     MONGO_URL=<your-mongodb-atlas-url>       # MongoDB Atlas connection URL.
     JWT_SECRET=<your-jwt-secret>             # The secret key for signing JWT tokens.
   ```

   Replace `<your-mongodb-connection-string>` and `<your-jwt-secret>` with your MongoDB connection string and a secure JWT secret, respectively.

## Running the Application

To run the application, use:

```bash
make run
```

The application will start a server on port `8080`. You can access the API at `http://localhost:8080/`.

## Testing

To run tests with coverage, use:

```bash
make test
```

To generate and view a detailed test coverage report, use:

```bash
make coverage
```

## API Endpoints

- **User Authentication**
  - **Register**: `POST /register`
  - **Login**: `POST /login`
- **Task Management**
  - **Add Task**: `POST /tasks`
  - **Get All Tasks**: `GET /tasks`
  - **Get Task by ID**: `GET /tasks/{id}`
  - **Update Task**: `PUT /tasks/{id}`
  - **Delete Task**: `DELETE /tasks/{id}`
- **User Management**
  - **Promote User**: `PATCH /promot/{id}`

Refer to [Documentation](docs/api_documentation.md) for detailed API usage and request/response formats.


