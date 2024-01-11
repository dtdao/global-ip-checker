package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var getCurrentGlobalIpCmd = &cobra.Command{
	Use:   "current",
	Short: "Get your current global ip",
	Run:   getCurrentGlobalIp,
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(getCurrentGlobalIpCmd)
}

func CurrentGlobalIp() (string, error) {
	response, error := http.Get(ipify)
	if error != nil {
		return "", error
	}

	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)

	if error != nil {
		return "", error
	}

	var globalIp PublicIp

	error = json.Unmarshal(body, &globalIp)

	if error != nil {
		return "", error
	}

	return globalIp.IP, nil
}

func getCurrentGlobalIp(cmd *cobra.Command, args []string) {
	globalIp, error := CurrentGlobalIp()
	if error != nil {
		fmt.Println("Error fetching the global ip address: ", error)
	}

	fmt.Println("Your current global ip is: ", globalIp)
	return

}
