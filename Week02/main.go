package main

import (
	"fmt"
	"train/Week02/Service"
)

func main()  {
	u, err := Service.UserInfo(100)
	if err == nil {
		fmt.Println(u)
	} else {
		fmt.Printf("%+v\n", err)
	}
}
