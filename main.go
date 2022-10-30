package main

import (
	"GroupBySize/helpers"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	// result := helpers.GetCurrentDir()
	var groupUncategorized = true
	ungroupedFolderName := "Remained"
	destination := "/home/yosia/Desktop/Projects/GroupBySize/"
	var groupingNumbers = []float32{2, 4, 6}
	// sorting the grouping numbers in right order
	var groupingFolders = []string{}
	handleIndex := func(i int) float32 {
		if i == 0 {
			return 0
		} else {
			return groupingNumbers[i-1]
		}
	}
	handleUnmoved := func (errorMsg, path string, response bool)  {
		if !response {
			helpers.HandlesError(fmt.Errorf("%v %v",errorMsg, path))
		}
	}

	// creatinng the grouping of folder slice from the grouping Numbers
	for i, groupingNumber := range groupingNumbers {
		groupingFolders = append(groupingFolders, fmt.Sprintf("%v-%v", handleIndex(i), groupingNumber))

	}

	moveFile := func (folderName, fileName, workingDirectory string)  {
		moved := helpers.RenameFile(helpers.JoinPath(workingDirectory, fileName), helpers.JoinPath(destination, folderName, fileName))
		handleUnmoved("Failed to Move file", helpers.JoinPath(workingDirectory, fileName),  moved)
	
	}
	handleDir := func (file fs.DirEntry, folderName string)  {
		err := filepath.WalkDir(file.Name(), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {

				// we create that directory
				// destination
				notch := strings.Contains(path, folderName)
				if !notch {
					created := helpers.CreateDir(helpers.JoinPath(destination, folderName, path), helpers.RdWrAll)
					if !created {
						log.Fatal("Failed to create a")
					}
				}

			}
			if !d.IsDir() {
				fmt.Println(path, "......")
				moved := helpers.RenameFile(path, helpers.JoinPath(destination, folderName, path))
				handleUnmoved("Failed to move", path, moved)
				
			}
			return err
		})
		if err != nil {
			helpers.HandlesError(err)

		}
		
	}
	// appending the remaining
	if groupUncategorized {
		groupingFolders = append(groupingFolders, ungroupedFolderName)

	}
	folderSelector := func(size int64) (int, string) {
		sizeMbs := helpers.SizeConverter(size, "mbs")
		var index int
		var folder string
		for i, groupingNumber := range groupingNumbers {
			if sizeMbs <= groupingNumber {
				index = i
				folder = groupingFolders[i]
				break
			}
		}
		return index, folder
	}
	for _, folder := range groupingFolders {
		createdFolder := helpers.CreateDir(folder, helpers.RdWrAll)
		if !createdFolder {
			return
		}
	}
	// // fmt.Printf("Data passed %v \n", result)
	helpers.Chdir("./Play")
	workingDirectory := helpers.GetCurrentDir()
	dirCreated := helpers.GetDirs(workingDirectory)
	for _, file := range dirCreated {
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("FileName: %v =>Size: %v \n", info.Name(), info.Size())
			_, folderName := folderSelector(info.Size())
			if folderName != "" {
				moveFile(folderName, file.Name(), workingDirectory)
			}
			if groupUncategorized && (folderName == "") {
				moveFile(ungroupedFolderName, file.Name(), workingDirectory)
			}
			// fmt.Println(folderSelector(8048576))
		} else {

			dirPath := helpers.JoinPath(workingDirectory, file.Name())
			size, _ := helpers.DirSizes(dirPath)
			_, folderName := folderSelector(size)
			if folderName != "" {
				handleDir(file, folderName)
			}
			if groupUncategorized && (folderName == "") {
				handleDir(file, ungroupedFolderName)
			}
		}

	}

}
