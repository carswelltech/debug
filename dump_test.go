package debug

import (
	`testing`
	`fmt`
	`os`
)


func TestSDump(t *testing.T) {
	fmt.Println(SDump("This is a test"))
	fmt.Println(SDump([]int{1,2,3,4,5}))
	fmt.Println(SDump(map[int]int{1:10,2:20,3:30,4:40,5:50}))
	fmt.Println(SDump(struct{A int; B int}{A:1, B:2}))

	Dump(struct{A int; B int}{A:1, B:2},os.Stderr)
	Dump(struct{A int; B int}{A:1, B:2})
	Dump(struct{A int; B int}{A:1, B:2},os.Stderr,os.Stdout)
}
