# Golang microservices assignment

## Deploy

 So that we can test our services, a Makefile was created to run commands of the Dockerfile.

 Therefore, to raise the application it's necessary to run the follow command:

`make run`

 And to finish the application we can run the follow command:

`make stop`

## Endpoints

### Add 
 
 Endpoint to read the json file and save data into the database.

 * **URL:** 
 localhost:8080/add

 * **Method:** 
 `GET`

 * **Success Response:** 
 Code: 200

### Get Data 
 
 Endpoint to return a list of data into the database.

 * **URL:** 
 localhost:8080/data

 * **Method:** 
 `GET`

 * **Success Response:** 
 Code: 200
 Content: 
 ```json
 {
    "data": [
        {
            "name": "Hebei",
            "city": "Hebei",
            "country": "China",
            "alias": [
                ""
            ],
            "coordinates": [
                115.27,
                39.88
            ],
            "province": "Hebei",
            "timezone": "Asia/Shanghai",
            "unlocs": [
                "CNHEB"
            ],
            "code": "57000",
            "regions": [
                ""
            ],
            "key": "CNHEB",
            "id": "48"
        },
    ]
 }
 ```

* **Param:**

    **Optional:**
    `id=[string]`