package manager

import "github.com/nsyszr/ngvpn/pkg/model"

type Manager interface {
	ManagedDevices() ManagedDeviceManager
	Groups() GroupManager
}

// ManagedDeviceManager is responsible for managing ManagedDevice model
type ManagedDeviceManager interface {
	FetchAll() (map[string]model.ManagedDevice, error)
	FindByID(id string) (*model.ManagedDevice, error)
	Create(m *model.ManagedDevice) error
}

// GroupManager is responsible for managing Group model
type GroupManager interface {
	FetchAll() (map[string]model.Group, error)
	FindByID(id string) (*model.Group, error)
	Create(m *model.Group) error
}
