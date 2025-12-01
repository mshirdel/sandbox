// Package cmd run command for just testing
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _runCmd = &cobra.Command{
	Use:   "run",
	Short: "run some code",
	RunE: func(cmd *cobra.Command, args []string) error {
		run()
		return nil
	},
}

func run() {
	nums := make([]int, 0, 10)
	for i := range 10 {
		nums = append(nums, i)
	}
	fmt.Println(nums)
}
