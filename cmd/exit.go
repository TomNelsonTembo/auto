/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// exitCmd represents the exit command
var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Automates closing of all Apllications",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		run_close()
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

func init() {
	rootCmd.AddCommand(exitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
