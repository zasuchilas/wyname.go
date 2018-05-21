package main

// Sector регистрирует лайферов
type Sector struct {
	members     map[*Lifer]bool // члены сектора
	subscribers map[*Lifer]bool // подписчики сектора

}
