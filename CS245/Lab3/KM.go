package main
import (
	"fmt"
	"os"
)

func main() {
	var number float32 = os.Args[0];
        var unit string = os.Args[1];
        var kilometers float32 = number * 1.609;
        var miles float32 = number / 1.609;
        if (unit == "m") {
            fmt.Println(kilometers + " km"); //Don't know how to print two different types
        } else if (unit == "k") {
            fmt.Println(miles + " miles");// Don't know how to print two different types
        }
}