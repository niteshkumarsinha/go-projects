package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type CmdFlags struct {
	Add string `flag:"add"`
	Del int `flag:"del"`
	Toggle int `flag:"toggle"`
	List bool `flag:"list"`
	Edit string `flag:"edit"`
}

func NewCommandFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo")
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) error {
	if cf.Add != "" {
		todos.Add(cf.Add)
		return nil
	}
	if cf.Del != -1 {
		return todos.Delete(cf.Del)
	}
	if cf.Toggle != -1 {
		return todos.Toggle(cf.Toggle)
	}
	if cf.List {
		todos.List()
		return nil
	}
	if cf.Edit != "" {
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid edit format")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid edit format")
			os.Exit(1)
		}
		return todos.Edit(index, parts[1])
	}
	fmt.Println("invalid command")
	os.Exit(1)	
	return nil	
}

