package main

func main() {
	_, err := wireApp()
	if err != nil {
		panic(err)
	}
}
