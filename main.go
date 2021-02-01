package main

import (
	"309pollution/pollution"
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("USAGE\n" +
		"\t./309pollution n file x y\n\n" +
		"DESCRIPTION\n" +
		"\tn\t number of points on the grid axis\n" +
		"\tfile\tcsv file containing the data points x;y;p\n" +
		"\tx\tabscissa of the point whose pollution level we want to know\n" +
		"\ty\tordinate of the point whose pollution level we want to know")
}

func main() {
	if pollution.CheckHelp() {
		printHelp()
		os.Exit(0)
	}
	err := pollution.CheckArgs(); if err != nil {
		fmt.Printf("%s\n\n", err)
		printHelp()
		os.Exit(84)
	}
	err = pollution.Main(); if err != nil {
		fmt.Printf("%s\n\n", err)
		printHelp()
		os.Exit(84)
	}
}