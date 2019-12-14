package main

import (
	"os"
	"strconv"

	"github.com/kissgyorgy/adventofcode2019/aoc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code helper",
	Long: `A Command Line helper for Advent of Code.
it is able to download input (well, that's it so far 😄)`,
}

var getInputCmd = &cobra.Command{
	Use:   "get-input [flags] day",
	Short: "Download input from website and save in the current directory.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.PrintErrln("You need to specify a day!\n")
			cmd.Usage()
			os.Exit(1)
		}
		day, err := strconv.ParseInt(args[0], 10, 0)
		if err != nil {
			cmd.PrintErrf("%q is not a day (should be int between 1-25)\n", args[0])
			os.Exit(1)
		}
		aoc.GetInput(int(day))
	},
}

func init() {
	rootCmd.AddCommand(getInputCmd)
}

func main() {
	rootCmd.Execute()
}
