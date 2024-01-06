package command

import (
	"log"

	"github.com/Mtt6300/workspace-tunnel/core/up"
	"github.com/spf13/cobra"
)

func init() {
	var upCmd = &cobra.Command{
		Use:   "up",
		Short: "port-foward specified services in config",
		Run: func(cmd *cobra.Command, args []string) {
			nameFlag, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatal(err)
			}

			configFlag, err := cmd.Flags().GetString("config")
			if err != nil {
				log.Fatal(err)
			}

			onlyFlag, err := cmd.Flags().GetStringSlice("only")
			if err != nil {
				log.Fatal(err)
			}

			up.UpCommand(nameFlag, configFlag, onlyFlag)
		},
	}

	rootCmd.AddCommand(upCmd)

	upCmd.Flags().String("name", "", "workspace name")
	upCmd.MarkFlagRequired("name")
	upCmd.Flags().String("config", "", "config path")
	upCmd.MarkFlagRequired("config")
	upCmd.Flags().StringSlice("only", []string{}, "port forward only specified resources")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
