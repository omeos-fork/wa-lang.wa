package main

func main() {
	println(fib(46))
}
func fib(n int) int {
	var aux func(n, acc1, acc2 int) int
	aux = func(n, acc1, acc2 int) int {
		switch n {
		case 0:
			return acc1
		case 1:
			return acc2
		default:
			return aux(n-1, acc2, acc1+acc2)
		}
	}
	return aux(n, 0, 1)
}
