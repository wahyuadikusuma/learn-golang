package main
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.prntln("You are blocked,", name)
	} else {
		fmt.println("Welcome,", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("eko", func(name string) bool {
		return name == "root"
	})
}