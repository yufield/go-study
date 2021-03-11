package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(strconv.Quote(`{"$db" : "local","collStats" : "startup_log","lsid" : {"id" : {"Subtype" : 4,"Data" : "HUHuhGSeS16N+nEaiYkk5w=="}}}`))
}
