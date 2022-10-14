package main
// create by whatislovevladislav@gmail.com
import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func IntAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
func mod(x1, y1, m float64) float64 {
	return math.Mod(math.Pow(x1, y1), m)

}

func GCD2(a, b int) int {
	if b == 0 {
		return IntAbs(a)
	}
	return GCD2(b, a%b)
}

func main() {

	var x1 float64
	var y1 float64
	var m float64
	var b int
	var a int

	rand.Seed(time.Now().UnixNano())

	fmt.Print("Введіть модуль: ")
	fmt.Fscan(os.Stdin, &m)
	fmt.Println("розв’язувати рівняння виду a^b mod m = x")
	fmt.Print("Введіть число: ")
	fmt.Fscan(os.Stdin, &x1)
	fmt.Print("Введите степень: ")
	fmt.Fscan(os.Stdin, &y1)
	fmt.Println("Остаток :", mod(x1, y1, m))
	fmt.Println("розв’язувати рівняння виду  a*x ≡ b mod m")
	fmt.Print("Введіть перше число: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Print("Введіть друге число : ")
	fmt.Fscan(os.Stdin, &b)
	fmt.Println("Найбільший спільний дільник GCD:", GCD2(a, b))
	fmt.Println("просте число у діапазоні від A до B")
	fmt.Println("Випадкове число з діапазону:", a+rand.Intn(b-a+1))

}
