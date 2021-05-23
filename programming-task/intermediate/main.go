package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {

}

// https://programming.in.th/task/rev2_problem.php?pid=1106
// k := 2
// n := []int{1, 2, 3, 4, 5}
// output => 3
func P1106(n []int, k int) int {
	if len(n) == 1 {
		return n[0]
	}
	n = append(n[k-1+1:], n[:k-1]...)
	P1106(n, k)
	return n[0]
}

//https://programming.in.th/task/rev2_problem.php?pid=1031
// โจทย์แม่งเชี้ยมาก งงชิบหาย
// สรุปคือ ถ้าพลัง >= a แล้ว เก็บ b ด้านขวาไว้
// จากนั้น b ไหนมากสุด ก็คือชั้นที่ไปได้
func P1031() {
	var k, n, m int
	fmt.Scanf("%d %d %d", &k, &n, &m)
	var floor = 1
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(a, b)
		if k >= a {
			floor = b
		}
	}
	fmt.Println(floor)
}

//https://programming.in.th/task/rev2_problem.php?pid=1000
func P1000() {
	//โซ่คำ คือลำดับของคำที่มีจำนวนอักขระเท่ากันและแต่ละคำที่มีลำดับติดกันจะต้องมีตำแหน่งที่มีตัวอักขระต่างกันไม่เกินสองตำแหน่ง
	//ตัวอย่างของโซ่คำที่ต่อเนื่องได้แก่ HEAD HEAP LEAP TEAR REAR และ EGG EAG GAE GAP TAP TIN
	//ตัวอย่างของโซ่คำที่ขาดได้แก่ LEAP TEAR REAR BAER BAET BEEP ซึ่งจะขาดที่ คำว่า BAER

	//input := []string{"LEAP", "TEAR", "REAR", "BAER", "BAET", "BEEP"}
	input := []string{"HEAD", "HEAP", "LEAP", "TEAR", "REAR", "BAER", "BAET", "BEEP", "JEEP", "JOIP", "JEIP", "AEIO"}
	for i := 0; i < len(input)-1; i++ {
		err := matching(input[i], input[i+1])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func matching(input1, input2 string) error {
	if len(input1) != len(input2) {
		return errors.New("len(input1) != len(input2)")
	}
	flagFail := 0
	for i, _ := range input1 {
		if input1[i] != input2[i] {
			flagFail++
		}
	}
	if flagFail > 2 {
		return errors.New(input2)
	}
	return nil
}
