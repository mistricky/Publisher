package utils

import (
	"fmt"
	"time"
)

func Log(msg string){
	fmt.Printf("[%s]: %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}