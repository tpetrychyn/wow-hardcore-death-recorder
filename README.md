## WoW Hardcore Player Death Recorder

## Repo Structure
This repository is structured as a monorepo. Microservices are contained under `/cmd`. Each microservice contains its own integration tests.
`/internal` directory contains shared code for standing up microservices and relevant business logic and models.

## API

#### Endpoints

#### Response Codes:
- 200 - OK
- 400 - bad request
- 429 - rate limited
- 500 - internal server error

### GET `/deaths/:guid`
Returns the list of deaths recorded on that GUID. Only trust entries where the player name matches the guid.

#### Request Example:
```shell
curl --location 'localhost:8080/deaths/Player-5139-020C481D'
```


#### Response:
```json
[
    {
        "Sender": "Zunter",
        "SourceID": "435",
        "RecordedAt": "2023-04-20T11:39:41Z",
        "GUID": "Player-5139-020C481D"
    },
    {
        "Sender": "Zunter",
        "SourceID": "435",
        "RecordedAt": "2023-04-20T11:39:55Z",
        "GUID": "Player-5139-020C481D"
    }
]
```
### POST `/deaths`
Insert a list of recorded deaths from the Wow_Hardcore deathlog addon

#### Request Example:
```shell
curl --location 'localhost:8080/deaths' \
--header 'Content-Type: text/plain' \
--data 'Recorded_Deaths = {
	{
		["sender"] = "Zunter",
		["source_id"] = "435",
		["time"] = 1682015980,
		["guid"] = "Player-5139-020C481D",
	}, -- [1]
	{
		["source_id"] = "435",
		["sender"] = "Zunter",
		["time"] = 1682015995,
		["guid"] = "Player-5139-020C481D",
	}, -- [2]
	{
		["source_id"] = "435",
		["sender"] = "Zunter",
		["time"] = 1682016033,
		["guid"] = "Player-5139-020C481D",
	}, -- [3]
}'
```

#### Response:
HTTP status code 200


### GET `/health`
Returns status 200 OK when API is healthy.
