#osaka 
## REST API for Posts

This is a simple CRUD (Create, Read, Update, Delete) API for managing posts. The API uses the Gin framework for routing and PostgreSQL as the database. Below is the documentation for the API, including examples for each CRUD operation.

### Base URL

`http://localhost:8080`

### Post Model

```go
type Post struct {
	ID    uint   `json:"id"`
	Text  string `json:"text"`
	Title string `json:"title"`
}
```

### Endpoints

#### 1. Get All Posts

- **Endpoint:** `GET /Posts`
- **Description:** Retrieve all posts from the database.
- **Example:**
  ```http
  GET http://localhost:8080/Posts
  ```

#### 2. Get a Post by ID

- **Endpoint:** `GET /Posts/:id`
- **Description:** Retrieve a specific post by its ID.
- **Example:**
  ```http
  GET http://localhost:8080/Posts/1
  ```

#### 3. Create a New Post

- **Endpoint:** `POST /Posts`
- **Description:** Create a new post with the provided JSON payload.
- **Example:**

  ```
  POST http://localhost:8080/Posts

  Request Body:
  {
    "text": "Content of the post",
    "title": "Title of the post"
  }
  ```

#### 4. Update a Post by ID

- **Endpoint:** `PUT /Posts/:id`
- **Description:** Update an existing post with the provided JSON payload.
- **Example:**
  ``` http
  PUT http://localhost:8080/Posts/1

  Request Body:
  {
    "text": "Updated content",
    "title": "Updated title"
  }
```

#### 5. Delete a Post by ID

- **Endpoint:** `DELETE /Posts/:id`
- **Description:** Delete a post by its ID.
- **Example:**
- 
```
  DELETE http://localhost:8080/Posts/1
```

### Response Formats

- **Success Response (Status Code: 200):**
  ```json
  {
    "id": 1,
    "text": "Content of the post",
    "title": "Title of the post"
  }
  ```

- **Error Response (Status Code: 404):**
  ```json
  {
    "error": "Post not found"
  }
  ```

- **Error Response (Status Code: 500):**
  ```json
  {
    "error": "Internal Server Error"
  }
  ```

### Note

- Make sure to replace `localhost:8080` with the appropriate base URL if the API is hosted elsewhere.
- For the `POST` and `PUT` requests, ensure the request body is a valid JSON payload as per the Post model structure.

This documentation provides a basic overview of the API. For detailed information about each endpoint and the expected request and response formats, refer to the code and comments in the provided Go program.
