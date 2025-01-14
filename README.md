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
- **Description**: Cancels a specific backfill task by its ID. Pass id -1 to stop system backfill task (backfill all exchanges).
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
- **Description**: Restarts a specific backfill task by its ID. Pass id -1 to run system backfill task (backfill all exchanges).
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

## Broadcast Server API Documentation

This server serve the endpoints to subscribe trade logs
### **WebSocket Connection Creation**

- **URL**: `/eventlogws`
- **Method**: `GET`
- **Description**: Connect WebSocket.
- **Request Params**:
    - `id` (string, optional): The ID of websocket connection, also the consumer group ID. 
    Empty if you need to create new connection and subscribe from the newest offset. 
    Else, pass the ID to continue from the most recently event.
    - `maker` (string, optional): The maker you want to subscribe.
    - `taker` (string, optional): The taker you want to subscribe.
    - `maker_token` (string, optional): The maker token you want to subscribe.
    - `taker_token` (string, optional): The taker token you want to subscribe.
    - `event_hash` (string, optional): The event hash you want to subscribe.

  The filter will combine non-empty fields (AND). If no field are pass, all the trade logs will be sent.
- **Response**:
    - **ID Message**: With new connection, the first message will be the ID of session.
      ```json
      {
        "id": "<session_id>"
      }
      ```
    - **Trade log Message**: with type `trade_log` and the data containing trade log information
      ```json
      {
        "type": "trade_log",
        "data": {
          "order_hash": "0xcccf91f6cc3f636f0c7864ae8bcbc7cddbb54971997f90d81891a139c36c33e9",
          "maker": "0x51C72848c68a965f66FA7a88855F9f7784502a7F",
          "taker": "0x22F9dCF4647084d6C31b2765F6910cd85C178C18",
          "maker_token": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
          "taker_token": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
          "maker_token_amount": "20210895540598740",
          "taker_token_amount": "50995428",
          "contract_address": "0xDef1C0ded9bec7F1a1670819833240f027b25EfF",
          "block_number": 21034739,
          "tx_hash": "0x6926077a7cfdc78cd1649c43e85da918f10b4ee0c6e930d5107a50de0c9d0837",
          "log_index": 489,
          "timestamp": 1729764251000,
          "event_hash": "0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f",
          "expiration_date": 1729764301,
          "maker_token_price": 0,
          "taker_token_price": 0,
          "maker_usd_amount": 0,
          "taker_usd_amount": 0,
          "state": ""
        }
      }
      ```
    - **Revert Message**: with type `revert` and data containing the block numbers of reverted blocks
      ```json
      {
        "type": "revert",
        "data": [11, 12, 13]
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

## Price-filler
### Refetch token price 

- **URL**: `/price_filler/refetch`
- **Method**: `POST`
- **Description**: Reset taker_token_price, maker_token_price, taker_usd_amount, maker_usd_amount of trades with faulty token.
- **URL Parameters**:
    - `address` (string): The token address.
    - `exchange` (string): The exchange name.
    - `from_ts` (int): The starting timestamp (millisecond).
    - `to_ts` (int): The ending timestamp (millisecond).
- **Response**:
    - **200 OK**: Success.
      ```json
      {
        "data": {
            "token": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
            "exchange": "bebop",
            "from": 1734204504800,
            "to": 1734809304800,
            "number of row updated": 2000
        },
        "success": true
      }
      ```
    - **400 Bad Request**: If the from_ts or to_ts is invalid.
      ```json
      {
          "error": "<error_message>",
          "success": false
      }
      ```
    - **500 Internal Server Error**: If there is an error resetting token price.
      ```json
      {
        "error": "<error_message>",
        "success": false
      }
      ```