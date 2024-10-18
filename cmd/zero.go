package cmd

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"

	// "log/slog/internal/buffer"
	"os"
	"path/filepath"
	"strings"

	// xj "github.com/basgys/goxml2json"
	"github.com/tradmark/config"
	"github.com/tradmark/model"
)

func Unzip(file string) (result string) {

	dst := "output"
	var filePath string
	archive, err := zip.OpenReader(file)
	if err != nil {
		log.Print("no such file or directory")
		return
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath = filepath.Join(dst, f.Name)

		log.Printf("unzipping %s file in %s\n", file, filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			log.Print("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			log.Print("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			log.Print(err)
			return
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Print(err)
			return
		}

		fileInArchive, err := f.Open()
		if err != nil {
			log.Print(err)
			return
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			log.Print(err)
			return
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	log.Print("Unzipped successfully")

	if strings.HasSuffix(filePath, ".xml") {
		log.Printf("inserting XML data into the database...")

		xmlData, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal("failed to read XML file:", err)
			return "failed to read XML file"
		}

		// xml := strings.NewReader(string(xmlData))

		// jsonData, err := xj.Convert(xml)
		// if err != nil {
		// 	log.Fatal("failed to covert XML file to json:", err)
		// 	return "failed to covert XML file to json"
		// }

		// var buf []byte
		caseFile := &model.CaseFile{}

		// buf, err = buffer.WriteTo(jsonData)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		if err := json.Unmarshal(xmlData, &caseFile); err != nil {
			log.Printf("failed to unmarshal XML: %v", err)
			return "failed to unmarshal XML"
		}

		result, err := xml.Marshal(caseFile)
		if err != nil {
			log.Printf("failed to marshal casefile: %v", err)
			return "failed to marshal casefile"
		}

		fmt.Println("%s\n", len(result))

		db := config.GetDB()
		if db == nil {
			return "Database is not initialized"
		}
		fmt.Println(caseFile.RegistrationNumber)
		if err := db.Table("case_files").Create(result).Error; err != nil {
			log.Printf("failed to insert XML data: %v", err)
			return "failed to insert XML data"
		}
		log.Printf("XML data inserted successfully")
	}
	return "XML data inserted successfully"
}
