// repository.go
package main

type DeviceRepository interface {
	GetAll() []*Device
	//Add(d *Device) *Device
	Update(id int, d *Device) (*Device, bool)
	//Delete(id int) bool
}
