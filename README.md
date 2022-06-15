# go-test-task

A simple CRUD site where you can add, edit and delete entries (products). Unusual in this task:

* Key-value store (built-in key-value store) - [Bolt](https://github.com/boltdb/bolt) will be used as the database.
* BoltDb will use [encoding/gob](https://pkg.go.dev/encoding/gob) as a data serializer instead of normal encoding/json.

## Endpoints for web pages:

- `/products`
- `/product/add`
- `/product/edit/{productId}`

## Endpoints for requests:

`/q/product-search-by-name`
- Request:
    ```json
    {
        "searchName": ""
    }
    ```
- Good Response:
    ```json
    {
        "status": "success",
        "product": {...}
    }
    ```
- Error Response:
    ```json
    {
        "status": "failure",
        "error": "..."
    }
    ```

## Endpoints for commands:

`/cmd/add-product`
- Request:
    ```json
    {
        "name": "Name",
        "price": "1000"
    }
    ```
- Good Response:
    ```json
    {
        "status": "success",
        "productId": 1 // id of the last created product
    }
    ```
- Error Response:
    ```json
    {
        "status": "failure",
        "error": "..."
    }
    ```

`/cmd/edit-product`
- Request:
    ```json
    {
        "id": 1,
        "name": "New name",
        "price": "1000"
    }
    ```
- Good Response:
    ```json
    {
        "status": "success"
    }
    ```
- Error Response:
    ```json
    {
        "status": "failure",
        "error": "..."
    }
    ```

`/cmd/delete-product`
- Request:
    ```json
    {
        "id": 1
    }
    ```
- Good Response:
    ```json
    {
        "status": "success"
    }
    ```
- Error Response:
    ```json
    {
        "status": "failure",
        "error": "..."
    }
    ```