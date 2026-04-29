package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "luluka sample/shader.kage",
	Short: "View and tweak shaders made with Kage for Ebiten",
	Long: `Luluka helps you display Ebiten shaders in Kage.
Simply load a shader by indicating its path.
Optionally specify uniforms using -u.
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uniformArgs, err := cmd.Flags().GetStringSlice("uniform")
		if err != nil {
			log.Panicf("Failed to get uniform args: %s", uniformArgs)
		}

		Run(args[0])
	},
}

func init() {
	rootCmd.Flags().StringSliceP("uniform", "u", []string{}, "specifies a uniform value")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
