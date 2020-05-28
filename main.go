package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	n := flag.Int("n", 1, "the number of rows")
	sep := flag.String("d", "\t", "separator")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	buf := make([][]string, *n)

	for i := 0; s.Scan(); i = (i + 1) % *n {
		buf[i] = append(buf[i], s.Text())
	}

	for _, x := range buf {
		fmt.Println(strings.Join(x, *sep))
	}

}
