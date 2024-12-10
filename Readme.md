# To-Do List API in Golang

## Overview

This project is a RESTful API built in Go for managing tasks. The API supports operations like creating tasks, retrieving tasks, caching frequent requests using Go routines and channels, and interacting with AWS DynamoDB for persistent storage.

The API is deployed on AWS Elastic Beanstalk and is accessible via the following URL:

[http://backend-todo-golang.us-east-1.elasticbeanstalk.com/](http://backend-todo-golang.us-east-1.elasticbeanstalk.com/)

You can use this URL to interact with the API for creating, retrieving tasks, and more.

## Features

- **Create a Task**  
  - Endpoint: `POST /api/tasks`  
  - Body:  
    ```json
    { "title": "Task Title", "description": "Task Description" }
    ```  
  - Response: `201 Created` with the created task object.

- **Get All Tasks**  
  - Endpoint: `GET /api/tasks`  
  - Response: List of tasks.

- **Get Task by ID**  
  - Endpoint: `GET /api/tasks/:id`  
  - Caches tasks for improved performance and fetches from DynamoDB if not cached.

## Architecture

- **Cache Management**: In-memory cache using Go routines and channels to store and fetch a task.
- **DynamoDB**: Persistent storage for tasks.
- **Concurrency**: Go routines and channels are used for efficient data fetching and handling concurrent requests.
- **AWS Elastic Beanstalk**: Deployment platform for the application.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/manuka99/backend-todo.git
   cd backend-todo
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up DynamoDB on AWS, create a table to store tasks.

4. Run the server:
   ```bash
   go run application.go || air
   ```

5. Test the API with `curl` or Postman:
   - **GET /api/tasks** to retrieve all tasks.
   - **GET /api/tasks/:id** to retrieve a task by ID.
   - **POST /api/tasks** to create tasks.

## AWS Elastic Beanstalk Deployment & Automation

This project is integrated with **AWS Elastic Beanstalk** (EB) for continuous deployment, using GitHub Actions for automation.

---

### Deployment Automation

A GitHub Actions workflow is set up to automatically deploy the application to EB whenever changes are pushed to the `main` branch.

- **Workflow File**: `.github/workflows/deploy.yml`
- **Triggers**: The deployment is triggered by `push` events on the `main` branch or manually using the `workflow_dispatch` event.
- **Steps**: 
  - The source code is checked out.
  - A deployment package (`deploy.zip`) is created, excluding `.git` files.
  - The application is deployed to Elastic Beanstalk using the `einaregilsson/beanstalk-deploy` action.

### Secrets Configuration

Ensure the following secrets are configured in the GitHub repository:
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `EB_APPLICATION_NAME`
- `EB_ENVIRONMENT_NAME`
- `AWS_REGION`

Variables
- `APP_VERSION` - The version of the application to deploy.

These secrets & variables are used in the deployment process to authenticate and deploy the application to AWS.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.