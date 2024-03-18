package main

import "runtime/debug"

func main() {
	type myString string

	var s myString
	s = "this is a my String type"

	var i1, ok = interface{}(s).(myString)

	println(i1)
	println(ok)

	gonum := debug.SetMaxThreads(10)

	println(gonum)

}
