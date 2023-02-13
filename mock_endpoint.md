# Defining Mock Endpoints
Mock Endpoints are defined by JSON files.  This file define the request and response for a mock.

The handler definition contains:
* HTTP Request - this will define the path, method, etc. of the mock
* HTTP Response - this will define the statuc code, response body, etc. of the mock

### Example
```
{
    "request" : {
        "method": "POST",
        "path": {
            "pattern": "/test/handler"    
        }
    },
    "response": {
        "status_code": 201,
        "body": {
            "message": "post handler was executed"
        }
    }
}
```
## Request
The mock request supports
* Request method defines `GET`, `POST`, etc.
* Path like `/this/is/my/path

Here is an example of defining a request for a mock.

```
    "request" : {
        "method": "POST",
        "path": {
            "pattern": "/test/handler"    
        }
    },
```

## Response
The mock response supports
* Response status code, `200`, `400`, etc.
* Reponse body as a JSON object

Here is an exmaple of defining a response for a mock.

```
    "response": {
        "status_code": 200,
        "body": {
            "message": "get handler was executed"
        }
    }
```