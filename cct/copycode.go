// copycode.go
package cct

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var excludedFiles map[string]struct{}

func findExcludeFile(rootDir string) (string, error) {
	// Используем рекурсивный обход всех файлов и папок внутри корневой директории
	var excludeFilePath string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "exclude.txt" {
			excludeFilePath = path
			return fmt.Errorf("файл exclude.txt найден") // прерываем поиск
		}
		return nil
	})
	if err != nil && err.Error() != "файл exclude.txt найден" {
		return "", err
	}
	if excludeFilePath == "" {
		return "", fmt.Errorf("файл exclude.txt не найден")
	}
	return excludeFilePath, nil
}

func initExcludedFilesMap() error {
	excludedFiles = make(map[string]struct{})

	// Находим корневую директорию проекта
	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}

	// Находим файл exclude.txt в проекте
	excludeFilePath, err := findExcludeFile(projectRoot)
	if err != nil {
		return err
	}

	// Открываем файл exclude.txt для чтения
	file, err := os.Open(excludeFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Читаем файл построчно и добавляем каждую строку в map исключений
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		excludedFiles[scanner.Text()] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func isFileExcluded(filename string) bool {
	// Проверяем, если filename в map исключенных файлов
	_, excluded := excludedFiles[filename]
	return excluded
}

func CopyCode() {
	// Вызываем инициализацию исключенных файлов из пакета cct
	if err := initExcludedFilesMap(); err != nil {
		fmt.Println("Ошибка инициализации исключенных файлов:", err)
		return
	}

	// Проверка, есть ли файлы для исключения
	if len(excludedFiles) == 0 {
		fmt.Println("Ничего не добавлено в файл exclude.txt.")
		return
	}

	// Создаем или перезаписываем файл output.txt в корне проекта
	outputFilePath := "partcode.txt"
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Ошибка создания выходного файла:", err)
		return
	}
	defer outputFile.Close()

	// Перебираем все файлы в проекте и записываем их содержимое в output.txt, исключая те, которые указаны в exclude.txt
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Исключаем только файлы, а не директории
		if !info.IsDir() && isFileExcluded(info.Name()) {
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Ошибка чтения файла %s: %v\n", path, err)
				return nil // продолжаем обход
			}
			outputFile.WriteString(fmt.Sprintf("Файл: %s\n", path))
			outputFile.Write(content)
			outputFile.WriteString("\n\n")
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Ошибка при копировании файлов: %v\n", err)
		return
	}

	fmt.Println("Программа завершила выполнение, результаты записаны в", outputFilePath)
}
