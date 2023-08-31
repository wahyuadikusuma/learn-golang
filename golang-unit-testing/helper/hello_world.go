package helper

// jangan lupa untuk menggunakan capital di awal nama function
func HelloWorld(name string) string {
	return "Hello " + name
}

/**
aturan dalam unit testing:
- jika ingin membuat function unit testing, awali dengan test. Misal: TestHelloWorld
- harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value
*/