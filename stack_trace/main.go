package main

func main() {
	example(make([]string, 2, 4), "hello", 10)
}

//go:noinline
func example(slice []string, str string, i int) {
	panic("Want stack trace!")
}

/*
panic: Want stack trace!

goroutine 1 [running]:
main.example({0x1400108?, 0x60?, 0x10b6860?}, {0xc000020060?, 0x0?}, 0xc0000021a0?)
        /Users/LittleYellowDog/project/pprof-practice/stack_trace/main.go:9 +0x27
main.main()
        /Users/LittleYellowDog/project/pprof-practice/stack_trace/main.go:4 +0x52
exit status 2
*/
