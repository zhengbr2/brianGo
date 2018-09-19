package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func main() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))




	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0],":", pair[1])
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	//检查 Scan 的错误。文件结束符是可以接受的，并且不会被 Scan 当作一个错误。

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
