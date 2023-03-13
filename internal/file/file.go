package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dep struct {
	Name string
	Url  string
}

func DepFromString(s string) Dep {
	var d Dep
	fields := strings.Split(s, " ")
	d.Name = fields[0]
	d.Url = fields[1]
	return d
}

func DepToString(d Dep) string {
	s := d.Name + " " + d.Url + "\n"
	return s
}

func Read() []Dep {
	file, err := os.OpenFile("dependencies.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening dependencies.txt: ", err)
		os.Exit(1)
	}
	defer file.Close()

	var deps []Dep
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		deps = append(deps, DepFromString(scanner.Text()))
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading dependencies.txt: ", err)
		os.Exit(1)
	}

	return deps
}

func Write(deps []Dep) {
	file, err := os.Create("dependencies.txt")
	if err != nil {
		fmt.Println("Error opening dependencies.txt: ", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, d := range deps {
		writer.WriteString(DepToString(d))
	}
	writer.Flush()
}
