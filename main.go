package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "laravel-sail-generator",
	Short: "Gera projetos Laravel Sail",
	Long:  `Gera projetos Laravel Sail de forma automatizada.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Cria um novo projeto Laravel Sail",
	Long:  `Cria um novo projeto Laravel Sail com as opções especificadas.`,
	Run:   createProject,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createProject(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's your app name? ")
	appName, _ := reader.ReadString('\n')
	appName = strings.TrimSpace(appName)

	fmt.Print("What are the services do you want to use? (separated by space) ")
	text, _ := reader.ReadString('\n')
	services := strings.ReplaceAll(text, " ", ",")

	curlCmd := fmt.Sprintf("curl -s \"https://laravel.build/%s?with=%s\" | bash",
		appName, services)

	// Executar o comando curl
	output, err := exec.Command("bash", "-c", curlCmd).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
