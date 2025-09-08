# OTP Backend Service

A backend service written in Golang for OTP-based login and registration, featuring basic user management, Swagger documentation, and a simple frontend for testing.

---

## 🧩 Features

- Send OTP to mobile number  
- Validate OTP with a 2-minute expiration  
- Limit OTP requests to 3 per phone number within 10 minutes  
- Register new users if not existing  
- Log in existing users if OTP is valid  
- Return JWT token upon successful login  
- Retrieve user list and individual user details  
- Support for pagination and search  
- Full API documentation via Swagger  
- Simple frontend for testing OTP flow  

---

## 🛠️ Technologies Used

### Backend:
- **Go (Golang)** for fast and secure API development  
- **PostgreSQL** for persistent user data storage  
- **Redis** for temporary OTP and rate-limiting storage  
- **Swagger** for API documentation and testing  

### Frontend:
- Basic HTML/CSS/JS for OTP testing  
- Sends requests to backend and displays responses  

---

## 📦 Installation & Running

### Run Locally

```bash
go run ./cmd


Run with Docker
bash
docker-compose up --build
🌐 Project URLs
Frontend: http://localhost:8080/front/

Swagger Documentation: http://localhost:8080/swagger/index.html

🧪 Frontend Testing
The frontend is located at /front/. Open the link in your browser to test the OTP login and registration flow.

🗃️ Why PostgreSQL and Redis?
PostgreSQL
Used for storing user data because:

Reliable and secure

Supports structured relational data

Ideal for search, filtering, and pagination

Easy integration with ORMs and query builders

Redis
Used for storing OTPs and rate limits because:

Extremely fast and lightweight

Perfect for temporary data with expiration

Supports atomic operations like INCR, EXPIRE, and SETNX

Ideal for OTP validation and request throttling

📄 API Documentation
All operations are exposed via REST APIs. Swagger documentation is available at:

Code
http://localhost:8080/swagger/index.html
Includes all endpoints, parameters, sample requests, and responses.

📋 Sample API Requests
Request OTP
http
POST /api/v1/otp/request
Content-Type: application/json

{
  "phone": "09123456789"
}
Verify OTP
http
POST /api/v1/otp/verify
Content-Type: application/json

{
  "phone": "09123456789",
  "otp": "123456"
}
Get Users List
http
GET /api/v1/users?page=1&limit=10&search=0912
Authorization: Bearer <JWT>
🧱 Architecture
Clean and maintainable structure

Separation of concerns across layers (Handler, Service, Repository)

RESTful API design

Proper error handling and logging

🐳 Containerization
Fully Dockerized application

PostgreSQL and Redis included in docker-compose.yml

Run everything with a single command:

bash
docker-compose up --build