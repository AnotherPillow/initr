package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

func goProject(name string, code bool, mkdir bool, overwrite bool) {
	if overwrite {
		//remove directory if exists
		err := os.RemoveAll(name)
		if err != nil {
			color.Red("Error removing directory: %s", err)
			return
		}

		color.Green("Removed directory: %s", name)
	}
	if mkdir {
		//check if windows
		if runtime.GOOS == "windows" {

			err := exec.Command("cmd", "/C", "md", name).Run()
			if err != nil {
				color.Red("Error making directory: %s", err)
				return
			}

			color.Green("Made directory: %s", name)
		} else {
			err := exec.Command("mkdir", name).Run()
			if err != nil {
				color.Red("Error making directory: %s", err)
				return
			}

			color.Green("Made directory: %s", name)
		}
	}

	//change CWD with os.Chdir
	if mkdir || overwrite {
		os.Chdir(name)
	}

	err := exec.Command("go", "mod", "init", name).Run()
	if err != nil {
		color.Red("Error making go project: %s", err)
		return
	}

	color.Green("Made go project: %s", name)

	if code {
		err = exec.Command("code", ".").Run()
		if err != nil {
			color.Red("Error opening project in vscode: %s", err)
			return
		}

		color.Green("Opened project in vscode: %s", name)
	}

}

func rustProject(name string, code bool, overwrite bool) {

	if overwrite {
		//remove directory if exists
		err := os.RemoveAll(name)
		if err != nil {
			color.Red("Error removing directory: %s", err)
			return
		}

		color.Green("Removed directory: %s", name)
	}

	err := exec.Command("cargo", "new", name).Run()
	if err != nil {
		color.Red("Error making rust project: %s", err)
		return
	}

	color.Green("Made rust project: %s", name)

	os.Chdir(name)

	if code {
		err = exec.Command("code", ".").Run()
		if err != nil {
			color.Red("Error opening project in vscode: %s", err)
			return
		}

		color.Green("Opened project in vscode: %s", name)
	}

}

func npmProject(name string, code bool, mkdir bool, overwrite bool) {
	if overwrite {
		//remove directory if exists
		err := os.RemoveAll(name)
		if err != nil {
			color.Red("Error removing directory: %s", err)
			return
		}

		color.Green("Removed directory: %s", name)
	}

	if mkdir {
		//check if windows
		if runtime.GOOS == "windows" {

			err := exec.Command("cmd", "/C", "md", name).Run()
			if err != nil {
				color.Red("Error making directory: %s", err)
				return
			}

			color.Green("Made directory: %s", name)
		} else {
			err := exec.Command("mkdir", name).Run()
			if err != nil {
				color.Red("Error making directory: %s", err)
				return
			}

			color.Green("Made directory: %s", name)
		}
	}

	//change CWD with os.Chdir
	if mkdir || overwrite {
		os.Chdir(name)
	}

	err := exec.Command("npm", "init", "-y").Run()

	if err != nil {
		color.Red("Error making npm project: %s", err)
		return
	}

	color.Green("Made npm project: %s", name)

}
