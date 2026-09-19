// SPDX-FileCopyrightText: 2026 first-uninteresting-username
//
// SPDX-License-Identifier: GPL-3.0-only
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	prompt      string
	commandMode bool
	codeMode    bool
)

func createConfig(cfgPath string) error {
	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("toml")

		configDir, err := os.UserConfigDir()
		if err != nil {
			if home, hErr := os.UserHomeDir(); hErr == nil {
				configDir = filepath.Join(home, ".config")
			} else {
				configDir = "."
			}
		}
		viper.AddConfigPath(filepath.Join(configDir, "hack"))

		viper.AddConfigPath("/etc/hack")
	}

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if errors.As(err, &notFound) {
			return nil
		}
		return fmt.Errorf("reading config: %w", err)
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		"config file path, $HOME/.config/hack/config.toml if not provided")

	rootCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "prompt for the LLM")
	rootCmd.Flags().StringP("key", "k", "", "API key for selected provider")
	rootCmd.Flags().StringP("key-file", "f", "", "Path to file containing API key")
	rootCmd.Flags().StringP("model", "m", "", "LLM used for response")
	rootCmd.Flags().StringP("base", "b", "", "base URL for the API (without /chat/completions)")
	rootCmd.Flags().BoolVarP(&commandMode, "shell", "s", false, "enable command mode")
	rootCmd.Flags().BoolVarP(&codeMode, "write", "w", false, "enable code mode")

	viper.BindPFlag("api_key", rootCmd.Flags().Lookup("key"))
	viper.BindPFlag("api_key_path", rootCmd.Flags().Lookup("key-file"))
	viper.BindPFlag("model", rootCmd.Flags().Lookup("model"))
	viper.BindPFlag("base_url", rootCmd.Flags().Lookup("base"))

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return createConfig(cfgFile)
	}
}

func determineMode(command, code bool) (string, error) {
	if command && code {
		return "", fmt.Errorf("cannot use command mode and code mode simultaneously")
	}

	switch {
	case command:
		return "command", nil
	case code:
		return "code", nil
	default:
		return "normal", nil
	}
}
