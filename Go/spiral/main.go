package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nCommands,_ := reader.ReadString('\n')
	nCommands = strings.TrimSuffix(nCommands, "\n")
	q,_ := strconv.Atoi(nCommands)

	commands := make([]string, q)
	for i := 0; i < q; i++ {
		command,_ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")
		commands[i] = command
	}

	for i := 0; i < q; i++ {
		var c string
		var n int
		fmt.Sscanf(commands[i], "%s %d", &c, &n)
		if c == "list" {
			list(n)
		}else {
			item(n - 1)
		}
	}
}

func list(n int)  {
	// 0 right 1 up 2 left 3 down
	x, y, direction, amount := 0, 0, 0, 1

	fmt.Printf("%d %d\n", x, y)

	i := 0
	f := 1
	for {
		amount = i/2 + 1
		for j := 0; j < amount; j++ {
			switch direction {
			case 0:
				x++
			case 1:
				y++
			case 2:
				x--
			case 3:
				y--
			}

			fmt.Printf("%d %d\n", x, y)
			f++
			if f == n {
				return
			}
		}

		i++

		direction++
		direction %= 4
	}
}

func item(n int) {
	x, y := 0, 0
	i := 0
	for {
		if (i+1)*(i+2) > n {
			break
		}

		i++
	}

	var r int
	if i%2 == 0 {
		r = hesabi(i/2, 1)
	} else {
		r = hesabi(i/2+1, 1)
	}

	l := hesabi(i/2, 2)

	x = x + r - l
	y = y + r - l

	var direction int
	if i%2 == 0 {
		direction = 0
	}else {
		direction = 2
	}

	// 0 right 1 up 2 left 3 down
	f := i * (i + 1)
	for {
		var d int
		if n - f > i + 1 {
			d = i + 1
		}else {
			d = n - f
		}

		switch direction {
		case 0:
			x += d
		case 1:
			y += d
		case 2:
			x -= d
		case 3:
			y -= d
		}

		f += d
		if f == n {
			fmt.Printf("%d %d\n", x, y)
			return
		}

		direction++
	}
}

func hesabi(n int, a int) int {
	return n * (a + n - 1)
}