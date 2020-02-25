package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now()
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format(time.Kitchen))
	fmt.Println(t.Second())
}