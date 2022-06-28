package main

import "github.com/takei0107/karasu"

func main() {
	if err := karasu.Run(); err != nil {
		panic(err)
	}
}
