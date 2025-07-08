package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type CommandEntry struct {
	Command     string   `json:"command"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

var storagePath = filepath.Join(os.Getenv("HOME"), ".term-cli", "commands.json")

func ensureStorage() map[string]CommandEntry {
	_ = os.MkdirAll(filepath.Dir(storagePath), 0755)
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		_ = ioutil.WriteFile(storagePath, []byte("{}"), 0644)
	}
	data, err := ioutil.ReadFile(storagePath)
	if err != nil {
		log.Fatalf("Failed to read storage: %v", err)
	}
	var cmds map[string]CommandEntry
	_ = json.Unmarshal(data, &cmds)
	return cmds
}

func saveStorage(cmds map[string]CommandEntry) {
	data, _ := json.MarshalIndent(cmds, "", "  ")
	_ = ioutil.WriteFile(storagePath, data, 0644)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "term",
		Short: "Save, manage, and run your frequently used terminal commands easily.",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "save [name] [command]",
		Short: "Save a new command",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			commandStr := strings.Join(args[1:], " ")

			cmds := ensureStorage()
			cmds[name] = CommandEntry{
				Command: commandStr,
			}
			saveStorage(cmds)
			fmt.Println("Saved:", name)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List saved commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmds := ensureStorage()
			for name, entry := range cmds {
				fmt.Printf("%s: %s\n", name, entry.Command)
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "run [name] [args...]",
		Short: "Run a saved command with arguments",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			runArgs := args[1:]

			cmds := ensureStorage()
			entry, ok := cmds[name]
			if !ok {
				fmt.Println("Command not found:", name)
				return
			}

			expanded := entry.Command
			for _, arg := range runArgs {
				expanded = strings.Replace(expanded, "{}", arg, 1)
			}

			fmt.Println("Running:", expanded)
			parts := strings.Fields(expanded)
			c := exec.Command(parts[0], parts[1:]...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Stdin = os.Stdin
			if err := c.Run(); err != nil {
				fmt.Println("Error:", err)
			}
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

