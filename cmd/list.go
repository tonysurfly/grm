package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const listPattern = "%-40s %-20s %-20s %s\n"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgs, err := loadAllInstalledFromDB()
		if err != nil {
			return err
		}
		if len(pkgs) > 0 {
			fmt.Printf(listPattern, "Package", "Version", "Locked", "Filter")
			for _, p := range pkgs {
				fmt.Printf(
					listPattern,
					p.Owner+"/"+p.Repo,
					p.Version,
					p.Locked,
					fmt.Sprintf(strings.Join(p.Filter, ", ")),
				)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
