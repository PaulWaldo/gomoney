package db

import (
	"fmt"
	"sync"

	"github.com/PaulWaldo/gomoney/pkg/domain"
)

// memoryStore implements domain.AccountDB with an in-memory storage
type memoryStore struct {
	accounts map[domain.AccountIDType]*domain.Account
	sync.Mutex
	nextId domain.AccountIDType
}

func (ms *memoryStore) String() string {
	return fmt.Sprintf("Memory Store: num accounts=%d, next ID=%d", len(ms.accounts), ms.nextId)
}

func NewMemoryStore() domain.AccountDB {
	ms := &memoryStore{
		accounts: make(map[domain.AccountIDType]*domain.Account),
	}
	ms.nextId = 0
	return ms
}

func (ms *memoryStore) Create(name string, accountType domain.AccountType) (domain.AccountIDType, error) {
	ms.Lock()
	defer ms.Unlock()
	acct := &domain.Account{
		ID:          ms.nextId,
		Name:        name,
		AccountType: accountType,
	}
	ms.accounts[ms.nextId] = acct
	ms.nextId++
	return acct.ID, nil
}

func (ms *memoryStore) Get(id domain.AccountIDType) (*domain.Account, error) {
	acct, exists := ms.accounts[id]
	if !exists {
		return nil, fmt.Errorf("no pet found with id: %d", id)
	}
	return acct, nil
}

func (ms *memoryStore) List() ([]*domain.Account, error) {
	accounts := []*domain.Account{}
	for _, p := range ms.accounts {
		accounts = append(accounts, p)
	}
	return accounts, nil
}

func (ms *memoryStore) Delete(id domain.AccountIDType) error {
	delete(ms.accounts, id)
	return nil
}
