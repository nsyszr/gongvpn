package memory

import "github.com/nsyszr/ngvpn/pkg/manager"

type MemoryManager struct {
	managedDevices *managedDeviceMemoryManager
	groups         *groupMemoryManager
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		managedDevices: newManagedDeviceMemoryManager(),
		groups:         newGroupMemoryManager(),
	}
}

func (m *MemoryManager) ManagedDevices() manager.ManagedDeviceManager {
	return m.managedDevices
}

func (m *MemoryManager) Groups() manager.GroupManager {
	return m.groups
}
