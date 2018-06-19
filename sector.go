package main

// Sector регистрирует лайферов
type Sector struct {
	members map[*Lifer]bool // члены сектора
	subscrs map[*Lifer]bool // подписчики сектора

}

func newsector() *Sector {
	return &Sector{
		members: make(map[*Lifer]bool, 200),
		subscrs: make(map[*Lifer]bool, 300),
	}
}
