package main // <-- comm: нужно указать главную точку входа

//var justString string <-- comm: объявление глобальной переменной не является хорошим тоном, если
// это не сбор ошибок, ключей и так далее

// comm: есть важный и опасный нюанс:
// 1) Если, к примеру, v имела бы 100000000 символов, а нам нужно было бы только 100, то тогда бы
// у нас была утечка памяти, так как сборщик мусора не сможет очистить память от v.

func createHugeString(len int) string {

	// string -> это массив рун
	return string(make([]byte, len))
}

func someFunc() string {
	v := createHugeString(1 << 10)
	// Мы копируем только нужный кусок, давай GC очистить память по v

	//return strings.Clone(v[:100])
	//или
	return string([]byte(v[:100]))

}

func main() {
	someFunc()
}
