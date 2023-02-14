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
    * Varaibles - this is where a variable can be used to validate a part of the path.  The variable is define by `{label}`.
         * Example: `/my/variable/path/{varaible}`

Here is an example of defining a request for a mock.

```
    "request" : {
        "method": "POST",
        "path": {
            "pattern": "/test/handler"    
        }
    },
```

### Request Path Variable
Path variables can be used to validate a path part using `regex`, `valiation`, etc.  The path variable is defined between `{}`.  The list of variables define the label and the validation func that can be used.

Values are defined by `func:pattern` where the `pattern` will be supplied to the `func` for evaluation.

Supported validation functions:
* `reqex` - Regex patthen 
* `validator` - uses the go [validator](https://github.com/go-playground/validator) library to valid the pattern

#### Example
The following example will compare the `id` label with regex for a path part starting with `t` and ending with `ing`.
```
        "path": {
            "pattern": "/test/handler/{id}",
            "variables": [
                {
                    "label": "id",
                    "value": "regex:t([a-z]+)ing"
                }
            ]    
        }
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