package api

import (
	"fmt"
	_ "plugins/api"
)

func init() {
	fmt.Println("############## plugins init ########################")
}
