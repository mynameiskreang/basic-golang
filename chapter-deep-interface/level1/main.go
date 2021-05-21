package main

import "fmt"

type Weapon interface {
	GetName() string
}

type Sword struct {
	Name string
}

type Bow struct {
	Name string
}

// Step 2:
// ดังนั้น Sword, Bow Struct ต้อง implement Weapon จะได้
func (s Sword) GetName() string {
	return s.Name
}
func (b Bow) GetName() string {
	return b.Name
}

// End Step 2

func main() {
	sword := Sword{"normal-sword"}
	bow := Bow{"magic-bow"}

	attack(sword)
	attack(bow)
}

// Step 1:
// func attack ต้องรับได้ struct ใดๆได้
// เพราะถ้าทำแบบ attack(sword Sword) หรือ attack(bow Bow) จะได้แค่อย่างเดียว
// ดังนั้นต้องทำ interface กลาง เลยได้เป็น attack(w Weapon)
func attack(w Weapon) {
	fmt.Println("Attack by:", w.GetName())
}
