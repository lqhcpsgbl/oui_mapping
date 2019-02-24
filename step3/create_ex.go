package main

import (
    "os"
)

func main() {
    filename := "Wed, 16 Jan 2019 11:04:36 GMT.data"
    f, _ := os.Create("./" + filename)
    f.Close()
}