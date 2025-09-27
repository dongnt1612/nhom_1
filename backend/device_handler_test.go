// device_handler_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockDeviceRepo implements DeviceRepository
type MockDeviceRepo struct {
	devices []*Device
}

func (m *MockDeviceRepo) GetAll() []*Device {
	return m.devices
}

//	func (m *MockDeviceRepo) Add(d *Device) *Device {
//		d.ID = len(m.devices) + 1
//		m.devices = append(m.devices, d)
//		return d
//	}
func (m *MockDeviceRepo) Update(id int, d *Device) (*Device, bool) {
	for _, dev := range m.devices {
		if dev.ID == id {
			dev.Name = d.Name
			dev.Type = d.Type
			dev.Status = d.Status
			return dev, true
		}
	}
	return nil, false
}

//func (m *MockDeviceRepo) Delete(id int) bool {
//	for i, dev := range m.devices {
//		if dev.ID == id {
//			m.devices = append(m.devices[:i], m.devices[i+1:]...)
//			return true
//		}
//	}
//	return false
//}

func TestUpdateDevice(t *testing.T) {
	mockRepo := &MockDeviceRepo{
		devices: []*Device{
			{ID: 1, Name: "Laptop", Type: "Computer", Status: "Active"},
		},
	}
	store = mockRepo // override store để test

	// Payload update
	body := []byte(`{"name":"Laptop Pro","type":"Computer","status":"Inactive"}`)
	req, _ := http.NewRequest("PUT", "/api/devices/1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(devicesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Expected status 200, got %v", status)
	}

	var updated Device
	if err := json.NewDecoder(rr.Body).Decode(&updated); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check field values
	if updated.Name != "Laptop Pro" {
		t.Errorf("Expected name Laptop Pro, got %s", updated.Name)
	}
	if updated.Status != "Inactive" {
		t.Errorf("Expected status Inactive, got %s", updated.Status)
	}
}

func TestUpdateDevice_InvalidJSON(t *testing.T) {
	mockRepo := &MockDeviceRepo{
		devices: []*Device{
			{ID: 1, Name: "Laptop", Type: "Computer", Status: "Active"},
		},
	}
	store = mockRepo

	// Body không phải JSON hợp lệ
	body := []byte(`{invalid-json}`)
	req, _ := http.NewRequest("PUT", "/api/devices/1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(devicesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fatalf("Expected status 400, got %v", status)
	}
}

func TestUpdateDevice_InvalidID(t *testing.T) {
	mockRepo := &MockDeviceRepo{
		devices: []*Device{
			{ID: 1, Name: "Laptop", Type: "Computer", Status: "Active"},
		},
	}
	store = mockRepo

	body := []byte(`{"name":"Laptop Pro","type":"Computer","status":"Active"}`)
	req, _ := http.NewRequest("PUT", "/api/devices/abc", bytes.NewBuffer(body)) // ID không hợp lệ
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(devicesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fatalf("Expected status 400, got %v", status)
	}
}

func TestAddDevice_MissingField(t *testing.T) {
	mockRepo := &MockDeviceRepo{}
	store = mockRepo

	// JSON thiếu field "type"
	body := []byte(`{"name":"Tablet","status":"Active"}`)
	req, _ := http.NewRequest("POST", "/api/devices", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(devicesHandler)
	handler.ServeHTTP(rr, req)

	// server hiện tại không validate field bắt buộc → vẫn trả 200
	// nhưng test này để check thực tế hành vi
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Expected status 200 (no validation yet), got %v", status)
	}

	var d Device
	json.NewDecoder(rr.Body).Decode(&d)
	if d.Type == "" {
		t.Logf("Warning: Device.Type is empty, consider adding validation")
	}
}
