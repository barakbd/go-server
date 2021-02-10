# go-server

## Endpoints

Single endpoint - **/replayed**

### POST

Request:

    {
        name: [STR, UNIQUE]
        age: INT
    }

Response:

    {
        id: UUID
    }

### GET

Request:

    **/replayed/:id**

Response
    {
        id: UUID
    }
>>>>>>> inital commit
