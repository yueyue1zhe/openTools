package commutil

import (
	"fmt"
	"testing"
	"time"
)

func TestSecondToDay(t *testing.T) {
	fmt.Println(SecondToDay(1))
	fmt.Println(SecondToDay(60 * 60 * 12))
	fmt.Println(SecondToDay(60 * 60 * 24))
	fmt.Println(SecondToDay(60 * 60 * 25))
	fmt.Println(SecondToDay(60 * 60 * 20))
}

func TestLastSecond(t *testing.T) {
	end := time.Now().AddDate(0, -1, 0)
	fmt.Println(SecondToDay(int(end.Unix() - time.Now().Unix())))
}

func TestAddQ(t *testing.T) {
	val := time.Now().AddDate(0, 3, 0).Add(time.Second*(60*60*12)).Unix() - time.Now().Unix()
	fmt.Println(SecondToDay(int(val)))
}
