package calc_test

import (
	"fmt"
	"testing"

	"github.com/ssoyyoung.p/Go-Example/goTest/calc"
)

func TestSum(t *testing.T) {
	res := calc.Sum(1, 2, 3)

	fmt.Println(res)

	if res != 6 {
		t.Error("Wrong Result")
	}
}

// test를 위한 명령어는 go test이다.
