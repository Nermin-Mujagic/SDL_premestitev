# SDL_premestitev

A Go-based student dormitory transfer request management system for Ljubljana's student housing network.

## Overview

SDL_premestitev is a domain-driven application that manages student housing transfer requests for Ljubljana's extensive dormitory system. The system handles the complexity of student preferences, room types, partner matching, and dormitory availability across multiple housing locations.

## Features

- **Comprehensive Validation System**: Input validation for student IDs, dormitory preferences, room types, and partner requirements
- **Flexible Room Configuration**: Support for traditional dormitory rooms and apartment-style housing with various occupancy options
- **Partner Matching**: System for students requesting shared accommodations with specific partners
- **Priority Queue Management**: Request processing with priority-based ordering (in development)
- **Extensive Dormitory Database**: Complete mapping of Ljubljana's student housing network including Rožna Dolina, Bežigrad, Center, and other locations

## Architecture

The project follows clean architecture principles with clear separation of concerns:

```
internal/
├── requests/           # Core domain logic and validation
├── priority_queue/     # Request processing and ordering
└── helpers/           # Testing utilities and shared functions
```

### Key Design Decisions

- **Validation at Creation**: Fail-fast approach with comprehensive error handling
- **Pointer-based Optional Fields**: Idiomatic Go handling of nullable data
- **Constants-based Configuration**: Maintainable dormitory and room type definitions
- **Comprehensive Test Coverage**: Unit tests covering happy paths and edge cases

## Domain Model

### Transfer Request
```go
type TransferRequest struct {
    RequestID      string
    StudentID      string
    PreferredDorms []string
    Apartment      bool           // Traditional vs apartment-style housing
    RoomType       *string        // "singleBed", "doubleBed", "coupleApartment"
    WithPartner    bool
    PartnerID      *string
    DateSubmitted  time.Time
    Status         RequestStatus  // "active", "inactive"
}
```

### Supported Dormitories
The system supports all major Ljubljana student housing locations:
- **Rožna Dolina**: Dom I through XIV
- **Bežigrad**: Dom A-D, FDV, AK, Topniška
- **Center**: Poljanska, VŠZ, Ilirska
- **Litostroj**: Dom Litostroj
- **Vič**: ŠD3, ŠD4, G59

## Getting Started

### Prerequisites
- Go 1.23.5 or later

### Installation
```bash
git clone https://github.com/yourusername/SDL_premestitev.git
cd SDL_premestitev
go mod tidy
```

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test -v ./internal/requests
go test -v ./internal/priority_queue
```

### Example Usage
```go
// Create a simple transfer request
request, err := CreateTransferRequest(
    "student123",           // Student ID
    []string{"dom_fdv"},    // Preferred dormitories
    true,                   // Apartment-style housing
    &roomType,             // Room type preference
    false,                 // No partner
    nil,                   // Partner ID
)

if err != nil {
    log.Fatal(err)
}

// Add to priority queue
queue := NewPriorityList()
queue.AddRequest(*request)
```

## Current Status

**🚧 Active Development**

This project is currently in active development with the following completed:
- ✅ Core domain model and validation
- ✅ Comprehensive test suite
- ✅ Request creation and basic validation
- ✅ Priority queue foundation

**Planned Features:**
- 🔄 Enhanced dormitory configuration modeling
- 🔄 Business logic validation layer
- 📋 CLI interface for request management
- 📋 Persistence layer (file-based or database)
- 📋 Request matching and assignment algorithms
- 📋 HTTP API endpoints

## Development Philosophy

This project emphasizes:
- **Clean, testable code** over premature optimization
- **Domain-driven design** reflecting real-world housing complexity
- **Incremental development** with working software at each stage
- **Comprehensive testing** to ensure reliability

## Contributing

This is a personal learning project, but feedback and suggestions are welcome! Please feel free to:
- Report issues or bugs
- Suggest improvements to the architecture
- Share insights about Go best practices

## Technical Details

- **Language**: Go 1.23.5
- **Architecture**: Clean Architecture with domain-driven design
- **Testing**: Standard Go testing with custom assertion helpers
- **Dependencies**: Standard library only (no external dependencies)

## License

This project is available under the MIT License. See LICENSE file for details.

---

*Built as a portfolio project demonstrating Go programming skills, software architecture, and domain modeling.*
