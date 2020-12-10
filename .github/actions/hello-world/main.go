package main
import (
   "fmt"
   "os"
)

func main() {
   input1 := os.Getenv("INPUT_GREETING1")
   input2 := os.Getenv("INPUT_GREETING2")
   input3 := os.Getenv("INPUT_GREETING3")
   
   fmt.Println(input1)
   fmt.Println(intpu2)
   
   if input3 != "" {
      fmt.Println(input3)
   }
}
