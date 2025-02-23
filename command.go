package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func execCommand() error {
	command := getArg(0, "Command is required")

	switch command {

	case "info":
		path := getArg(1, "Enter path")
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s\nSize: %d\nMode: %s,\nIsDir: %t\nModification: %s\n",
			fileInfo.Name(),
			fileInfo.Size(),
			fileInfo.Mode(),
			fileInfo.IsDir(),
			fileInfo.ModTime(),
		)

	case "ls":
		files, err := os.ReadDir(".")
		if err != nil {
			return err
		}
		for _, file := range files {
			fmt.Printf("%s  ", file.Name())
		}
		fmt.Println()

	case "rm":
		path := getArg(1, "Enter path")
		return os.RemoveAll(path)

	case "pwd":
		currentPath, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(currentPath)

	case "mv":
		path1 := getArg(1, "Enter path 1")
		path2 := getArg(2, "Enter path 2")

		return os.Rename(path1, path2)

	case "mkdir":
		path := getArg(1, "Enter Dir")
		return os.MkdirAll(path, 0777)

	case "cat":
		path := getArg(1, "Enter path")
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		os.Stdout.Write(data)

	case "touch":
		path := getArg(1, "Enter path")

		_, err := os.Create(path)
		return err

	case "cp": // copy file
		// path old file
		path1 := getArg(1, "Enter path 1")
		// path new file
		path2 := getArg(2, "Enter path 2")

		data, err := os.ReadFile(path1)
		if err != nil {
			return err
		}
		return os.WriteFile(path2, data, 0777)

	case "echo": // add content to file
		content := getArg(1, "Enter content")
		path := getArg(2, "Enter path")

		return os.WriteFile(path, []byte(content), 0777)
	default:
		fmt.Print(helpMess)
	}
	return nil

}

func getArg(indexArg int, message string) string {
	val := flag.Arg(indexArg)
	if val == "" {
		log.Fatal(message)
	}
	return val
}
