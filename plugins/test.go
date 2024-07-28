package plugins

import "fmt"

// 如果不想让业务与系统服务的代码耦合在一块可以把业务写在此模块下
func init() {
	fmt.Println("########## plugins init #####################")
}
