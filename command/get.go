package command

import (
	"log"

	"github.com/Mtt6300/workspace-tunnel/core/get"
	"github.com/spf13/cobra"
)

func init() {
	// getCmd represents the get command
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "show available service to port forward",
		Run: func(cmd *cobra.Command, args []string) {
			resourceFlag, err := cmd.Flags().GetString("resource")
			if err != nil {
				log.Fatal(err)
			}

			kubeConfigFlag, err := cmd.Flags().GetString("kube-config")
			if err != nil {
				log.Fatal(err)
			}

			get.GetCommand(resourceFlag, kubeConfigFlag)
		},
	}

	rootCmd.AddCommand(getCmd)

	getCmd.Flags().String("resource", "", "kubernetes resource")
	getCmd.MarkFlagRequired("resource")
	getCmd.Flags().String("kube-config", "", "kubnerenets yaml config file path")
	getCmd.MarkFlagRequired("kube-config")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
