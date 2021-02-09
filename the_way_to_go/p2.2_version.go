//在命令行输入go version打印go版本信息

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s", runtime.Version())
}
