# Entain Technical Test

## Overview

This repository is my solution for the Entain Technical Test. The project is structured as a multi-service Go application demonstrating:

- A basic REST gateway in the **api** directory.  
- A bare-bones racing service in the **racing** directory.

Recent modifications include:

1. **“Visible only”** filter on `ListRaces`.  
2. **Ordering** support via an `order_by` field.  
3. **Derived `status`** on each `Race` (OPEN/CLOSED based on time).  
4. **Attempted “GetRaceById”** RPC—currently incomplete due to import errors and time constraints.
5. To be made

---

## Project Structure

```
Entain_Test/
├─ api/                   
│  ├─ proto/
│  ├─ go.mod
│  ├─ main.go
│  ├─ tools.go
├─ racing/                
│  ├─ db/
│  |  ├─ db.go
│  |  ├─ queries.go
│  |  ├─ queries_test.go                            # ← Testing for all tasks
│  ├─ proto/
│  |  ├─ racing/
│  |  |  ├─ racing.proto                            # ← Service proto (no HTTP imports)
│  |  ├─ racing.go           
│  ├─ service/                                      
│  |  ├─ racing.go                                  # ← Implements status
│  ├─ go.mod
│  ├─ main.go
│  ├─ tools.go            
├─ example.png                                      # The example PNG for the test
├─ givenREADME.md                                   # The example README as instructions for the test      
├─ README.md                                        # This file (updated)
```

---

## Getting Started

### Prerequisites

- **Go ≥ 1.24.1**  
- **protoc** with `protoc-gen-go` & `protoc-gen-go-grpc` plugins  
- **SQLite3** driver (`github.com/mattn/go-sqlite3`)

## Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/SylvanSol/Entain_Test.git
   cd Entain_Test
   ```

2. **Generate Protos**  
   *Racing service:*
   ```powershell
   & 'C:\ProgramData\chocolatey\bin\protoc.exe' -I racing/proto \
     --go_out=racing/proto --go_opt paths=source_relative \
     --go-grpc_out=racing/proto --go-grpc_opt paths=source_relative \
     racing/proto/racing.proto
   ```
   *API service:*
   ```powershell
   & 'C:\ProgramData\chocolatey\bin\protoc.exe' -I api/proto/google/api -I api/proto/racing \
     --go_out=api/proto/racing --go_opt paths=source_relative \
     --go-grpc_out=api/proto/racing --go-grpc_opt paths=source_relative \
     --grpc-gateway_out=api/proto/racing --grpc-gateway_opt paths=source_relative \
     api/proto/racing/racing.proto
   ```

3. **Build & Run**  
   *Racing:*
   ```powershell
   $env:CGO_ENABLED=1
   $env:PATH += ";C:\ProgramData\mingw64\mingw64\bin"
   cd racing && go build && ./racing
   ```
   *API:*
   ```bash
   cd api && go build && ./api
   ```

---

## Recent Changes

### Task 1: “Visible Only” Filter
- **Proto:** added `bool only_visible = 2` to `ListRacesRequestFilter`.  
- **DB repo:** `applyFilter` appends `visible = 1` when `only_visible` is true.

#### Example Request
```bash
curl -X POST http://localhost:8000/v1/list-races \
  -H 'Content-Type: application/json' \
  -d '{
    "filter": {
      "only_visible": true
    }
  }'
  ```
### Task 2: Ordering Support
- **Proto:** added `optional string order_by = 3` to `ListRacesRequestFilter`.  
- **DB repo:** `applyFilter` now appends `ORDER BY advertised_start_time` or custom column.

#### Example Request
```bash
curl -X POST http://localhost:8000/v1/list-races \
  -H 'Content-Type: application/json' \
  -d '{
    "filter": {
      "only_visible": true
    },
    "sort": {
      "field": "name",
      "direction": "desc"
     }
}
```
### Task 3: Derived `status` Field
- **Proto:** added
  ```proto
  enum RaceStatus { UNSPECIFIED=0; OPEN=1; CLOSED=2; }
  message Race { … RaceStatus status = 7; }
  ```
- **Service:** in `service/racing.go`, after scanning `advertised_start_time`, set
  ```go
  if advertisedStart.Before(time.Now()) {
    race.Status = racing.RaceStatus_CLOSED
  } else {
    race.Status = racing.RaceStatus_OPEN
  }
  ```
#### Example Curl

```bash
curl -X POST http://localhost:8000/v1/list-races \
  -H 'Content-Type: application/json' \
  -d '{
    "filter": {}
  }'
```
#### Example Response

```json
{
  "races": [
    {
      "id": 301,
      "name": "Past Race",
      "advertised_start_time": "2024-01-01T12:00:00Z",
      "status": "CLOSED"
    },
    {
      "id": 302,
      "name": "Future Race",
      "advertised_start_time": "2026-01-01T12:00:00Z",
      "status": "OPEN"
    }
  ]
}
```

### Task 4: GetRaceById RPC
* **Proto:** Added the following to `racing.proto`:

  ```proto
  message GetRaceRequest {
    int64 id = 1;
  }

  message GetRaceResponse {
    Race race = 1;
  }

  rpc GetRace(GetRaceRequest) returns (GetRaceResponse);
  ```
* **Interface & Repository:**

  * Extended the `RacesRepo` interface with `GetByID(id int64) (*racing.Race, error)`.
  * Reused the existing `GetByID` implementation in `races.go`.
* **Service Layer:**

  * Implemented `GetRace` in `racing/service/racing.go`, which returns:

    * `NOT_FOUND` if the race doesn’t exist.
    * `INTERNAL` on DB error.
* **Tests:**

  * Added `TestGetByID` in `queries_test.go` to verify:

    * Correct field values.
    * Derived `status` is set to `OPEN` or `CLOSED`.
* **Note:**

  * This RPC fetches a single race by ID.
  * Status is always derived at runtime and returned with the race.

#### Example Request (via curl)

```bash
curl -X POST http://localhost:8000/v1/get-race \
  -H "Content-Type: application/json" \
  -d '{
    "id": 500
  }'
```

#### Example Response

```json
{
  "race": {
    "id": 500,
    "name": "Solo Race",
    "advertised_start_time": "2025-01-01T12:00:00Z",
    "status": "OPEN"
  }
}
```

### Task 5: Sports Service with `ListEvents` RPC

* **New Service:**
  Created a new standalone gRPC service called `sports`, with its own proto, service logic, and `main.go` file.

* **Directory Structure:**

  ```
  sports/
  ├── proto/
  │   └── sports/
  │       └── sports.proto
  ├── service/
  │   └── sports.go
  ├── main.go
  └── go.mod
  ```

* **Proto Definition:**
  `sports.proto` includes:

  ```proto
  message Event {
    int64 id = 1;
    string name = 2;
    string location = 3;
    google.protobuf.Timestamp advertised_start_time = 4;
  }
  ```

  Along with:

  * `ListEventsRequest`
  * `ListEventsResponse`
  * `rpc ListEvents`

* **Service Logic:**
  `ListEvents` returns 3 mocked sports matches with future start times. Each includes a name, location, and advertised start time.

* **gRPC Server:**
  Sports service runs on `localhost:9100`. The API service does not forward to this (no REST setup required).

* **Tests:**
  A unit test `TestListEvents_ReturnsMockEvents` in `sports/service/sports_test.go` verifies:

  * No errors on call
  * 3 expected events returned
  * Correct fields and ordering

* **How to Run:**

  ```bash
  cd sports
  go build && ./sports
  ```

* **How to Call (with grpcurl):**

  ```bash
  grpcurl -plaintext localhost:9100 sports.Sports/ListEvents
  ```

* **Sample Response:**

  ```json
  {
    "events": [
      {
        "id": "1",
        "name": "Red Hawks vs Blue Titans",
        "location": "Thunder Dome",
        "advertisedStartTime": "2025-05-09T05:37:58Z"
      },
      ...
    ]
  }
  ```

## Testing

All implemented tests live in **racing/db/queries_test.go** or **sports/service/sports_test.go**

Run:
```bash
go test ./racing/db
```

## Contact

If you have any questions or require further clarification, please feel free to reach out.

---