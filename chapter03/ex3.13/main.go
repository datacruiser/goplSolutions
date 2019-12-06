package main

import "fmt"

const (
	KB = 1024
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println(KB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(EB)
}
