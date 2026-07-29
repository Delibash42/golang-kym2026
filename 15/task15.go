package main

// в памяти хранится вся строка а не только ее срез в v и хранится по сути 924 байта без дела

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
}
func main() {
	someFunc()
}

// так мы помещаем данные а не срез из v
