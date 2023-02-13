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

Here is an example of defing a request for a mock.

```
    "request" : {
        "method": "POST",
        "path": {
            "pattern": "/test/handler"    
        }
    },
```