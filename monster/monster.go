package monster

type Monster interface {
	GetHealth() int
	SetHealth(health int)
	SetName(name string)
	GetName() string
}

type monster struct {
	health int
	name   string
}

func (m *monster) GetHealth() int {
	return m.health
}

func (m *monster) SetHealth(health int) {
	m.health = health
}

func (m *monster) SetName(name string) {
	m.name = name
}

func (m *monster) GetName() string {
	return m.name
}

func New() Monster {
	return &monster{}
}
