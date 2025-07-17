# Library Management System Documentation

## Overview

This is a simple Library Management System built with Go. It allows you to manage books and users, and supports basic borrowing and returning operations via a terminal interface.

## Features

- Add, delete, and list books
- Register and manage users
- Borrow and return books

## Setup

1. **Clone the repository**
   ```sh
   git clone https://github.com/Edenbel27/goland.git
   cd library_management
   ```

2. **Install dependencies**
   ```sh
   go mod tidy
   ```

3. **Run the application**
   ```sh
   go run main.go
   ```

## Usage

- Interact with the system through the terminal.
- Follow the on-screen prompts to manage books and users, and perform borrowing/returning operations.

## Project Structure

```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request