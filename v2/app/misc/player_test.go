package misc

func testP(name string) Player {
	return New(&Options{For2P: true, Name: name})
}
