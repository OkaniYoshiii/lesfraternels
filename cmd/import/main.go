package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const CommandName = "import"

var hFlag = flag.Bool("h", false, "print help")

type Args struct {
	Filepath string
	Model    string
}

func main() {
	flag.Parse()
	args := os.Args
	askHelp := *hFlag

	if askHelp == true {
		fmt.Println(HelpMessage(CommandName))
		os.Exit(0)
	}

	if len(args) != 2 {
		fmt.Printf("Error : expected 1 argument but got 0.\n\n%s\n", UsageMessage(CommandName))
	}
}

func PossibleModels() []string {
	possibleModels := [...]string{
		"mods",
		"team-members",
	}

	return possibleModels[:]
}

func BNFMessage(commandName string) string {
	args := Args{
		Filepath: "filepath",
		Model:    "model",
	}

	message := ExampleMessage(commandName, args)

	return message
}

func ModelExampleMessage(commandName string) string {
	models := PossibleModels()
	args := Args{
		Filepath: "/home/user/downloads/mods.json",
		Model:    models[0],
	}

	message := ExampleMessage(commandName, args)

	return message
}

func UsageMessage(commandName string) string {
	bnf := BNFMessage(commandName)
	possibleValues := strings.Join(PossibleModels(), ", ")
	modelExample := ModelExampleMessage(commandName)
	usage := fmt.Sprintf(`Usage : "%s"

Arguments:
  filepath (mandatory) :
    - relative or absolute path to the JSON file you want to import
  model (mandatory) :
    - what type of data you want to import. Possible values : (%s)
    - ex: %q
	`, bnf, possibleValues, modelExample)

	return usage
}

func ExampleMessage(commandName string, args Args) string {
	message := fmt.Sprintf("%s %s %s", commandName, args.Filepath, args.Model)

	return message
}

func HelpMessage(commandName string) string {
	description := `Load data from a JSON file based on a relative/absolute path and store the data in the database.`
	usage := UsageMessage(commandName)

	return description + "\n\n" + usage
}
