package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var exportCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
	refDataPath  string
}{}

var exportCMD = &cobra.Command{
	Use:   "export",
	Short: "export the data of storage server",
	Long: `export the data of storage that related to current system
	<warn>	No Storage Server for this project 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver Static Site Generator v0.9 -- HEAD")
		// if len(args) > 0 {
		// 	fmt.Println(args)
		// }
		// // fmt.Println(exportCMDInput.cfPath)
		// var configPoint *common.ConfigTemp
		// var err error
		// if strings.Contains(exportCMDInput.cfPath, ".toml") {
		// 	configPoint, err = Cf.OpenToml(exportCMDInput.cfPath)
		// } else if strings.Contains(exportCMDInput.cfPath, ".yaml") {
		// 	configPoint, err = Cf.OpenYaml(exportCMDInput.cfPath)
		// }
		// log.Println(configPoint)
		// log.Println(exportCMDInput.mode)
		// log.Println("refDataPath:", exportCMDInput.refDataPath)
		// if err == nil {
		// 	Wb.ServerExportProc(configPoint, &exportCMDInput.refDataPath)
		// } else {
		// 	panic(err)
		// }
	},
}

func init() {
	callPath, _ := os.Getwd()
	exportCMD.Flags().StringVarP(
		&exportCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start db server with specific config file")
	exportCMD.Flags().StringVarP(
		&exportCMDInput.refDataPath,
		"export", "x",
		filepath.Join(callPath, "/data"),
		"export json file for mongodb")
	rootCmd.AddCommand(exportCMD)
}
