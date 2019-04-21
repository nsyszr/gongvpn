package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nsyszr/ngvpn/pkg/model"
)

type groupMemoryManager struct {
	store map[string]model.Group
	sync.RWMutex
}

func newGroupMemoryManager() *groupMemoryManager {
	return &groupMemoryManager{
		store: make(map[string]model.Group),
	}
}

func (m *groupMemoryManager) FetchAll() (groups map[string]model.Group, err error) {
	m.RLock()
	defer m.RUnlock()
	groups = make(map[string]model.Group, len(m.store))

	for id, group := range m.store {
		groups[id] = group
	}

	return groups, nil
}

func (m *groupMemoryManager) FindByID(id string) (*model.Group, error) {
	m.RLock()
	defer m.RUnlock()
	if group, ok := m.store[id]; ok {
		return &group, nil
	}

	return nil, fmt.Errorf("not found")
}

func (m *groupMemoryManager) Create(group *model.Group) error {
	id := uuid.New().String()

	if _, err := m.FindByID(id); err == nil {
		return fmt.Errorf("unique violation")
	}

	m.Lock()
	defer m.Unlock()

	group.ID = id
	group.CreatedAt = time.Now().Round(time.Second)
	group.UpdatedAt = time.Now().Round(time.Second)

	m.store[group.ID] = *group

	return nil
}
