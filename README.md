# Entain Technical Test

## Overview

This repository is my solution for the Entain Technical Test. The project is structured as a multi-service Go application that demonstrates the following:

---

## Project Structure

- `api`: A basic REST gateway, forwarding requests onto service(s).
- `racing`: A very bare-bones racing service.
```
entain/
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
│  |  ├─ races.go
│  |  ├─ racing.go                
│  ├─ proto/
│  |  ├─ racing/
│  |  |  ├─ racing.pd.go
│  |  |  ├─ racing.proto
│  |  |  ├─ racing_grpc.pd.go
│  |  ├─ racing.go           
│  ├─ service/
│  |  ├─ racing.go
│  ├─ go.mod
│  ├─ go.sum      
│  ├─ main.go
│  ├─ tools.go            
├─ example.png                                      # The Example PNG for the test
├─ givenREADME.md                                   # The Example README as instrucitons for the test      
├─ README.md                                        # This file
```

---

## Getting Started

### Prerequisites

- [Go (latest version)](https://golang.org/doc/install)
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
- Required Go modules and tools:
  ```bash
  nothing here yet
  ```

### Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/yourusername/entain-technical-test.git
   cd entain-technical-test
   ```

## Testing



---

## Additional Information

---

## Future Improvements

---

## Contact

If you have any questions or require further clarification, please feel free to reach out.

---
