# Tradelogs

This is service to indexing tradelogs from event

# Tradelogs V2

## Backfill Server API Documentation

This server serve the endpoints to manage backfill task: list, create, cancel, restart tasks
### 1. **Backfill Task Creation**

- **URL**: `/backfill`
- **Method**: `POST`
- **Description**: Creates a new backfill task.
- **Request Body**:
    - `from_block` (uint64, required): The starting block number.
    - `to_block` (uint64, required): The ending block number.
    - `exchange` (string, required): The exchange name.
- **Response**:
    - **200 OK**: On success.
      ```json
      {
        "success": true,
        "id": "<task_id>",
        "message": "<message>"
      }
      ```
    - **400 Bad Request**: If there is a validation error (e.g., missing fields, invalid exchange).
      ```json
      {
        "success": false,
        "error": "<error_message>"
      }
      ```
    - **500 Internal Server Error**: If there is an error during task creation.
      ```json
      {
        "success": false,
        "error": "<error_message>"
      }
      ```

### 2. **Get All Backfill Tasks**

- **URL**: `/backfill`
- **Method**: `GET`
- **Description**: Retrieves all backfill tasks.
- **Response**:
    - **200 OK**: On success. The task with id -1 is the service's default backfill flow.
      ```json
      {
        "success": true,
        "tasks": [
          {
            "id": -1,
            "exchange": "",
            "from_block": 20926953,
            "to_block": 20962657,
            "processed_block": 20962657,
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "status": "processing"
          },
          {
            "id": 1,
            "exchange": "zerox",
            "from_block": 20962657,
            "to_block": 20962658,
            "processed_block": 20962657,
            "created_at": "2024-10-14T09:07:01.059135Z",
            "updated_at": "2024-10-14T17:18:32.814065Z",
            "status": "done"
          }
        ]
      }
      ```
    - **500 Internal Server Error**: If there is an error retrieving the tasks.
      ```json
      {
        "success": false,
        "error": "<error_message>"
      }
      ```

### 3. **Get Backfill Task By ID**

- **URL**: `/backfill/:id`
- **Method**: `GET`
- **Description**: Retrieves a specific backfill task by its ID.
- **URL Parameters**:
    - `id` (int, required): The task ID.
- **Response**:
    - **200 OK**: On success.
      ```json
      {
        "success": true,
        "task": {
            "id": 1,
            "exchange": "zerox",
            "from_block": 20962657,
            "to_block": 20962658,
            "processed_block": 20962657,
            "created_at": "2024-10-14T09:07:01.059135Z",
            "updated_at": "2024-10-14T17:18:32.814065Z",
            "status": "done"
          }
      }
      ```
    - **400 Bad Request**: If the task ID is invalid.
      ```json
      {
        "success": false,
        "error": "invalid task id: <id>"
      }
      ```
    - **500 Internal Server Error**: If there is an error retrieving the task.
      ```json
      {
        "success": false,
        "error": "<error_message>"
      }
      ```

### 4. **Cancel Backfill Task**

- **URL**: `/backfill/cancel/:id`
- **Method**: `GET`
- **Description**: Cancels a specific backfill task by its ID.
- **URL Parameters**:
    - `id` (int, required): The task ID.
- **Response**:
    - **200 OK**: On success.
      ```json
      {
        "success": true
      }
      ```
    - **400 Bad Request**: If the task ID is invalid.
      ```json
      {
        "success": false,
        "error": "invalid task id: <error_message>"
      }
      ```
    - **500 Internal Server Error**: If there is an error canceling the task.
      ```json
      {
        "success": false,
        "error": "<error_message>"
      }
      ```

### 5. **Restart Backfill Task**

- **URL**: `/backfill/restart/:id`
- **Method**: `GET`
- **Description**: Restarts a specific backfill task by its ID.
- **URL Parameters**:
    - `id` (int, required): The task ID.
- **Response**:
    - **200 OK**: On success.
      ```json
      {
        "success": true
      }
      ```
    - **400 Bad Request**: If the task ID is invalid.
      ```json
      {
        "success": false,
        "error": "invalid task id: <error_message>"
      }
      ```
    - **500 Internal Server Error**: If there is an error restarting the task.
      ```json
      {
        "success": false,
        "error": "<error_message>"
      }
      ```

## Re-generate mock file

First, you need to install `mockery`

- https://vektra.github.io/mockery/latest/installation/
- https://github.com/vektra/mockery

Then you use the `mockery` to generate mock files from interface

```
mockery --dir=<interface directory> --name=<interface name> --output=<output dir> --structname=<name of output struct> --filename=<output filename>
```

Example: generate `Storage` interface to a struct name `MockStorage`

```
mockery --dir=v2/pkg/storage/state --name=Storage --output=v2/mocks/ --structname=MockState --filename=State.go
```

For more information `mockery --help`