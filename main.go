package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello Lango!")
	fmt.Println("------------")

    context = make(map[string]result)

	for {
        fmt.Print("(lango) : ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

        index = 0
        parse()
        for CurTok.tocat != TokenCats[EOF] {
            Advance()
        }
        Advance()

        eval()
	}
}
