package generator

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/mrbns/assetLoader/helper"
	conf "github.com/mrbns/assetLoader/internal/config"
)

var (
	config         = conf.GetConfig()
	assets         = []string{} // Contain Only File names;
	projectRoot, _ = os.Getwd()
)

func isValidFiles(fileName string) bool {
	ext := filepath.Ext(fileName)

	switch ext {
	case ".png", ".jpg", ".jpeg", ".webp", ".avif", ".gif", ".mp4", ".svg":
		return true
	default:
		return false
	}
}

func getFiles() error {
	if config.AssetDir == "" {
		return errors.New("please provide asset directory eg. --dir=/src/assets")
	}

	fileDir := path.Join(projectRoot, config.AssetDir)
	entries, err := os.ReadDir(fileDir)
	if err != nil || len(entries) <= 0 {
		fmt.Print(color.RedString("❌ %s Directory Doesn't Exist. \n Please Provide a Valid Directory name", fileDir))
		os.Exit(1)
	}

	for _, ent := range entries {
		// Code to handle each entry goes here
		if isValidFiles(ent.Name()) {
			assets = append(assets, ent.Name())
		}
	}

	return nil
}

// func restorePrevious() ([]string, error) {
// 	outputFile := path.Join(projectRoot, config.AssetDir, config.OutputFile)
// 	oldFiles := make([]string, 0)

// 	file, err := os.Open(filepath.Join(outputFile))
// 	helper.FatalIfError(err)
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 	}
// 	helper.FatalIfError(scanner.Err())

// 	return make([]string, 0), nil
// }

func GenerateAsset(directoryPath string) {
	getFiles()
	projectRoot, _ := os.Getwd()
	outputImgFile := filepath.Join(projectRoot, fmt.Sprintf("%v/%v", config.AssetDir, config.OutputFile))

	// Writing to file
	file, err := os.Create(outputImgFile)

	helper.FatalIfError(err)
	defer file.Close()

	if config.AssetDir == "" {
		return
	}
	assetDir := path.Join(projectRoot, config.AssetDir)

	if len(assets) < 1 {
		fmt.Println(color.RedString("❌ No Files Found in Folder - %s", assetDir))
		file.WriteString("")
		os.Exit(1)
		return
	}

	for index, name := range assets {
		ext := filepath.Ext(name)
		fileNameWithoutExtension := name[:len(name)-len(ext)]

		//Normalizing Name for JS variable
		normalizeNameRegex := regexp.MustCompile(`[-. !)(]`)
		normalizedFileName := normalizeNameRegex.ReplaceAllString(fileNameWithoutExtension, "_")

		if isValidFiles(name) {
			_, err := file.WriteString(fmt.Sprintf("export { default as %v_%v } from \"./%v\";\n", config.AssetPrefix, strings.ToUpper(normalizedFileName), name))
			fmt.Println(color.YellowString("✅ %d -  %s", index, name))
			helper.FatalIfError(err)
		}
	}

	fmt.Println(color.GreenString("🔥 added %s file at %s", config.OutputFile, outputImgFile))
	fmt.Println(color.BlueString("👋 Bye. see you again"))
}
