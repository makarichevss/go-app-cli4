/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:
 
 Bash:
 
 $ source <(yourprogram completion bash)
 
 # To load completions for each session, execute once:
 Linux:
  $ yourprogram completion bash > /etc/bash_completion.d/yourprogram
 MacOS:
  $ yourprogram completion bash > /usr/local/etc/bash_completion.d/yourprogram
 
 Zsh:
 
 # If shell completion is not already enabled in your environment,
 # you will need to enable it.  You can execute the following once:
 
 $ echo "autoload -U compinit; compinit" >> ~/.zshrc
 
 # To load completions for each session, execute once:
 $ yourprogram completion zsh > "${fpath[1]}/_yourprogram"
 
 # You will need to start a new shell for this setup to take effect.
 
 Fish:
 
 $ yourprogram completion fish | source
 
 # To load completions for each session, execute once:
 $ yourprogram completion fish > ~/.config/fish/completions/yourprogram.fish

PowerShell:

  PS> %[1]s completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> %[1]s completion powershell > %[1]s.ps1
  # and source this file from your PowerShell profile.
 `,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			if err := cmd.Root().GenBashCompletion(os.Stdout); err != nil {
				fmt.Println("Unable to generate bash completion script:", err)
			}
		case "zsh":
			if err := cmd.Root().GenZshCompletion(os.Stdout); err != nil {
				fmt.Println("Unable to generate zsh completion script:", err)
			}
		case "fish":
			if err := cmd.Root().GenFishCompletion(os.Stdout, true); err != nil {
				fmt.Println("Unable to generate fish completion: ", err)
			}
		case "powershell":
			if err := cmd.Root().GenPowerShellCompletion(os.Stdout); err != nil {
				fmt.Println("Unable to generate powershell completion: ", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
