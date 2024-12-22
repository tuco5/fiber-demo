# Notes API - Go & Fiber Demo

A simple REST API built with Go and the Fiber framework that performs CRUD (Create, Read, Update, Delete) operations on a "note" resource.

This project demonstrates how to build a basic API with Go and Fiber, handling operations for notes, which include creating, reading, updating, and deleting notes.
Features

- Create a new note
- Read all notes or a single note by ID
- Update a note by ID
- Delete a note by ID

Technologies Used

- Go (Golang)
- Fiber (Fast and minimalist web framework for Go)
- PostgreSQL

## Getting Started
### Prerequisites
Before you begin, make sure you have Go installed on your machine.

1. Install Go
2. Clone the repository to your local machine:

        git clone https://github.com/yourusername/fiber-notes-api.git
        cd fiber-notes-api

3. Install the necessary Go dependencies:

        go mod tidy

4. Set up the database:
- You can use SQLite, which is included in the project by default, or configure another database.
- If you are using SQLite, the database file will be created automatically when you first run the app.

## Running the Application

To run the server locally:

    go run server.go

The server will start on port :3000 by default. You should see something like:
_Server is running on http://localhost:3000_

## Test the API with tools like Postman or curl.

### API Endpoints

Create a Note
POST /api/v1/notes

Request Body:
    
    {
      "title": "Note Title",
      "content": "Note Content"
    }

Response:

    {
      "id": 1,
      "title": "Note Title",
      "content": "Note Content"
    }

Get All Notes
GET /api/v1/notes

    Response:

    [
      {
        "id": 1,
        "title": "Note Title 1",
        "content": "Note Content 1"
      },
      {
        "id": 2,
        "title": "Note Title 2",
        "content": "Note Content 2"
      }
    ]

Get a Single Note
GET /api/v1/notes/:id

    Response:

    {
      "id": 1,
      "title": "Note Title",
      "content": "Note Content"
    }

Update a Note
PUT /api/v1/notes/:id

    Request Body:

{
  "title": "Updated Note Title",
  "content": "Updated Note Content"
}

Response:

    {
      "id": 1,
      "title": "Updated Note Title",
      "content": "Updated Note Content"
    }

Delete a Note
DELETE /api/v1/notes/:id

    Response:

        {
          "message": "Note deleted successfully"
        }

### Testing the API

You can test the endpoints using any HTTP client like Postman or curl.
Example with curl:

Create a Note:

    curl -X POST http://localhost:3000/api/notes -d '{"title":"Note Title", "content":"Note Content"}' -H "Content-Type: application/json"

Get All Notes:

    curl http://localhost:3000/api/notes

Get a Single Note:

    curl http://localhost:3000/api/notes/1

Update a Note:

    curl -X PUT http://localhost:3000/api/notes/1 -d '{"title":"Updated Note", "content":"Updated Content"}' -H "Content-Type: application/json"

Delete a Note:

    curl -X DELETE http://localhost:3000/api/notes/1


## License

This project is licensed under the MIT License - see the LICENSE file for details.
