package dynamic

import (
	"fmt"
	"testing"
)

func TestMaxSubSeq(t *testing.T) {
	fmt.Println(MaxSubSeq("1A2C3D4B56", "B1D23CA45B6A"))
	fmt.Println(MaxSubSeq2("1A2C3D4B56", "B1D23CA45B6A"))
}

func TestMaxSubStr(t *testing.T) {
	fmt.Println(MaxSubStr("1AB2345CD", "12345EF"))
}

func TestOperateStrCost(t *testing.T) {
	fmt.Println(OperateStrCost("abc", "adc", 5, 3, 2))
}

func TestIsMix(t *testing.T) {
	fmt.Println(IsMix("AB", "12", "A1B2"))
	fmt.Println(IsMix("AB", "12", "AB12"))
	fmt.Println(IsMix("AB", "12", "AB13"))
}

func TestConvertLetter(t *testing.T) {
	fmt.Println(ConvertLetter("01"))
}
