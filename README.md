# Okul - School Listing Project

Okul is a web application that allows users to create an account, receive email confirmations, list schools, and add them to their favorites. The backend of the application is built using Golang and Fiber for the server, and it uses Postgres as the database. The frontend is developed with React.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Starting the Server](#starting-the-server)
  - [Running the Frontend](#running-the-frontend)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before you start, make sure you have the following installed on your system:

- [Golang](https://golang.org/dl/)
- [Node.js](https://nodejs.org/)
- [Postgres](https://www.postgresql.org/download/)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/k61b/okul.git
cd okul
```

2. Create a config.dev.yml configuration file in the config directory like config.prod.yml

3. Install dependencies:

```bash
# Backend (Golang/Fiber)
go mod tidy

# Frontend (React)
cd frontend
npm install
```

### Configuration
The project uses a YAML configuration file (config.dev.yml) for environment-specific settings. Make sure to provide appropriate values for the keys in this file, especially the database connection details, JWT secret, and email credentials.

## Usage

### Starting the Server

1. Start the Postgres database.

2. Migrate the database schema:
```bash
# Manually run the migration file
```

3. Start the Golang server:
```bash
make run
```

### Running the Frontend

In a separate terminal window:
```bash
cd client
npm run dev
```

The frontend will be accessible at http://localhost:3000.

## Features
- User authentication with email confirmation.
- Listing schools with details.
- Adding schools to favorites.
- Responsive design for both desktop and mobile.

## Contributing
Contributions are welcome! Feel free to open an issue or create a pull request.

## License

[MIT](LICENSE)
