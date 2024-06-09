package main

import (
	"os"
	"text/template"
)

type Mine struct {
	Location string
}

func main() {
	// Define the data to be used in the template
	mine := Mine{
		Location: "🇺🇸 Austin, TX, US",
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("text/README.tmpl")
	if err != nil {
		panic(err)
	}

	// Create a README file
	file, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Execute the template and write to the README file
	err = tmpl.Execute(file, mine)
	if err != nil {
		panic(err)
	}

	println("README.md generated successfully.")
}
