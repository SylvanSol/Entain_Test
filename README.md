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
│  |  ├─ google/
│  |  |  ├─ api/
│  |  |  |  ├─ annotations.proto
│  |  |  |  ├─ http.proto
│  |  ├─ racing/
│  |  |  ├─ racing.pb.go
│  |  |  ├─ racing.pb.gw.go
│  |  |  ├─ racing.proto
│  |  |  ├─ racing_grpc.pb.go
│  |  ├─ api.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  ├─ tools.go
├─ racing/                
│  ├─ db/
│  |  ├─ db.go
│  |  ├─ queries.go
│  |  ├─ queries_test.go
│  |  ├─ races.go
│  |  ├─ racing.go                
│  ├─ proto/
│  |  ├─ racing/
│  |  |  ├─ racing.pb.go
│  |  |  ├─ racing.proto
│  |  |  ├─ racing_grpc.pb.go
│  |  ├─ racing.go           
│  ├─ service/
│  |  ├─ racing.go
│  ├─ go.mod
│  ├─ go.sum      
│  ├─ main.go
│  ├─ tools.go            
├─ example.png                                      # The example PNG for the test
├─ givenREADME.md                                   # The example README as instructions for the test      
├─ README.md                                        # This file (updated)
```

---

## Getting Started

### Prerequisites

- [Go (latest version)](https://golang.org/doc/install)  
  Windows:  
  ```bash
  choco install go
  ```  
  MacOS / Linux:  
  ```bash
  brew install go
  ```

- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)  
  Windows:  
  ```bash
  choco install protoc
  ```  
  MacOS / Linux:  
  ```bash
  brew install protobuf 
  ```

- Required Go modules and tools are defined in each module's **go.mod** file.

---

## Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/SylvanSol/Entain_Test.git
   cd Entain_Test
   ```

2. **Update Dependencies and Generate Protobuf Files:**

   Ensure you are using the updated Go toolchain (as specified in your go.mod) and that your modules use recent versions (for instance, gRPC-Go v1.71.1). Then regenerate your protobuf files as follows:

   **For the Racing Service:**

   - **Generating Protobuf Files (Racing):**
     ```powershell
     & 'C:\ProgramData\chocolatey\bin\protoc.exe' -I . --go_out=. --go_opt paths=source_relative --go-grpc_out=. --go-grpc_opt paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt paths=source_relative racing.proto
     ```

   **For the API Service:**

   - **Generating Protobuf Files (API):**
     ```powershell
     & 'C:\ProgramData\chocolatey\bin\protoc.exe' -I .. -I . --go_out=. --go_opt paths=source_relative --go-grpc_out=. --go-grpc_opt paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt paths=source_relative racing.proto
     ```

3. **Build the Services:**

   **For the Racing Service:**

   - Enable CGO and update your PATH for MinGW:
     ```powershell
     $env:CGO_ENABLED = "1"
     $env:PATH += ";C:\ProgramData\mingw64\mingw64\bin"
     go build
     ./racing
     ```

   **For the API Service:**
   ```bash
   go build
   ./api
   ```

---

## Recent Changes

### Addition of the "Visible Only" Filter

- **Proto Modification:**  
  In the `ListRacesRequestFilter` message (in `racing.proto`), a new optional boolean field has been added:
  ```proto
  optional bool visible_only = 2;
  ```
  This allows API consumers to request that only races with `visible = true` are returned.

- **Repository Update:**  
  The `applyFilter` function in **/racing/db/races.go** has been modified to append the clause `"visible = 1"` when `visible_only` is true.

- **Client Example:**
  ```bash
  curl -X POST http://localhost:8000/v1/list-races \
       -H 'Content-Type: application/json' \
       -d '{
         "filter": {
           "visibleOnly": true
         }
       }'
  ```

### gRPC Server Interface Issue and Resolution

While upgrading dependencies (with Go 1.24.1 and gRPC-Go v1.71.1), the following compiler error was encountered:
```
cannot use service.NewRacingService(racesRepo) (value of interface type service.Racing) as racing.RacingServer value in argument to racing.RegisterRacingServer: service.Racing does not implement racing.RacingServer (missing method mustEmbedUnimplementedRacingServer)
```
  
**Resolution Steps:**

1. **Remove Local Interface:**  
   Any custom local interface definitions (e.g. type Racing) were removed from the service package to rely solely on the generated `racing.RacingServer` interface.

2. **Embed Unimplemented Server:**  
   In **/racing/service/racing.go**, the service struct was updated to embed `racing.UnimplementedRacingServer` so that it automatically implements the required method:
   ```go
   type racingService struct {
       racing.UnimplementedRacingServer  // This embeds mustEmbedUnimplementedRacingServer.
       racesRepo db.RacesRepo
   }
   ```

3. **Constructor Return Type:**  
   The constructor was modified to return `racing.RacingServer`:
   ```go
   func NewRacingService(racesRepo db.RacesRepo) racing.RacingServer {
       return &racingService{racesRepo: racesRepo}
   }
   ```

4. **Re-generation and Clean:**  
   After these changes, the protobuf files were regenerated using the latest protoc plugins, and the build cache was cleaned using:
   ```bash
   go clean -cache -modcache
   ```

Despite these modifications, the error persisted until careful verification of import paths, regeneration of the proto files, and cleaning of cached modules resolved the conflict.

---
## Testing

A test file has been added in **/racing/db/queries_test.go** to verify the "visible only" filter and to serve as a basis for future testing:
  
- The test creates an in‑memory SQLite database, seeds it with test data (including one visible and one non‑visible race), and checks that filtering by visible races returns only races with visible = 1.

---
## Known Issues & Future Steps

- **Current Stuck Point:**  
  The primary issue (missing method `mustEmbedUnimplementedRacingServer`) has been addressed by ensuring your service implementation exactly matches the generated interface. If the error persists, verify that only one version of the generated package is used throughout the project.
  
- **Future Improvements:**  
  - Add additional sorting options for the `ListRaces` RPC.
  - Implement further endpoints, such as fetching a single race by ID and creating a separate Sports service.
  - Improve unit and integration tests.
  
---

## Contact

If you have any questions or require further clarification, please feel free to reach out.

---

## Bibliography

- [gRPC-Go Documentation – Official Site](https://pkg.go.dev/google.golang.org/grpc) citeturn7file4
- [Protocol Buffers Documentation](https://developers.google.com/protocol-buffers) citeturn7file4
- [gRPC-Gateway Documentation](https://grpc-ecosystem.github.io/grpc-gateway/) citeturn7file4

---