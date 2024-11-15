package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version identifier populated via the CI/CD process.
var Version = "HEAD"

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version of the plugin",
		Run: func(*cobra.Command, []string) {
			fmt.Println(Version)
		},
	}
}