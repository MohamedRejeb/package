package main

import (
	"github.com/adembc/package/fn"
	"github.com/adembc/package/pkg"
)

func main() {
	me := fn.WhoAmI(pkg.Name, pkg.Age)
	fn.DoSmth(me)

}
