# Eight-Sup API

**Eight-Sup API** is an open-source backend project developed to facilitate seamless interactions between streamers and their clients. Built with **Go (Golang)** and the **Echo** framework, it serves as a robust foundation for a streamer support platform.

This project handles everything from user authentication to managing subscription plans and service requests, utilizing a **MongoDB** database for efficient data management.

## 🚀 Key Features

- **RESTful Architecture**: Designed scalable API endpoints using the high-performance **Echo** framework.
- **Database Management**: Architected a non-relational database schema using **MongoDB** to manage Users, Plans, Events, and complex Streamer/Client request flows ("Wants").
- **Authentication**: Implemented secure user authentication and authorization utilizing **JWT (JSON Web Tokens)**.
- **Business Logic**: Comprehensive endpoints for subscription management, event tracking, and marketplace-like features.
- **Clean Code**: Structured application with clean architecture principles, separating models, handlers, and utility packages.

## 🛠 Tech Stack

- **Language**: Go (Golang) 1.17+
- **Framework**: Echo (v4)
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens)

## 📦 Installation & Setup

1.  **Initialize and download dependencies:**

    ```bash
    go mod init github.com/fleimkeipa/eight-sup-api
    go mod tidy
    ```

2.  **Run the application:**
    ```bash
    go run cmd/api/main.go
    ```

_Note: If you do not have the Go command on your system, you need to [Install Go](http://golang.org/doc/install) first._
