📝 Project Documentation: Task Management System (Clean architecture + Gin + MongoDB)
## 📄 API Documentation

You can find detailed API documentation in the following link:
https://documenter.getpostman.com/view/39486846/2sB34mgx6T

### 📌 Project Overview

This is a secure task management REST API built using:

* **Gin** for routing and middleware
* **MongoDB** for data persistence
* **JWT** for authentication and user identification
* Organized following **Clean Architecture principles** to ensure separation of concerns

---

### 🧱 Project Structure

```
/config         → Loads environment variables and sets up DB  
/controller     → Handles HTTP requests and responses  
/docs           → Contains the documentation  
/domain         → Defines core domain models  
/dto            → Structures for incoming and outgoing JSON data  
/infrastructure → Core services: JWT, password hashing, auth middleware  
/repository     → Handles all MongoDB database operations  
/usecases       → Contains the business logic and use case flows  
/main.go        → App entry point with server and route initialization  
```


### ✅ Features Implemented

* 🔐 **Authentication**

  * Signup and login via `username` and `password`
  * Secure password hashing
  * Token generation using JWT

* 📄 **Task Management**

  * Create, read, update, delete tasks
  * Each task is tied to a specific user by `username`
  * Only authenticated users can access their own tasks


### 🔐 Authentication Flow

1. **Signup**

   * User provides `username`, `password`, and `role`
   * Password is securely hashed before storing in the database

2. **Login**

   * System verifies credentials and returns a signed JWT

3. **Protected Routes**

   * Requires `Authorization: Bearer <token>` header
   * Middleware extracts and verifies token
   * Routes use `username` from the token to authorize access


### 📁 Environment Setup (`.env`)

All sensitive settings are defined in a `.env` file, such as:

* `PORT` → Server port
* `MONGODB_URI` → MongoDB connection string
* `JWT_SECRET` → Secret for signing JWTs
* `DB_NAME` → MongoDB database name

> The `.env` is loaded via the `/config` package to keep secrets out of the source code.



### 📦 DTO Usage

DTOs (**Data Transfer Objects**) are used to:

* Accept structured input for creating or logging in users
* Return filtered and safe responses to the client

This abstraction:

* Keeps the internal domain logic clean
* Makes it easy to change the format of requests/responses without affecting core logic


### 🧱 Clean Architecture (Inspired by Uncle Bob)

This project follows **Clean Architecture principles** with strict layer separation:

#### 1. **Entities (Domain Layer)**

* Core business objects (e.g., `User`, `Task`)
* Independent of frameworks or technologies

#### 2. **Use Cases**

* Contains business logic
* Interacts with entities and defines workflows (e.g., register user, create task)
* Depends only on domain models and repository interfaces

#### 3. **Interface Adapters**

* Includes:

  * **Controllers** (translates HTTP requests to use case calls)
  * **DTOs** (binds/returns JSON)
  * **Repositories** (implements DB logic via interfaces)

#### 4. **Frameworks and Drivers (Infrastructure Layer)**

* Includes:

  * **Gin** (HTTP framework)
  * **MongoDB driver**
  * **JWT and hashing libraries**

> Dependency flows **inward** only — outer layers depend on inner layers, never the reverse.
