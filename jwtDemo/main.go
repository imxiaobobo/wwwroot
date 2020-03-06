package main

import "fmt"

type Keyfunc func(int int) string

func New(k Keyfunc) {
	var a = 1
	s := k(a)
	fmt.Println(s)
}

func main() {
	New(func(i int) string {
		return "wb"
	})
	//r := gin.Default()
	//r = router.UserRouter(r)
	//r.Run()

}
