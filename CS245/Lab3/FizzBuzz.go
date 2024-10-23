package main
import "fmt"

func main() {
	count := 0
        for i := 0; i < 100; i++ {
            count++;
            if ((count%3) == 0) && ((count%5) == 0){
                fmt.Println("FizzBuzz");
            } else if (count%5) == 0 {
                fmt.Println("Buzz");
            } else if (count%3) == 0 {
                fmt.Println("Fizz");
            } else {
                fmt.Println(count);
            }
        }
}