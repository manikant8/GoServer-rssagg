# GoServer-rssagg

## Overview

GoServer-rssagg is a web server built with Go that fetches and aggregates RSS feeds. It allows users to follow their favorite feeds and retrieves the latest posts, making it easier to stay updated with the content they love.

## Features

- Fetches RSS feeds from various sources
- Aggregates and displays posts in a user-friendly manner
- Supports multiple users with a follow feature for feeds
- Implements efficient database management using SQL

## Technologies Used

- **Go**: Programming language used to build the server.
- **PostgreSQL**: Database used for storing feed and post information.
- **Goroutines**: For concurrent fetching of feeds.
- **Git**: Version control for managing the source code.

## Getting Started

### Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (version 1.17 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- Git (for version control)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/manikant8/GoServer-rssagg.git
   cd GoServer-rssagg
2. Set up the PostgreSQL database:
    1.Create a new database and user for the project.
    2.Update the database connection settings in the project as needed.
3. Install the required Go modules: go mod tidy
4. Running the Server : To start the server, use the following command: go run main.go

API Endpoints
    GET /feeds: Retrieve a list of all feeds.
    POST /feeds: Add a new feed to track.
    GET /posts: Retrieve posts from followed feeds.


Contributions are welcome! Please follow these steps:

Fork the repository.
Create your feature branch (git checkout -b feature/YourFeature).
Commit your changes (git commit -m 'Add some feature').
Push to the branch (git push origin feature/YourFeature).
Open a pull request.
