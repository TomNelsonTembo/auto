/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package close

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// exitCmd represents the exit command
var MustcloseCmd = &cobra.Command{
	Use:   "Auto Close",
	Short: "Automates closing of all Apllications",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func run_close() {
	python := path.Clean(strings.Join([]string{os.Getenv("userprofile"), "Anaconda3", "python.exe"}, "/"))
	script := "close.py"
	cmd := exec.Command("cmd", python, script)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	run_close()
}
