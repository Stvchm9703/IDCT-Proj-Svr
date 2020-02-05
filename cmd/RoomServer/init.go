package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
	refDataPath  string
	schemaPath   string
	isExport     bool
}{}

var initCMD = &cobra.Command{
	Use:   "init",
	Short: "init the storage server of webserver",
	Long:  `webserver server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver Static Site Generator v0.9 -- HEAD")
		// if len(args) > 0 {
		// 	fmt.Println(args)
		// }
		// // fmt.Println(initCMDInput.cfPath)
		// var configPoint *common.ConfigTemp
		// var err error
		// if strings.Contains(initCMDInput.cfPath, ".toml") {
		// 	configPoint, err = Cf.OpenToml(initCMDInput.cfPath)
		// } else if strings.Contains(initCMDInput.cfPath, ".yaml") {
		// 	configPoint, err = Cf.OpenYaml(initCMDInput.cfPath)
		// }
		// log.Println(configPoint)
		// log.Println(initCMDInput.mode)
		// if err == nil {

		// 	if initCMDInput.isExport {
		// 		Wb.ServerInitProc(configPoint, &initCMDInput.schemaPath)
		// 	} else {
		// 		Wb.ServerInitProc(configPoint, nil)
		// 	}

		// } else {
		// 	panic(err)
		// }
	},
}

func init() {
	callPath, _ := os.Getwd()
	initCMD.Flags().StringVarP(
		&initCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start server with specific config file")

	initCMD.Flags().BoolVarP(
		&initCMDInput.isExport,
		"exportSchema", "x",
		false,
		"export the createCollecction validator.$jsonSchema build json")
	initCMD.Flags().StringVarP(
		&initCMDInput.schemaPath,
		"schema", "m",
		filepath.Join(callPath, "doc", "log", "create_schema.json"),
		"start server with specific config file")
	rootCmd.AddCommand(initCMD)
}
