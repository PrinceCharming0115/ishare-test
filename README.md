# Secure REST API with GO

## Objective
Develop a secure REST API in Go that utilizes JSON Web Signatures (JWS) for signing and verifying tokens and implements OAuth 2.0 for authorization.

## Description
You are required to build a RESTful API for a simple task management application. The API should have endpoints for creating, retrieving, updating, and deleting tasks. The access to these endpoints should be secured using OAuth 2.0, and the tokens should be signed using JSON Web Signatures (JWS).

## Requirements
1. ### API Endpoints
    - **Create Task**: `POST /tasks`
    - **Get Task**: `GET /tasks/{id}`
    - **Update Task**: `PUT /tasks/{id}`
    - **Delete Task**: `DELETE /tasks/{id}`
    - **List Tasks**: `GET /tasks`
2. ### Task Model 
    - `ID` (UUID)
    - `Title`
    - `Description`
    - `Status` (e.g., pending, completed)
    - `CreatedAt`
    - `UpdatedAt`
3. ### Authentication and Authorization
    - Implement OAuth 2.0 Authorization Code Flow.
    - Secure all endpoints except the OAuth token endpoint.
    - Use JWS to sign and verify tokens.
4. ### Technology Stack
    - Go (Golang)
    - Any preferred Go web framework
    - Use any suitable database

## Deliverables
- Source code in a public Git repository (e.g., GitHub).
- README file with setup instructions, explanations, and any assumptions made.
- API documentation (Swagger or similar).

## Run Application
## Run Application
1. Install docker.
2. Copy `backend/.env.local` and rename it to `.env`.
3. Start docker container.

    ```bash
    make up
    ```

    If it doesn't work, use this command.

    ```bash
    docker compose up
    ```

4. Stop docker container.

    ```bash
    make down
    ```

    If it doesn't work, use this comand

    ```bash
    docker compose down
    ```
5. You can see swagger documentation at this url.

    <http://127.0.0.1:8000/docs/index.html>

6. You can get OAuth2.0 access token from this url.

    <http://127.0.0.1:8000/docs/index.html>

### Note:
If you make assumptions, please note them down in documentation.