/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/dmark/reimagined-funicular/todo"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo",
	Long:  `Add will create a new todo item on the list.`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(datafile)
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority) // my priority values are not being saved. They are defaulting to zero.
		items = append(items, item)
	}
	err = todo.SaveItems(datafile, items)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}
