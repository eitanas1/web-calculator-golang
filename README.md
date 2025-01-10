
# Web-calculator | [English](README.md) | [Русский](README.ru.md)

web-calculator is a web service that allows users to send arithmetic expressions via HTTP and receive their results.

## Functionality

- Supports addition, subtraction, multiplication, division operations, and expressions in parentheses
- Expressions can be entered with or without spaces between numbers and operands
- The calculator accepts positive integers as input

## Requirements

- Go version 1.22 or newer

## Installation

1. Clone the repository

```bash
git clone https://github.com/bulbosaur/web-calculator-golang
```

2. Run the server from the project repository
   
``` bash
go run cmd/main.go
```

## Environment Variables

| Variable | Description | Default Value |
|----------|----------------|----------------|
| PORT | Server port | 8080 |
| HOST | Server host | localhost |

To change PORT and HOST values, you need to set them manually in the command line / terminal before starting the server.

### Windows

```bash
set PORT=3000
set HOST=0.0.0.0
```

Check set variables:

```bash
echo %PORT%
echo %HOST%
```

### Linux / MacOS

```bash
export PORT=3000
export HOST=0.0.0.0
```

Check set variables:

``` bash
echo $PORT
echo $HOST
```

## API

Default base URL: http://localhost:8080

| API endpoint | Method | Request Body | Server Response | Response Code |
|--------------|--------|--------------|-----------------|---------------|
| /api/v1/calculate | POST | {"expression": "2 * 2"} | {"result":"4"} | 200 |
| /api/v1/calculate | POST | "expression": "2 * 2" | {"error":"Bad request","error_message":"invalid request body"} | 400 |
| /api/v1/calculate | GET | {"expression": "2 * 2"} | Method Not Allowed | 405 |
| /coffee | | | I'm a teapot | 418 |
| /api/v1/tea | | | 404 page not found | 404 |

### Response Codes

- 200 - Successful request
- 400 - Bad request
- 404 - Resource not found
- 405 - Method not allowed
- 422 - Invalid expression (e.g., English letter instead of a number)
- 500 - Internal server error

### Usage Examples

For sending POST requests, it's most convenient to use [Postman](https://www.postman.com/downloads/).

1. StatusOK 200

```bash
curl 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "42 + 5 * 2"
}'

# {"result":52}
```

```bash
curl 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "6-8"
}'

# {"result":-2}
```

```bash
curl 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "123(3/2)"
}'

# {"result":184.5}
```

2. Bad Request 400

```bash
curl 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 * 2
}'

# {"error":"Bad request","error_message":"invalid request body"}
```

3. Unprocessable Entity 422

```bash
curl 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "cat + 100500"
}'

# {"error":"Expression is not valid","error_message":"invalid characters in expression"}
```

```bash
curl 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "()"
}'

# {"error":"Expression is not valid","error_message":"the brackets are empty"}
```

```bash
curl 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1/(2 - 3 + 1)"
}'

# {"error":"Expression is not valid","error_message":"division by zero is not allowed"}
```

```bash
curl 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1 000 000 + 6"
}'

# {"error":"Expression is not valid","error_message":"missing operand"}
```