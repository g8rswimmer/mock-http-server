# mock-http-server
This is a mock HTTP server that can be used for testing of downstream services.

It is a simple mock with new features being added in the future.

## Release
Please refer to [changle log](change_log.md).

This is currently in BETA while being developed.

## Enviromental Variables
| Variable              | Required | Default | Description                                    |
|-----------------------|----------|---------|------------------------------------------------|
| MOCK_HTTP_SRVR_PORT   | N        | 8080    | The HTTP server port to handle requests        |
| MOCK_HTTP_HANDLER_DIR | Y        | None    | The directory to load all of the mock handlers |

## Defining Mock Endpoints
Please refer to the documentation (here)[mock_handler.md] for more information on how to define mock endpoints.

## Running
A makefile is provide to run the server in 'default' mode.
| Variable              | Value           |
|-----------------------|-----------------|
| MOCK_HTTP_SRVR_PORT   | 8080            |
| MOCK_HTTP_HANDLER_DIR | _mock_handlers_ |

```
make run-default
```

The command line will display
```
2023/02/12 19:07:06 Starting Mock HTTP Server...
2023/02/12 19:07:06 Config Vars...
2023/02/12 19:07:06 {
    "Server": {
        "Port": "8080"
    },
    "Handler": {
        "Directory": "_mock_handlers"
    }
}
2023/02/12 19:07:06 Happy Mocking...
```