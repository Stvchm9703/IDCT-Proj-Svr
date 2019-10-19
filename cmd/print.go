package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cfPath string

// serveCmd represents the serve command
var PrintCfCMD = &cobra.Command{
	Use:   "print",
	Short: "print the building config",
	Long:  `print the import building config.toml (default : the argurment directory)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		fmt.Println(cfPath)
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(PrintCfCMD)

	callPath, _ := os.Getwd()
	PrintCfCMD.Flags().StringVarP(
		&cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"import building config file")

}
