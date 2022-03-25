# Visits service project

### Prerequisites

`go version >= go1.17`

### Server
upon startup
server is hosted on `localhost` and is available on port
`3000`

e.g. `http://localhost:3000`

### Endpoints

`GET /v1/health`

params:

`health: string representing state of service for monitoring`

```
Status: 200 OK

Body:

{
health: "healthy"
}
```

`GET /v1/visits`

params:

`date_as_string: string. represents the date in dd:mm:yyyy format`

`bathroom_visit_count: float. represents the total visits for a date`

`lin_reg_trend: float. represents the y axis value for a particular date satisfying the linear regression
equation y = a(x)+b`

`above_trend: bool. true | false given bathroom_visit_count > lin_reg_trend 
will allow client to colour code based on above or below`

`equation: string. y = a(x)+b trend line equation for Any(x,y) - same for each object`
```
Status : 200 OK

Body:

{
    "data": [
        {
            "date_as_string": "13:1:2021",
            "bathroom_visit_count": 1,
            "lin_reg_trend": 2.8577397360518138,
            "above_trend": false,
            "equation": "y = -8.99199391605682e-09(x) + 2.8695374209015525"
        },
        {
            "date_as_string": "1:12:2021",
            "bathroom_visit_count": 1,
            "lin_reg_trend": 2.8594482148958646,
            "above_trend": false,
            "equation": "y = -8.99199391605682e-09(x) + 2.8695374209015525"
        },
        {
            "date_as_string": "14:1:2022",
            "bathroom_visit_count": 4,
            "lin_reg_trend": 2.856840527668214,
            "above_trend": true,
            "equation": "y = -8.99199391605682e-09(x) + 2.8695374209015525"
        },
        {
            "date_as_string": "20:1:2022",
            "bathroom_visit_count": 2,
            "lin_reg_trend": 2.85144533131858,
            "above_trend": false,
            "equation": "y = -8.99199391605682e-09(x) + 2.8695374209015525"
        },
        {
            "date_as_string": "28:5:2021",
            "bathroom_visit_count": 1,
            "lin_reg_trend": 2.843892065421086,
            "above_trend": false,
            "equation": "y = -8.99199391605682e-09(x) + 2.8695374209015525"
        },
        {
            "date_as_string": "21:5:2021",
            "bathroom_visit_count": 3,
            "lin_reg_trend": 2.850186461162326,
            "above_trend": true,
            "equation": "y = -8.99199391605682e-09(x) + 2.8695374209015525"
        }
}
```

### Running project
start the server with:
`$ go run cmd/main.go`

run tests:
`$ go test ./... -v`



### Structure

`cmd` - contains `main.go` file

`internal/filehandler` - package responsible for 
handling file and calculating trend

`internal/service` - package responsible for starting
and running server

#### Tests
filehandler: `internal/filehandler/test`

service: `internal/service/test`
