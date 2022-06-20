package main

import (
	"fmt"
)

//Для того чтобы поменять i-ый бит на 1, создадим маску c единственной единицей в позиции i, отсчитывая справа.
//Пример:
//   число  11 // 1011
//   поменяем в нем  2-ой бит ( 0 ) на 1
//   создаем маску 0100
//   создать такую маску можем сдвинув бит на 2 позицию (mask = int64(1) << 2)
//   выполнив побитовое или, биты на соотвествующих позициях поменяются на 0 если a[i] == 0 и b[i] == 0,
//   в остальных случаях будут раны 1
//
//   исходное число : 1011
//   маска :          0100
//   битвое или |
//   результат        1111
//
//Для того, чтобы поменять i-ый бит на 0 воспользуемся оператором очистки &^ (еще говорят, опустим i-ый бит)
//то есть нам надо создать аналогичную маску и использовать этот оператор. Здесь уже важно, что
//маска будет стоять справа от оператора.
//Если i-ый бит числа равен 1, то он будет заменен на 0, а если он был равен 0, то останется нулем.
//Пример:
//   число 11 // 1011
//   поменяем в нем 1-ый бит на 0
//   создадим маску 0010 // это число 2
//   11 &^ 2 == 9 // 1001
//
//   если применить подобну операцию для чисел 1001 и 0010 то получим все равно 1001, так как 1-ый бит в
//   исходном числе и так равен 0

func ChangeByte(value int64, i int, bitNumber int) int64 {
	if i == 1 {
		// сдвиг влево на bitNumber позиций
		// было 10, сдвинули на 3
		// стало 10000
		// |= - логическое или, что означает
		// будет ноль, когда оба ноль, но мы сдвигаем едицину, тем самым обеспечивая, что в данной позиции меняется бит
		// в других позициях ничего не изменится из-за одной единициы и остальных нулей
		value |= int64(1) << bitNumber
	}
	if i == 0 {
		// сдвиг влево на bitNumber позиций
		// было 10, сдвинули на 3
		// стало 10000
		// z = x &^ y - логическое И НЕ. в z - 0, когда  y - 1. если бит в y равен 0, то берется бит из x
		// z = 0
		// сдвигаем
		value &^= int64(1) << bitNumber
	}
	return value
}

func main() {
	value := 156
	i := 1
	bitNumber := 2
	fmt.Println(ChangeByte(int64(value), i, bitNumber))
}