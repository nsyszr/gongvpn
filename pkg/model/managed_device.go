package model

import "time"

/*
{
	"id": 1,
	"deviceType": "M3",
	"name": "mrx-1",
	"location": "Regensburg",
	"assetTag": "100200",
	"group": "Standard",
	"hardwareModel": "M3CPU",
	"hardwareRevision": "3",
	"hardwareSerialNumber": "21_of_33",
	"firmwareVersion": "3.4",
	"networkHostname": "icom",
	"networkDomainname": "local",
	"networkPrimaryIPv4Address": "192.168.10.1",
	"availabilitySessionTimeout": 120,
	"availabilityPingInterval": 104,
	"availabilityPongResponseInterval": 16,
	"availabilityLastMessageAt": "2019-04-13T10:00:00Z",
	"availabilityStatus": "CONNECTED",
	"connectionStatus": "CONNECTED"
}
*/

// ManagedDevice contains all properties of a managed device
type ManagedDevice struct {
	ID                               string    `json:"id"`
	DeviceType                       string    `json:"deviceType"`
	Name                             string    `json:"name"`
	Location                         string    `json:"location"`
	AssetTag                         string    `json:"assetTag"`
	Group                            Group     `json:"group"`
	HardwareModel                    string    `json:"hardwareModel"`
	HardwareRevision                 string    `json:"hardwareRevision"`
	HardwareSerialNumber             string    `json:"hardwareSerialNumber"`
	FirmwareVersion                  string    `json:"firmwareVersion"`
	NetworkHostname                  string    `json:"networkHostname"`
	NetworkDomainname                string    `json:"networkDomainname"`
	NetworkPrimaryIPv4Address        string    `json:"networkPrimaryIPv4Address"`
	AvailabilitySessionTimeout       int       `json:"availabilitySessionTimeout"`
	AvailabilityPingInterval         int       `json:"availabilityPingInterval"`
	AvailabilityPongResponseInterval int       `json:"availabilityPongResponseInterval"`
	AvailabilityLastMessageAt        time.Time `json:"availabilityLastMessageAt"`
	AvailabilityStatus               string    `json:"availabilityStatus"`
	ConnectionStatus                 string    `json:"connectionStatus"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
