package hello

import (
    "fmt"
    "time"
)

var Routes = map[string]interface{}{
    "/today": today,
    "/(.*)": hello,
}

func hello(val string) string { 
    return fmt.Sprintf("hello %s", val) 
}

func today() string {
    return fmt.Sprintf("The time is currently %s", time.LocalTime().Asctime())
}

