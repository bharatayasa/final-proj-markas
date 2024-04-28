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
                 "database_name": "db_2",
                 "latest_backup": {
                     "file_name": "hahaahahs.zip",
                     "id": 2,
                     "timestamp": "2024-04-06T14:10:53.94+08:00"
                 }
             },
             {
                 "database_name": "db_5",
                 "latest_backup": {
                     "file_name": "mysql-2023-10-29-00-00-00-cv_kucing_oren-8634bf3f-23b5-45a7-8b78-fe9b1a3bcf66.sql.zip",
                     "id": 12,
                     "timestamp": "2024-04-08T06:05:20.031+08:00"
                 }
             }
         ],
         "message": "Success"
     }
     ```

2. **GET `/{database_name}`**

   - Middleware: bearer token (Fill the value with your secret key)
   - Response:
     ```json
     {
         "data": {
             "database_name": "db_1",
             "histories": [
                 {
                     "file_name": "tes copy.zip",
                     "id": 23,
                     "timestamp": "2024-04-08T15:14:58.521+08:00"
                 },
                 {
                     "file_name": "Screenshot 2024-03-29 at 19.25.59.png",
                     "id": 24,
                     "timestamp": "2024-04-08T18:58:31.703+08:00"
                 }
             ]
         },
         "message": "Success"
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
     {
         "data": {
             "database_name": "db_1",
             "file_name": "heheh.zip",
             "id": 39,
             "timestamp": "2024-04-11T02:18:22.754+08:00"
         },
         "message": "Success"
     }
     ```

4. **GET `/{id}/{download}`**

   - Middleware: bearer token (Fill the value with your secret key)
   - Response:
     ```json
     {
         "message": "File successfully retrieved and saved to the specified download path"
     }
     ```