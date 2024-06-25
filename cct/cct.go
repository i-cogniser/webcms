package cct

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isExcludedDir(dirname string) bool {
	excludedDirs := []string{".idea", ".git", "node_modules", "administration", "cct"}
	for _, dir := range excludedDirs {
		if dirname == dir {
			return true
		}
	}
	return false
}

func isExcludedFile(filename, outputFile string) bool {
	excludedFiles := []string{"go.mod", "go.sum", "text.txt", "copycode.go", ".gitignore", ".gitattributes", "project_structure.md", filepath.Base(outputFile)}
	for _, file := range excludedFiles {
		if filename == file {
			return true
		}
	}
	return false
}

func CopyProjectToTextFile(sourceDir, outputFile string) error {
	// Создание/открытие файла для записи
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	// Функция для обработки каждого файла
	var processFile = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Пропускать директории, которые нужно исключить
		if info.IsDir() && isExcludedDir(info.Name()) {
			return filepath.SkipDir
		}

		// Пропускать файлы, которые нужно исключить
		if !info.IsDir() && isExcludedFile(info.Name(), outputFile) {
			return nil
		}

		// Проверяем, содержит ли путь к файлу строку "templates/favicon.ico" или "templates\favicon.ico"
		if strings.Contains(path, "templates/favicon.ico") || strings.Contains(path, "templates\\favicon.ico") {
			return nil
		}

		// Если это файл, читаем и записываем его содержимое
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Запись информации о файле и его содержимом в выходной файл
			relativePath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				return err
			}

			header := fmt.Sprintf("Файл: %s\n\n", relativePath)
			if _, err := output.WriteString(header); err != nil {
				return err
			}

			if _, err := output.Write(content); err != nil {
				return err
			}

			if _, err := output.WriteString("\n\n"); err != nil {
				return err
			}
		}

		return nil
	}

	// Рекурсивное прохождение по директории
	if err := filepath.Walk(sourceDir, processFile); err != nil {
		return err
	}

	return nil
}
func CcT() {

	sourceDir := "."
	outputFile := "codeall.txt"

	if err := CopyProjectToTextFile(sourceDir, outputFile); err != nil {
		fmt.Println("Ошибка cct → завершено не успешно :(", err)
		return
	}

	fmt.Println("cct → :)")
}
