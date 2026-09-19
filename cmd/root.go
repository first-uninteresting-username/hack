// SPDX-FileCopyrightText: 2026 first-uninteresting-username
//
// SPDX-License-Identifier: GPL-3.0-only
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "5",
	Use:     "hack",
	Short:   "CLI tool for interacting with LLMs",
	Long: `
Interact with LLMs from the command line

hack is a simple tool for interacting with LLMs.
It is made to be scriptable, extensible and easy to use.
There're no agentic capabilities built in,
but because of how it works, it's possible to create an agent based on it.

Example usage:
	hack -p "your prompt here"
	echo "some content" | hack -p "summarize this"
	ls -la | hack -sp "delete the largest file"

Modes:
	shell   (-s/--shell)    Generate shell commands from a prompt
	code    (-w/--write)    Output executable code (jq, python3, bash, or POSIX sh)
	normal                  Standard prompt-and-response
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runHack()
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runHack() error {
	response, err := makeRequest()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
