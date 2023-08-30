package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func incrementMajorVersion(major int) string {
	// major, _ := strconv.Atoi(versionSegments[0])
	return fmt.Sprintf("%d.0.0", major+1)
}

func incrementMinorVersion(major int, minor int) string {
	// major, _ := strconv.Atoi(versionSegments[0])
	// minor, _ := strconv.Atoi(versionSegments[1])
	return fmt.Sprintf("%d.%d.0", major, minor+1)
}

func incrementPatchVersion(major int, minor int, patch int) string {
	// patch, _ := strconv.Atoi(versionSegments[2])
	return fmt.Sprintf("%d.%d.%d", major, minor, patch+1)
}

func getVersionDifferenceTypeAndIncrement(version, current, updated string) string {
	currentSegments := strings.Split(current, ".")
	updatedSegments := strings.Split(updated, ".")
	versionSegments := strings.Split(version, ".")

	major, _ := strconv.Atoi(versionSegments[0])
	minor, _ := strconv.Atoi(versionSegments[1])
	patch, _ := strconv.Atoi(versionSegments[2])

	if currentSegments[0] != updatedSegments[0] {
		return incrementMajorVersion(major)
	} else if currentSegments[1] != updatedSegments[1] {
		return incrementMinorVersion(major, minor)
	} else if currentSegments[2] != updatedSegments[2] {
		return incrementPatchVersion(major, minor, patch)
	} else {
		return "error"
	}
}

func main() {
	args := os.Args
	fmt.Println(args)
	if len(args) != 4 {
		fmt.Println("Usage: program version current updated")
		os.Exit(1)
	}

	version := args[1]
	current := args[2]
	updated := args[3]

	result := getVersionDifferenceTypeAndIncrement(version, current, updated)
	fmt.Println(result)
	// fmt.Printf("::set-output name=result::%s\n", result)
	// fmt.Printf(`"{result}={result}" >> $GITHUB_OUTPUT`)
	fmt.Printf(`"%s=%s" >> $GITHUB_OUTPUT`, "result", result)
}
