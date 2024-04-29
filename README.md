# Final Project Markas

This final project is the culmination of the Beta intensive program.

## How to Run the Code

1. Run `go mod tidy` to install all dependencies.
2. Create a `.env` file and fill the key as per the `.env.example` file.
3. Run `go run main.go` to execute the program. You will receive the endpoint link http://127.0.0.1:3000.

Now you can test using Postman with http://127.0.0.1:3000.

### Endpoints

1. **GET `/`**

   - Middleware: bearer token (Fill the value with your secret key)
   - Response:
     ```json
     {
        "data": [
            {
                "database_name": "information_schema",
                "latest_backup": {
                    "file_name": "msql-2024-04-29-14-33-28-information_schema-ed9a9059-86aa-4914-81b8-31a2c9c07cc1.sql.zip",
                    "id": 9,
                    "timestamp": "2024-04-29T14:33:28.328+08:00"
                }
            },
            {
                "database_name": "perpus_markas",
                "latest_backup": {
                    "file_name": "msql-2024-04-29-14-33-28-perpus_markas-089f5a41-1876-47bd-8199-7203953768e5.sql.zip",
                    "id": 10,
                    "timestamp": "2024-04-29T14:33:28.327+08:00"
                }
            },
        ],
        "message": "Success"
     }
     ```

2. **GET `/{database_name}`**

## Middleware: Bearer Token 
  - Fill the value with your secret key

## Response:
    ```json
        {
        "data": {
            "database_name": "information_schema",
            "histories": [
            {
                "file_name": "msql-2024-04-29-14-31-44-information_schema-df658c80-3973-4351-9165-05c207da7315.sql.zip",
                "id": 5,
                "timestamp": "2024-04-29T14:31:44.034+08:00"
            },
            {
                "file_name": "msql-2024-04-29-14-33-28-information_schema-ed9a9059-86aa-4914-81b8-31a2c9c07cc1.sql.zip",
                "id": 9,
                "timestamp": "2024-04-29T14:33:28.328+08:00"
            }
            ]
        },
        "message": "success"
        }
    ```


3. **POST `/{db_name}`**

   - Middleware: bearer token (Fill the value with your secret key)
   - Body Request form-data:
     - Key1 type text: `database_name`
     - Key2 type file: `file_name`
   - Example values:
     - Value1: `db_name`
     - Value2: `file.zip`
   - Response:
     ```json
            "data": {
                "database_name": "tomatify",
                "file_name": "msql-2024-04-29-14-13-48-sekolahbeta-ab6f2452-3ba8-45e6-bbb6-3e747dd4b82d.sql.zip",
                "id": 7,
                "timestamp": "2024-04-29T14:29:04.319797+08:00"
            },
            "message": "success"
     ```

4. **GET `/{id}/{download}`**

   - Middleware: bearer token (Fill the value with your secret key)
   - Response:
     ```json
            {
                "message": "File successfully retrieved and saved to the specified download path"
            }
     ```