package main

import "fmt"

const (
	VERSION  = "1.0.0"
	REQARGS  = 2
	USAGE    = "Usage: olins \"command bar | command foo\" name"
	EXAMPLE  = "olins \"curl -X POST {site}/{endpoint} \" post_to_endpoint"
	EXAMPLE1 = "post_to_endpoint http://example.com /api/v1/fooz"
)

func Run(Args []string) {
	if len(Args) < 2 {
		Help()
		return
	}

	var CommandString = Args[1]
	var Name = Args[2]

	var saver = Saver{
		Command: CommandString,
		Name:    Name,
	}

	saver.Save()
}

func Help() {
	fmt.Println("OneLineSaver v" + VERSION + " by @pabloskubert\n")
	fmt.Println(USAGE)
	fmt.Println("Example: \t\t" + EXAMPLE)
	fmt.Println("Then: \t\t\t" + EXAMPLE1)
	fmt.Println("\n")
}
