package main

import (
	"github.com/Adem20021920/package/fn"
	"github.com/Adem20021920/package/pkg"
)

func main() {
	me := fn.WhoAmI(pkg.Name, pkg.Age)
	fn.DoSmth(me)

}
