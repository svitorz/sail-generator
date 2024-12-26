package main

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "sail-generator"
  Short: "Generate Laravel Sail projects"
  Long: `Generate Laravel Sail projects automaticly`
  Run: func(cmd *cobra.Command, args []string){ cmd.Help() },
},

