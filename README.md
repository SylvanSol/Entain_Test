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
│  |  ├─ queries_test.go                            # ← Updated with Task 1, 2 & 3 tests        
│  ├─ proto/
│  |  ├─ racing/
│  |  |  ├─ racing.proto                            # ← Defines RaceStatus enum
│  |  ├─ racing.go           
│  ├─ service/                                      
│  |  ├─ racing.go                                  # ← Populates Race.Status
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

### Task 2: Ordering Support
- **Proto:** added `optional string order_by = 3` to `ListRacesRequestFilter`.  
- **DB repo:** `applyFilter` now appends `ORDER BY advertised_start_time` or custom column.

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
- **Tests:** new `TestListRaces_Status` in `db/queries_test.go` verifies CLOSED vs. OPEN.

### Task 4: GetRaceById RPC (Incomplete)  
- **Proto (service):** attempted to add:
  ```proto
  rpc GetRace(GetRaceRequest) returns (GetRaceResponse);
  ```  
  to `/racing/proto/racing.proto`.  
- **Error encountered:**
  ```
  Import "google/api/annotations.proto" was not found or had errors.
  Import "google/api/http.proto" was not found or had errors.
  ```  
  because the racing service proto should not include HTTP annotations.  
- **Current status:** Task 4 could not be completed in time due to these import errors and the elapsed time since assignment.

---

## Testing

All implemented tests live in **racing/db/queries_test.go** (Tasks 1–3 covered). Task 4 tests are pending completion of the service stub.

Run:
```bash
go test ./racing/db
```

---

## Known Issues & Future Steps

- **Current Stuck Point:**  
  The primary issue (missing method `mustEmbedUnimplementedRacingServer`) has been addressed by ensuring your service implementation exactly matches the generated interface. If the error persists, verify that only one version of the generated package is used throughout the project.
  
- **Future Improvements:**  
  - Add additional sorting options for the `ListRaces` RPC.
  - Implement further endpoints, such as fetching a single race by ID and creating a separate Sports service.
  - Improve unit and integration tests.
  
- **Task 4 Incomplete:**  
  GetRaceById RPC stubs in `racing/proto` fail to build due to inappropriate HTTP imports.  
- **Next steps:**  
  - Strip gateway annotations from service proto, regenerate, then implement the repo and service methods.  
  - Add corresponding unit tests.  
  - Implement additional RPCs (e.g. GetRaceById).

---

## Contact

If you have any questions or require further clarification, please feel free to reach out.

---