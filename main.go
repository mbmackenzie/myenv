package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
)

type Var struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type State struct {
	Vars []Var `json:"vars"`
}

func writeState(state State) {
	file, _ := json.MarshalIndent(state, "", " ")
	_ = ioutil.WriteFile("myenv.json", file, 0644)
}

func readState() State {
	file, _ := ioutil.ReadFile("myenv.json")

	var state State
	_ = json.Unmarshal(file, &state)

	return state
}

func printMainHelp() {
	fmt.Println("myenv manages what variables and files are required in a local environment.")
	fmt.Println("")
	fmt.Println("  myenv COMMAND [RESOURCE] [ARG ...]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  list                Display resource names")
	fmt.Println("  get                 Display resource name and value")
	fmt.Println("  create              Create a resource")
	fmt.Println("  local               Add resource to local environment")
	fmt.Println("")
	fmt.Println("Resources:")
	fmt.Println("  var                 Environment variables")
	fmt.Println("  ref                 File references")
	fmt.Println("  config              Configuration files")
}

func printGetHelp() {
	fmt.Println("Get a resource")
	fmt.Println()
	fmt.Println("  Usage: myenv get RESOURCE NAME")
	fmt.Println()
}

func printCreateHelp() {
	fmt.Println("Create a new resource")
	fmt.Println()
	fmt.Println("  Usage: myenv create RESOURCE NAME [VALUE]")
	fmt.Println()
}

func main() {
	state := readState()
	arguments := os.Args

	if len(arguments) == 1 {
		printMainHelp()
		return
	}

	switch arguments[1] {
	case "list":
		if len(arguments) < 3 {
			fmt.Println("Show list help...")
			return
		}

		switch arguments[2] {
		case "vars", "var":

			namesOnly := false

			for _, arg := range arguments[3:] {
				if arg == "--names-only" {
					namesOnly = true
				}
			}

			w := tabwriter.NewWriter(os.Stdout, 16, 1, 1, ' ', 0)
			fmt.Fprintln(w, "NAME\t\tCHAR LEN\tCREATED\tMODIFIED")
			for _, v := range state.Vars {
				numChars := len(v.Value)
				fmt.Fprintf(w, "%s\t\t%d\t----\t----\n", v.Name, numChars)
				if namesOnly {
					fmt.Println(v.Name)
				}
			}

			if !namesOnly {
				w.Flush()
			}
		default:
			fmt.Printf("Unknown resource: %s\n", arguments[2])
		}
	case "get":
		if len(arguments) < 4 {
			printGetHelp()
			return
		}

		switch arguments[2] {
		case "var", "vars":

			for _, arg := range arguments[3:] {

				found := false

				for _, v := range state.Vars {
					if v.Name == arg {
						found = true
						fmt.Println(v.Value)
					}
				}

				if !found {
					fmt.Printf("Unknown variable: %s\n", arg)
				}

				found = false
			}

		default:
			fmt.Printf("Unknown resource: %s\n", arguments[2])
		}
	case "create":
		if len(arguments) < 3 {
			printCreateHelp()
			return
		}

		switch arguments[2] {
		case "var":
			if len(arguments) != 5 {
				fmt.Println("myenv var create NAME VALUE")
				return
			}

			name := arguments[3]
			value := arguments[4]

			for _, v := range state.Vars {
				if name == v.Name {
					fmt.Println("Variable already exists - use update instead.")
					return
				}
			}

			newVar := Var{name, value}
			state.Vars = append(state.Vars, newVar)
		default:
			printCreateHelp()
			return
		}

		writeState(state)
	default:
		printMainHelp()
	}

}
