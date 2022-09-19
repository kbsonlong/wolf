/*
Copyright © 2021 kbsonlong <kbsonlong@gmail.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"

	"github.com/kbsonlong/wolf/internal/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "运维平台",
	Long: `运维管理平台,功能CMDB、Kubernetes控制台、监控系统
竞价实例等.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(cmd.Short)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		r := router.InitRouter()
		r.Run(fmt.Sprintf(":%d", viper.Get("PORT")))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
