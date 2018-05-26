package main;

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{ "Red", "Green", "Blue" };
	fmt.Printf("%v:%T \n",colors,colors);

	colors = append(colors, "Purple");
	fmt.Printf("%v:%T \n",colors,colors);

	colors = append(colors[1:len(colors)]);
	fmt.Printf("%v:%T \n",colors,colors);

	numbers := make([]int, 5, 5); // size 5 capacity 5
	for i:=0 ; i<5 ; i++ {
		numbers[i] = i-1000*i;
	}

	fmt.Println(numbers);
	sort.Ints(numbers);
	fmt.Println(numbers);

}