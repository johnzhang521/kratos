package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/urfave/cli"
)

func runNew(ctx *cli.Context) error {
	if len(ctx.Args()) == 0 {
		return errors.New("project name required, please enter a project name you want to create")
	}
	p.Name = ctx.Args()[0]

	if p.ModuleName == "" {
		p.ModuleName = p.Name
	}

	if p.Path != "" {
		p.Path = path.Join(p.Path, p.Name)
	} else {
		pwd, _ := os.Getwd()
		p.Path = path.Join(pwd, p.Name)
	}
	p.Path = filepath.FromSlash(p.Path)
	// Create a project
	if err := create(); err != nil {
		return err
	}
	fmt.Printf("Project: %s\n", p.Name)
	fmt.Printf("Owner: %s\n", p.Owner)
	fmt.Printf("Module Name: %s\n", p.ModuleName)
	fmt.Printf("WithGRPC: %t\n", p.WithGRPC)
	fmt.Printf("Directory: %s\n\n", p.Path)
	fmt.Println("The application has been created.")
	return nil
}
