## TODO app

### Backend API
- POST /tasks/{id}
- GET /tasks
- GET /tasks/{id}
- PUT /tasks/{id}
- DELETE /tasks/{id}

The API will perform CRUD operations on a `Task` object.

### DB
`Task` object schema:
{
    id: string,
    description: string,
    isComplete: boolean,
    createdAt: ISO-8601 timestamp,
    updatedAt: ISO-8601 timestamp,
}

### Frontend
Build a very simple HTML page that's divided in two vertical sections:
Section 1: Form to add a TODO
Section 2 is a table which lists all the TODOs.