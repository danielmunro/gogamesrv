package main

import "github.com/danielmunro/gogamesrv"

func main() {
	s := gogamesrv.NewServer()
	s.Listen(5555)
}
