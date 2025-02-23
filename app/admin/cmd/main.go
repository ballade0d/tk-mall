package main

func main() {
	err, _ := wireApp()
	if err != nil {
		panic(err)
	}
}
