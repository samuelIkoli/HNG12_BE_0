# HNG12_BE_0
HNG stage 0 task with goLang. Simple endpoint to return my email address, the current date time in ISO 8601 format and the codebase URL of this project (this repository)

## 📁 Project Configuration

The project is divided into:

- Handlers: found in `/handlers` folder. The functions that get executed when the endpoints are called is defined here.

- Routes: found in `/routes` directory. URL endpoints and their corresponding method/action.

- Entites: found in `/entity` directory. To ensure conformity of output.

## Getting Started: Running the Server

### 🔧 Tech Stack

- GoLang
- Gin
- Air

### 📝 Requirements

This project requires go 1.23.3, Gin backend framework for go and OPTIONALLY Air module for live reload

### 💻 Running Locally

1. Clone this repository by running:
   ```bash
   git clone https://github.com/samuelIkoli/HNG12_BE_0
   cd HNG12_BE_0
   ```
2. Install the dependencies:
   ```bash
   go mod tidy
   ```
3. Start the server in dev mode (if you have air installed):
   ```bash
   air
   ```
   or (if you do not want to use air live reload)
   ```bash
   go run server.go
   ```

### 💻 Testing

Online API testing tools such as **Postman** and **Thunderclient** can be used to test the endpoints or just easily make a get request from your browser to the endpoints in the documentation. Or better still, use your terminal and Curl 😈 .

## 📖 Documentation

## Endpoints


- **URL**: `/task`
- **Method**: `GET`
- **Response**:
  ```json
  {
  "email": "myemail@gmail.com",
  "current_datetime": "",
  "github_url": "user_id"
    }
  ```



## 🔗 Link(s)

- [You can interact with the project here](https://hng12-be-0.onrender.com/task)
- [Backlink] (https://hng.tech/hire/golang-developers)
Built by SAMUEL IKOLI
