package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// CLI flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addURL := addCmd.String("url", "", "The original long URL")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getSlug := getCmd.String("slug", "", "The short slug to resolve")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'get' or 'list' subcommands")
		os.Exit(1)
	}

	store := NewURLStore("urls.json") // Pass the filename to store the data

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addURL == "" {
			fmt.Println("Please provide a URL with -url")
			os.Exit(1)
		}
		slug := store.Add(*addURL)
		fmt.Printf("Shortened URL: %s\n", slug)
	case "get":
		getCmd.Parse(os.Args[2:])
		if *getSlug == "" {
			fmt.Println("Please provide a slug with -slug")
			os.Exit(1)
		}
		url, ok := store.Get(*getSlug)
		if !ok {
			fmt.Println("Slug not found.")
		} else {
			fmt.Printf("Original URL: %s\n", url)
		}
	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Stored URLs:")
		for slug, long := range store.All() {
			fmt.Printf("%s → %s\n", slug, long)
		}
	default:
		fmt.Println("expected 'add', 'get' or 'list'")
		os.Exit(1)
	}
}
