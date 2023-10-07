# How to run

1. make sure that Go 1.20+, Docker, makefile is installed on the machine
2. navigate to root of the project and run `go mod tidy`
3. run `make run`

# event create endpoint

## call `http://localhost:1323/v1/event` endpoint with test data

```
{
"name": "My Event",
"date": "2023-04-20T14:00:00Z",
"languages": ["English", "French"],
"videoQuality": ["720p", "1080p"],
"audioQuality": ["High", "Low"],
"invitees": ["example1@gmail.com", "example2@gmail.com"],
"description": "A short description of the event"
}
```
