package tinyadd

func Add(args ...int) (res int) {
	for _, arg := range args {
		res += arg
	}
	return
}
