package main

import (
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/application"
	"github.com/mikolajsemeniuk/go-elasticsearch-react-fullstack/data"
)

func main() {
	data.GetInfo()
	application.Listen()
}
