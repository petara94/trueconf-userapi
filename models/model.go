package models

import (
	"sync"
	"time"
)

type Model struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ModelList map[int]Model

type ModelStore struct {
	sync.RWMutex
}
