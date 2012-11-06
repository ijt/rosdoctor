package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Package struct {
	name string
	dir string
}

func main() {
	if checkLocalPackages() == 0 {
		fmt.Println("No issues found.")
	}
}

// checkLocalPackages warns if there are local (pip or easy_install) versions of
// Python packages and returns how many were found.
func checkLocalPackages() int {
	localPkgs := findLocalPackages()
	if len(localPkgs) > 0 {
		for _, p := range(localPkgs) {
			warn("Found a local version of", p.name, "in", p.dir)
		}
		names := []string {}
		for _, p := range(localPkgs) {
			names = append(names, p.name)
		}
		cmd := "sudo pip uninstall " + strings.Join(names, " ")
		fmt.Println("You may want to run ", cmd)
		fmt.Println(`It is a good idea to keep running this command
	until it stops doing anything.`)
	}
	return len(localPkgs)
}

func warn(args ...string) {
	fmt.Println("Warning: ", strings.Join(args, " "))
}

// checkForLocalPackages looks for local packages that shadow ones that should
// have been installed using apt or another package manager.
func findLocalPackages() (foundPkgs []Package) {
	dir := "/usr/local/lib/python2.7/dist-packages/"
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(dir, " does not exist so not checking it.")
		return
	}
	pkgs := []string { "rosinstall", "rosdep", "rospkg", "vcstools", "catkin", "bloom" }
	for _, name := range(pkgs) {
		matches, _ := filepath.Glob(name + "-*")
		if len(matches) > 0 {
			pkg := Package{ name, dir }
			foundPkgs = append(foundPkgs, pkg)
		}
	}
	return
}

