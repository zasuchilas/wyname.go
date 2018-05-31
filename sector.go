package main

// Sector регистрирует лайферов
type Sector struct {
	members     map[*Lifer]bool // члены сектора
	subscribers map[*Lifer]bool // подписчики сектора

}

func newsector() *Sector {
	return &Sector{
		members:     make(map[*Lifer]bool, 200),
		subscribers: make(map[*Lifer]bool, 300),
	}
}
