package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	log.Println("hello")
	command := &cobra.Command{
		Use:   "run",
		Short: "run my profiler",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
