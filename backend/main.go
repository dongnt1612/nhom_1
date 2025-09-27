package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Device represents a device in our system
type Device struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

// DeviceStore manages devices in memory
type DeviceStore struct {
	mu      sync.RWMutex
	devices []*Device
}

// NewDeviceStore creates a new device store with sample data
func NewDeviceStore() *DeviceStore {
	return &DeviceStore{
		devices: []*Device{
			{ID: 1, Name: "Laptop", Type: "Computer", Status: "Active"},
			{ID: 2, Name: "Printer", Type: "Peripheral", Status: "Inactive"},
			{ID: 3, Name: "Router", Type: "Network", Status: "Active"},
		},
	}
}

// GetAll returns all devices
func (ds *DeviceStore) GetAll() []*Device {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.devices
}

// Update modifies an existing device by ID
func (ds *DeviceStore) Update(id int, d *Device) (*Device, bool) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for _, dev := range ds.devices {
		if dev.ID == id {
			dev.Name = d.Name
			dev.Type = d.Type
			dev.Status = d.Status
			return dev, true
		}
	}
	return nil, false
}

var store = NewDeviceStore()

// CORS middleware
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Helper: extract ID từ URL
func getIDFromPath(path string) (int, bool) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 {
		return 0, false
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, false
	}
	return id, true
}

// HTTP Handlers
func devicesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case "GET":
		devices := store.GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(devices)

	case "PUT":
		id, ok := getIDFromPath(r.URL.Path)
		if !ok {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		var d Device
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		err := d.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated, found := store.Update(id, &d)
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	//if r.Method != "GET" {
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//	return
	//}
	//
	//devices := store.GetAll()
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(devices)
}

func main() {
	// Serve static files (our HTML frontend)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "frontend/index.html")
		} else {
			http.NotFound(w, r)
		}
	})

	// API routes
	http.HandleFunc("/api/devices", devicesHandler)
	http.HandleFunc("/api/devices/", devicesHandler)

	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("Frontend available at: http://localhost:8080")
	fmt.Println("API endpoints:")
	fmt.Println("  GET    /api/devices")

	log.Fatal(http.ListenAndServe(port, nil))
}

func (d *Device) Validate() error {
    if d.ID <= 0 {
        return fmt.Errorf("id must be greater than 0")
    }
    if strings.TrimSpace(d.Name) == "" {
        return fmt.Errorf("name cannot be empty")
    }

    return nil
}