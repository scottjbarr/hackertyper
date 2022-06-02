package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	delay := flag.Int("delay", 1, "Delay in ms per character")
	flag.Parse()

	lines := []string{}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	data := strings.Join(lines, "\n")

	r := strings.NewReader(data)

	in := []byte(data)
	d := time.Duration(time.Millisecond * time.Duration(*delay))

	for i := 0; i < len(in); i++ {
		b, err := r.ReadByte()
		if err != nil {
			fmt.Printf("%c\n", err)
			return
		}

		time.Sleep(d)
		fmt.Printf("%s", string(b))
	}

	fmt.Printf("\n\n")
}
