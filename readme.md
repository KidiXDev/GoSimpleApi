# Simple API

This project is a simple API built with Go. It provides endpoints to manage user coin balances and authentication.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [License](#license)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/KidiXDev/GoSimpleApi.git
    ```
2. Navigate to the project directory:
    ```bash
    cd gosimpleapi
    ```
3. Install the dependencies:
    ```bash
    go mod tidy
    ```

## Usage

1. Start the server:
    ```bash
    go run cmd/main.go
    ```
2. The server will start on `localhost:8000`.

## Endpoints

### Get Coin Balance

- **URL:** `/account/coins`
- **Method:** `GET`
- **Headers:**
  - `Authorization: <AuthToken>`
- **Query Parameters:**
  - `username: <Username>`
- **Response:**
  ```json
  {
    "Code": 200,
    "Balance": 100
  }
  ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.