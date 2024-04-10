# final-proj-markas

This final project is the last project from the Beta intensive program.

## How to run this code

1. Run `go mod tidy` to install all dependencies.
2. create .env file ad fill the key by the .env.example file
3. Run `go run main.go` to run the program and you will get the endpoint link http://127.0.0.1:3000

now u can tes the http://127.0.0.1:3000 at postman

1. GET `/`
* Middleware: Authorization (fill the value with your secret key)
* Response:
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

2. GET `/{database_name}`
* Middleware: Authorization (fill the value with your secret key)
* Response:
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
    "message": "success"
}

3. POST `/{db_name}`
* Middleware: Authorization (fill the value with your secret key)
* Body Request form-data: 
key1 type text: database_name 
key2 type file: file_name

value1: db_name
value2: file.zip

* Response:
{
    "data": {
        "database_name": "db_1",
        "file_name": "heheh.zip",
        "id": 39,
        "timestamp": "2024-04-11T02:18:22.754+08:00"
    },
    "message": "success"
}

4. GET `/{id}/{download}`
* Middleware: Authorization (fill the value with your secret key)
* Response:
{
    "message": "File successfully retrieved and saved to the specified download path"
}

