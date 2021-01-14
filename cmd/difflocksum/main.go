package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dgnorton/gomod"
)

func usage() {
	fmt.Println("usage: difflocksum <path/Gopkg.lock> <path/go.sum>")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	depLockPath, goSumPath := os.Args[1], os.Args[2]

	diffs, err := gomod.DiffLockSum(depLockPath, goSumPath)
	check(err, "diff failed")

	if len(diffs) == 0 {
		fmt.Printf("%q and %q are equal\n", depLockPath, goSumPath)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Dependency\tLockVer\tModVer")
	for _, diff := range diffs {
		fmt.Fprintf(w, "%s\t%s\t%s\n", diff.ProjectName, diff.DepVer, diff.ModVer)
	}
	w.Flush()
}

func check(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		os.Exit(1)
	}
}
