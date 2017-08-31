package branchingAndLooping

import (
	"fmt"
)

// func main() {
// 	loops()
// 	//switches()
// 	//ifs()
// }

func loops() {
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["thrid"] = "buz"

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("--------------------------")

	s := []string{"foo", "bar", "buz"}

	for idx, v := range s {
		fmt.Println(idx, v)
	}

	fmt.Println("--------------------------")

	i := 0
	for {
		i++
		fmt.Println(i)

		if i > 5 {
			break
		}
	}

	fmt.Println("--------------------------")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func switches() {
	foo := 5

	switch {
	case foo == 1:
		fmt.Println("one")
	case foo > 2:
		fmt.Println("two")
	}

	switch foo := 1; foo {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("one")
	}
}

func ifs() {
	if foo := 2; foo == 1 {
		fmt.Println("bar")
	} else {
		fmt.Println("buz")
	}
}
