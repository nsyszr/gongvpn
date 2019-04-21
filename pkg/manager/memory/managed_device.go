package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nsyszr/ngvpn/pkg/model"
)

type managedDeviceMemoryManager struct {
	store map[string]model.ManagedDevice
	sync.RWMutex
}

func newManagedDeviceMemoryManager() *managedDeviceMemoryManager {
	return &managedDeviceMemoryManager{
		store: make(map[string]model.ManagedDevice),
	}
}

func (m *managedDeviceMemoryManager) FetchAll() (devices map[string]model.ManagedDevice, err error) {
	m.RLock()
	defer m.RUnlock()
	devices = make(map[string]model.ManagedDevice, len(m.store))

	for id, device := range m.store {
		devices[id] = device
	}

	return devices, nil
}

func (m *managedDeviceMemoryManager) FindByID(id string) (*model.ManagedDevice, error) {
	m.RLock()
	defer m.RUnlock()
	if device, ok := m.store[id]; ok {
		return &device, nil
	}

	return nil, fmt.Errorf("not found")
}

func (m *managedDeviceMemoryManager) Create(device *model.ManagedDevice) error {
	id := uuid.New().String()

	if _, err := m.FindByID(id); err != nil {
		return fmt.Errorf("unique violation")
	}

	m.Lock()
	defer m.Unlock()

	device.ID = id
	device.CreatedAt = time.Now().Round(time.Second)
	device.UpdatedAt = time.Now().Round(time.Second)

	m.store[device.ID] = *device

	return nil
}
