package main

import "fmt"

type Weapon interface {
	GetName() string
}

// Step 2:
// MagicWeapon ก็ต้องมีชื่อเหมือนกัน เลย implement Weapon
type MagicWeapon interface {
	// ตรงนี้เรียกว่า embed
	Weapon
	GetMagic() string
}

type Sword struct {
	Name string
}

type Bow struct {
	Name string
}

func (s Sword) GetName() string {
	return s.Name
}
func (b Bow) GetName() string {
	return b.Name
}

// Step 3: bow implement GetMagic()
func (b Bow) GetMagic() string {
	return "Magic Attack!!!"
}

func main() {
	sword := Sword{"normal-sword"}
	bow := Bow{"magic-bow"}

	attack(sword)
	attack(bow)

	// Step 1:
	// จาก level1 ที่ bow attack ธรรมดา
	// มา level2 จะใช้ magic
	magic(bow)
	// alternatively สามารถสร้าง interface แล้ว pass bow ได้
	// แต่จะ MagicWeapon(sword) ไม่ได้
	mw := MagicWeapon(bow)
	magic(mw)
}

func magic(mw MagicWeapon) {
	fmt.Println(mw.GetName(), ":", mw.GetMagic())
}

func attack(w Weapon) {
	fmt.Println("Attack by:", w.GetName())
}
