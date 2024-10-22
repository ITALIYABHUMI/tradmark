package cmd

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"os"
	"path/filepath"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/tradmark/config"
	"github.com/tradmark/model"
)

func Unzip(file string) (result string) {

	if strings.HasSuffix(file, ".zip") {
		dst := "output"
		var filePath string
		archive, err := zip.OpenReader(file)
		if err != nil {
			log.Fatal("no such file or directory")
			return "no such file or directory"
		}
		defer archive.Close()

		for _, f := range archive.File {
			filePath = filepath.Join(dst, f.Name)

			log.Printf("unzipping %s file in %s\n", file, filePath)

			if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
				log.Fatal("invalid file path")
				return
			}
			if f.FileInfo().IsDir() {
				log.Print("creating directory...")
				os.MkdirAll(filePath, os.ModePerm)
				continue
			}

			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				log.Fatal(err)
				return
			}

			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.Fatal(err)
				return
			}

			fileInArchive, err := f.Open()
			if err != nil {
				log.Fatal(err)
				return
			}

			if _, err := io.Copy(dstFile, fileInArchive); err != nil {
				log.Fatal(err)
				return
			}

			dstFile.Close()
			fileInArchive.Close()
		}

		log.Print("Unzipped successfully")

		jsonData, err := ConverToJsonData(filePath)
		if err != nil {
			return fmt.Sprintf("failed to covert data to Json: %v", err)
		}

		if err = SaveDataToDb(jsonData); err != nil {
			return fmt.Sprintf("failed to save data to database: %v", err)
		}
	}

	if strings.HasSuffix(file, ".xml") {
		log.Println("processing XML data and converting it to JSON...")

		jsonData, err := ConverToJsonData(file)
		if err != nil {
			return fmt.Sprintf("failed to covert data to Json: %v", err)
		}

		log.Println("unmarshalled JSON")

		if err = SaveDataToDb(jsonData); err != nil {
			return fmt.Sprintf("failed to save data to database: %v", err)
		}
	}

	return "XML data inserted successfully"
}

func ConverToJsonData(file string) (model.TrademarkApplicationsDailyWrapper, error) {
	xmlData, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("failed to read XML file:", err)
		return model.TrademarkApplicationsDailyWrapper{}, fmt.Errorf("failed to read XML file")
	}

	// Convert XML to JSON
	xmlReader := strings.NewReader(string(xmlData))
	var jsonData *bytes.Buffer
	jsonData, err = xj.Convert(xmlReader)
	if err != nil {
		log.Fatal("failed to convert XML to JSON:", err)
		return model.TrademarkApplicationsDailyWrapper{}, fmt.Errorf("failed to convert XML to JSON")
	}

	// Unmarshal JSON into the root structure
	var trademarkApplicationsDailyWrapper model.TrademarkApplicationsDailyWrapper
	err = json.Unmarshal(jsonData.Bytes(), &trademarkApplicationsDailyWrapper)
	if err != nil {
		log.Printf("failed to unmarshal JSON: %v", err)
		return model.TrademarkApplicationsDailyWrapper{}, fmt.Errorf("failed to unmarshal JSON")
	}
	return trademarkApplicationsDailyWrapper, nil
}

func SaveDataToDb(trademarkApplicationsDailyWrapper model.TrademarkApplicationsDailyWrapper) error {
	db := config.GetDB()
	if db == nil {
		return fmt.Errorf("Database is not initialized")
	}
	log.Println("Database is initialized")
	for _, caseFile := range trademarkApplicationsDailyWrapper.TrademarkApplicationsDaily.ApplicationInformation.FileSegments.ActionKeys[0].CaseFile {
		if err := db.Table("case_files").Create(&caseFile).Error; err != nil {
			log.Printf("failed to insert CaseFile: %v", err)
			return err
		}
	}
	log.Println("All CaseFiles inserted successfully")
	return nil
}
