package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	Cf "RoomStatus/pkg/config"
	wb "RoomStatus/pkg/serverctlNoRedis"

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
	Short: "start the server of grpc server",
	Long:  `grpc server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc server  Generator v0.9 -- HEAD")
		if len(args) > 0 {
			fmt.Println(args)
		}
		// , _ = os.Getwd()
		fmt.Println(runCMDInput.cfPath)

		var configPoint *Cf.ConfTmp
		var err error
		if strings.Contains(runCMDInput.cfPath, ".toml") {
			configPoint, err = Cf.OpenToml(runCMDInput.cfPath)
		} else if strings.Contains(runCMDInput.cfPath, ".yaml") {
			configPoint, err = Cf.OpenYaml(runCMDInput.cfPath)
		}
		fmt.Printf("\n %#v\n", configPoint)
		fmt.Println(runCMDInput.mode)
		if err == nil {
			// Wb.ServerMainProcess(configPoint, callPath, runCMDInput.mode)
			wb.ServerMainProcess(configPoint)
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
