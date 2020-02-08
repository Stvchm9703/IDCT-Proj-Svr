package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	// Wb "webserver/server"
	Cf "RoomStatus/config"
	// "webserver/server/common"

	"github.com/spf13/cobra"
)

var runCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
}{}

var runCmd = &cobra.Command{
	Use:   "start",
	Short: "start the all service of grpc server",
	Long:  `grpc server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc server  Generator v0.9 -- HEAD")
		if len(args) > 0 {
			fmt.Println(args)
		}
		callPath, _ := os.Getwd()
		fmt.Println(runCMDInput.cfPath)

		var configPoint *Cf.ConfTmp
		var err error
		if strings.Contains(runCMDInput.cfPath, ".toml") {
			configPoint, err = Cf.OpenToml(runCMDInput.cfPath)
		} else if strings.Contains(runCMDInput.cfPath, ".yaml") {
			configPoint, err = Cf.OpenYaml(runCMDInput.cfPath)
		}
		log.Println(configPoint)
		log.Println(runCMDInput.mode)
		if err == nil {
			Wb.ServerMainProcess(configPoint, callPath, runCMDInput.mode)
		} else {
			panic(err)
		}
	},
}

func init() {
	callPath, _ := os.Getwd()

	runCmd.Flags().StringVarP(
		&runCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start server with specific config file")

	runCmd.Flags().StringVarP(
		&runCMDInput.mode,
		"mode", "m",
		"prod",
		"server running mode [prod / dev / test]")

	rootCmd.AddCommand(runCmd)
}
