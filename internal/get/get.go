package get

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/BenLee8602/repodrum/internal/file"
)

func Clone(d file.Dep) {
	cmd := exec.Command("git", "clone", d.Url, "dependencies/"+d.Name)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error cloning \"%s\": ", d.Url)
		log.Fatal(err)
		os.Exit(1)
	}
}
