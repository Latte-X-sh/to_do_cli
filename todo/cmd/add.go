/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	tododata "to_do_cli/todo/todo_data"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
// var addCmd = &cobra.Command{
// 	Use:   "add",
// 	Short: "Add a new todo",
// 	Long:  `Add will create a new todo item to the list.`,
// 	Run:   addRun,
// }

func init() {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new todo",
		Long:  `Add will create a new todo item to the list.`,
		Run:   addRun,
	}
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addRun(cmd *cobra.Command, args []string) {
	// [] string represents an array | slice
	//  so the for loop will loop over 2 values
	//  The index of the array
	var myTask tododata.Item
	// myTask.Text
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
	// tutorial implementaion
	// good for iterating over channels, slices, maps arrays etc
	// for index, value := range args {
	// 	fmt.Println(x, z)
	// }

}
