package main

import (
	"fmt"
	"os"
)

const URL = "https://dev.pamongdesa.id/catalog"

func main() {
	cookie := os.Getenv("LARATI")
	fmt.Println(cookie)
	RunGET(URL, cookie)
}
