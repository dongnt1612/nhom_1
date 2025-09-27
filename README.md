# Device Management System

A simple CRUD application for managing devices with a Go backend and vanilla HTML/CSS/JS frontend.

## Features

- **Backend**: Go server with REST API (no external dependencies)
- **Frontend**: Single HTML file with inline CSS and JavaScript
- **Storage**: In-memory storage (resets on server restart)
- **CRUD Operations**: Create, Read, Update, Delete devices
- **Device Properties**: ID, Name, Type, Status

## Project Structure

```
/
├── backend/
│   ├── main.go       # Go server with all logic
│   ├── main_test.go  # Comprehensive tests
│   └── go.mod        # Go module file
├── frontend/
│   └── index.html    # Single HTML file with UI
└── package.json      # NPM scripts for convenience
```

## Prerequisites

- Go 1.21 or later
- A web browser

## Running the Application

### Option 1: Using npm scripts (if you have Node.js)

```bash
# Run the server
npm run dev

# Run tests
npm run test:backend

# Build binary
npm run build:backend
```

### Option 2: Using Go directly

```bash
# Run the server
go run backend/main.go

# Run tests
go test -v ./backend

# Build binary
go build -o device-management ./backend
```

## Usage

1. Start the server:
   ```bash
   npm run dev
   # or
   go run backend/main.go
   ```

2. Open your browser and go to: `http://localhost:8080`

3. Use the web interface to:
   - View all devices in the table
   - Add new devices using the form
   - Edit existing devices by clicking "Edit"
   - Delete devices by clicking "Delete"

## API Endpoints

The backend provides a REST API:

- `GET /api/devices` - Get all devices
- `POST /api/devices` - Create a new device
- `PUT /api/devices/{id}` - Update a device
- `DELETE /api/devices/{id}` - Delete a device

### Example API Usage

```bash
# Get all devices
curl http://localhost:8080/api/devices

# Create a device
curl -X POST http://localhost:8080/api/devices \
  -H "Content-Type: application/json" \
  -d '{"name":"My Laptop","type":"laptop","status":"active"}'

# Update a device
curl -X PUT http://localhost:8080/api/devices/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Laptop","type":"laptop","status":"inactive"}'

# Delete a device
curl -X DELETE http://localhost:8080/api/devices/1
```

## Sample Data

The application starts with 3 sample devices:
1. MacBook Pro (laptop, active)
2. iPhone 15 (phone, active)
3. iPad Air (tablet, inactive)

## Device Types

- laptop
- phone
- tablet
- desktop
- server
- other

## Device Statuses

- active
- inactive
- maintenance

## Testing

The backend includes comprehensive tests covering:
- Device store operations (CRUD)
- HTTP handlers
- Error cases
- Edge cases

Run tests with:
```bash
npm run test:backend
# or
go test -v ./backend
```

## Architecture

### Backend (Go)

- **Single file**: All logic in `main.go` for simplicity
- **In-memory storage**: Thread-safe map with mutex
- **Standard library only**: No external dependencies
- **CORS enabled**: For frontend integration
- **RESTful API**: Standard HTTP methods and status codes

### Frontend (HTML/CSS/JS)

- **Single file**: Everything in `index.html`
- **Vanilla JavaScript**: No frameworks or build tools
- **Responsive design**: Works on desktop and mobile
- **Real-time updates**: UI updates immediately after operations
- **Error handling**: User-friendly error messages
