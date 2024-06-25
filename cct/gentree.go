package cct

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Функция для проверки, является ли имя файла или папки игнорируемым
func isIgnored(name string) bool {
	ignoredNames := []string{".idea", ".git", "cct", "project_structure.md", "output.txt"}
	for _, ignored := range ignoredNames {
		if name == ignored {
			return true
		}
	}
	return false
}

func printDirTree(path string, prefix string, file *os.File) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for i, fileInfo := range files {
		if isIgnored(fileInfo.Name()) { // Проверяем, является ли файл или папка игнорируемым
			continue
		}

		line := ""
		if i == len(files)-1 {
			line = fmt.Sprintf("%s└── %s\n", prefix, fileInfo.Name())
		} else {
			line = fmt.Sprintf("%s├── %s\n", prefix, fileInfo.Name())
		}

		if _, err := file.WriteString(line); err != nil {
			log.Fatal(err)
		}

		if fileInfo.IsDir() {
			newPrefix := prefix + "│   "
			if i == len(files)-1 {
				newPrefix = prefix + "    "
			}
			printDirTree(filepath.Join(path, fileInfo.Name()), newPrefix, file)
		}
	}
}

func generateProjectStructure(root string) {
	// Открытие файла для записи
	file, err := os.Create("project_structure.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Запись заголовка Markdown
	if _, err := file.WriteString("# Project Structure\n\n"); err != nil {
		log.Fatal(err)
	}

	printDirTree(root, "", file)
}

func GenTree() {
	root := "."
	generateProjectStructure(root)
}
