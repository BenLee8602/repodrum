package cli

import (
	"fmt"
	"os"

	"github.com/BenLee8602/repodrum/internal/file"
	"github.com/BenLee8602/repodrum/internal/get"
)

func Install(args []string) {
	os.RemoveAll("dependencies/")
	deps := file.Read()
	for _, d := range deps {
		fmt.Printf("Cloning %s from %s ...\n", d.Name, d.Url)
		get.Clone(d)
	}
	fmt.Println("Installed all dependencies")
}

func Uninstall() {
	os.RemoveAll("dependencies/")
	fmt.Println("Uninstalled all dependencies")
}

func Add(args []string) {
	if len(args) < 2 {
		fmt.Println("add: must provide a git repo url and a name")
		os.Exit(1)
	}
	curName := args[0]
	curUrl := args[1]

	deps := file.Read()
	for _, d := range deps {
		if d.Name == curName {
			fmt.Printf("add: dependency already exists called %s\n", curName)
			os.Exit(1)
		} else if d.Url == curUrl {
			fmt.Printf("add: dependency already installed as %s\n", d.Name)
			os.Exit(1)
		}
	}

	d := file.Dep{Name: curName, Url: curUrl}
	get.Clone(d)
	deps = append(deps, d)
	file.Write(deps)

	fmt.Printf("Added dependency %s from %s\n", d.Name, d.Url)
}

func Remove(args []string) {
	if len(args) == 0 {
		fmt.Println("remove: must provide a package to remove")
		os.Exit(1)
	}
	curName := args[0]

	deps := file.Read()
	for i, d := range deps {
		if d.Name == curName {
			deps = append(deps[:i], deps[i+1:]...)
			os.RemoveAll("dependencies/" + d.Name)
			file.Write(deps)
			fmt.Printf("Removed dependency %s (%s)\n", d.Name, d.Url)
			return
		}
	}

	fmt.Printf("remove: dependency %s not found\n", curName)
	os.Exit(1)
}
