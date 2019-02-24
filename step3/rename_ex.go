package main

import (
    "os"
)

func main() {
    os.Rename("./" + "a", "./" + "b")
}