package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}
	defer file.Close()

	fs := NewFileSystem()
	processShellData(file, fs)

	sum := 0
	freeSpace := fs.FreeSpace()
	spaceRequired := 30000000 - freeSpace
	var bigEnough []*treeNode

	fs.Walk(func(node *treeNode) {
		sz := node.Size()
		if sz <= 100000 {
			sum += node.Size()
		}
		if sz > spaceRequired {
			bigEnough = append(bigEnough, node)
		}
	})

	sort.SliceStable(bigEnough, func(i, j int) bool {
		return bigEnough[i].Size() < bigEnough[j].Size()
	})
	maxSz := bigEnough[0].Size()
	fmt.Printf("Part 1: size = %d\n", sum)
	fmt.Printf("Part 2: size = %d\n", maxSz)
}

func processShellData(file io.Reader, fs *fileSystem) {
	scanner := bufio.NewScanner(file)
	output := ""
	var command string
	var currentDir string
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(fmt.Sprintf("error reading: %s", scanner.Err()))
		}
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			// process existing command if any
			if command != "" {
				switch command {
				case "ls":
					processListFilesOutput(fs, currentDir, strings.TrimSpace(output))
					break
				}
			}
			// command
			cmdAndArgs := strings.Split(line[2:], " ")
			command = cmdAndArgs[0]
			switch command {
			case "cd":
				if currentDir != "" && currentDir != "/" {
					currentDir = currentDir + "/" + cmdAndArgs[1]
					currentDir = processDots(currentDir)
				} else {
					currentDir = cmdAndArgs[1]
				}
				break
			case "ls":
				break
			}
			output = ""
		} else {
			switch command {
			case "ls":
				output = fmt.Sprintf("%s\n%s", output, line)
				break
			}
		}
	}
	if command == "ls" {
		processListFilesOutput(fs, currentDir, strings.TrimSpace(output))
	}
}

func processDots(dir string) string {
	// /test/foo/bar/.. -> /test/foo/bar

	elements := strings.Split(dir, "/")
	for i, component := range elements {
		if component == ".." {
			return strings.Join(elements[0:i-1], "/")
		}
	}
	return dir
}

func processListFilesOutput(fs *fileSystem, dir string, output string) {
	scanner := bufio.NewScanner(strings.NewReader(output))

	var currentNode *treeNode
	if dir == "/" {
		currentNode = fs.Root
	} else {
		currentNode = fs.GetPathNode(strings.Split(dir, "/"))
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "dir ") {
			dirName := line[4:]
			_, err := currentNode.Mkdir(dirName)
			if err != nil {
				panic(fmt.Sprintf("failed to mkdir: %s", err))
			}
		} else {
			elements := strings.Split(line, " ")
			sz, err := strconv.Atoi(elements[0])
			if err != nil {
				panic(fmt.Sprintf("failed to parse size: %s", elements[0]))
			}
			filename := elements[1]

			currentNode.AddFile(&file{
				Name: filename,
				Size: sz,
			})
		}
	}
}
