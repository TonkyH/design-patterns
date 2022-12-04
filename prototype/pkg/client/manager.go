package client

import "prototype/pkg/prototype"

type Manager struct {
	Showcase map[string]prototype.Product
}

func NewManager() *Manager {
	return &Manager{Showcase: map[string]prototype.Product{}}
}

func (m *Manager) Register(name string, prototype prototype.Product) {
	m.Showcase[name] = prototype
}

func (m *Manager) Create(prototypeName string) prototype.Product {
	p := m.Showcase[prototypeName]
	return p.CreateCopy()
}
