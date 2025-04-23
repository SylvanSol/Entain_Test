# Entain Technical Test

## Overview

This repository is my solution for the Entain Technical Test. The project is structured as a multi-service Go application demonstrating:
  
- A basic REST gateway in the **api** directory.
- A bare-bones racing service in the **racing** directory.

Recent modifications include adding a new filter to the `ListRaces` RPC to allow clients to request "visible only" races.

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

---

## Known Issues & Future Steps

- **Current Stuck Point:**  
  The primary issue (missing method `mustEmbedUnimplementedRacingServer`) has been addressed by ensuring your service implementation exactly matches the generated interface. If the error persists, verify that only one version of the generated package is used throughout the project.
  
- **Future Improvements:**  
  - Add additional sorting options for the `ListRaces` RPC.
  - Implement further endpoints, such as fetching a single race by ID and creating a separate Sports service.
  - Improve unit and integration tests.
  
---

## Testing

All tests live in **racing/db/queries_test.go**. They cover:

1. Visible-only filtering.  
2. Default and custom ordering.  
3. Derived OPEN/CLOSED status.

Run:
```bash
go test ./racing/db
```

---

## Contact

If you have any questions or require further clarification, please feel free to reach out.

---

## Bibliography

- [gRPC-Go Documentation – Official Site](https://pkg.go.dev/google.golang.org/grpc) citeturn7file4
- [Protocol Buffers Documentation](https://developers.google.com/protocol-buffers) citeturn7file4
- [gRPC-Gateway Documentation](https://grpc-ecosystem.github.io/grpc-gateway/) citeturn7file4

---