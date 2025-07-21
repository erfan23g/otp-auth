# Go OTP Authentication Service

This project provides a simple, Dockerized One-Time Password (OTP) authentication service built with Go. It allows users to request an OTP code via their phone number, verify that code, and receive a JSON Web Token (JWT) upon successful verification.

---

## Features

* **OTP Generation & Verification**: Securely generates and verifies 6-digit OTP codes.
* **2-Minute TTL**: Each generated OTP code is valid for 2 minutes.
* **Rate Limiting**: Prevents abuse by limiting OTP requests to 10 per minute per user.
* **Redis Integration**: Stores OTP codes and rate limit information for fast access and expiry.
* **PostgreSQL Integration**: Persists user phone numbers and creation timestamps.
* **JWT Authentication**: Issues a JWT upon successful OTP verification for secure subsequent requests.
* **Dockerized**: Easily set up and run the entire service using Docker Compose.

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have the following installed:

* [Docker Desktop](https://www.docker.com/products/docker-desktop) (includes Docker Engine and Docker Compose)

### Installation and Setup

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/erfan23g/otp-auth.git
    cd otp-auth
    ```

2.  **Start the Service**:
    Navigate to the root directory of the project and run:
    ```bash
    docker compose up
    ```
    This command will build the Docker images (if not already built), start the PostgreSQL database, Redis instance, and the Go application.

---

## Usage

Once the service is running, you can interact with it using `curl` or any API client. The service will be accessible on `http://localhost:8080`.

### 1. Request an OTP Code

Send a POST request to the `/api/v1/otp/send` endpoint with the user's phone number.

**Endpoint**: `POST /api/v1/otp/send`
**Content-Type**: `application/json`

**Example Request**:

1. Request an otp code
```bash
curl -X POST http://localhost:8080/api/v1/otp/send \
  -H "Content-Type: application/json" \
  -d '{"phone": "+11111111111"}'
  ```
Note: After sending the request, the generated OTP code will be logged in the app.log file.

2. Verify OTP and Get JWT

Send a POST request to the /api/v1/otp/verify endpoint with the phone number and the OTP code you received.

Endpoint: POST /api/v1/otp/verify
Content-Type: application/json

Example Request:

```bash
curl -X POST http://localhost:8080/api/v1/otp/verify \
  -H "Content-Type: application/json" \
  -d '{"phone": "+11111111111", "code": "123456"}' # Replace "123456" with the actual
  OTP from the logs
```
Example Successful Response:

```JSON
{
  "token": "your_generated_jwt_token_here"
}
```
Important: The OTP code is only valid for 2 minutes from the time it was generated.


## Technologies Used
Go: Primary programming language

Redis: In-memory data store for OTPs and rate limiting. Used because of it's fast response for large amount of requests.

PostgreSQL: Relational database for user data. Used because of it's simplicity for not too large number of users.

## Contact
[Erfan Ghorbani/erfan.ghorbani.1204@gmail.com/https://github.com/erfan23g]