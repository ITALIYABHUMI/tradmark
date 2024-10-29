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
	"github.com/tradmark/api/model"
	"github.com/tradmark/config"
	"github.com/tradmark/pkg"
)

func Unzip(file string) (string, error) {

	dst := "output"
	var filePath string
	archive, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal("no such file or directory")
		return "", err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath = filepath.Join(dst, f.Name)

		log.Printf("unzipping %s file in %s\n", file, filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return "", fmt.Errorf("invalid file path")

		}
		if f.FileInfo().IsDir() {
			log.Print("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			log.Fatal(err)
			return "", err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Fatal(err)
			return "", err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			log.Fatal(err)
			return "", err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			log.Fatal(err)
			return "", err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return "Unzipped successfully", nil
}

func ProcessFile(file string) error {
	log.Println("processing XML data and converting it to JSON...")

	jsonData, err := ConvertToJSON(file)
	if err != nil {
		return err
	}

	log.Println("unmarshalled JSON")

	if err = SaveDataToDB(jsonData); err != nil {
		return err
	}
	return nil
}

func ConvertToJSON(file string) (model.TrademarkApplicationsDailyWrapper, error) {
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

	// Unmarshal JSON into model.TrademarkApplicationsDailyWrapper
	var trademarkApplicationsDailyWrapper model.TrademarkApplicationsDailyWrapper
	err = json.Unmarshal(jsonData.Bytes(), &trademarkApplicationsDailyWrapper)
	if err != nil {
		log.Printf("failed to unmarshal JSON: %v", err)
		return model.TrademarkApplicationsDailyWrapper{}, fmt.Errorf("failed to unmarshal JSON")
	}
	return trademarkApplicationsDailyWrapper, nil
}

func SaveDataToDB(trademarkApplicationsDailyWrapper model.TrademarkApplicationsDailyWrapper) error {

	db := config.GetDB()
	if db == nil {
		return fmt.Errorf("Database is not initialized")
	}
	log.Println("Database is initialized")
	for _, caseFile := range trademarkApplicationsDailyWrapper.TrademarkApplicationsDaily.ApplicationInformation.FileSegments.ActionKeys[0].CaseFile {
		if err := pkg.TradesRepository.CreateCaseFiles(config.GetDB(), &caseFile); err != nil {
			return fmt.Errorf("CaseFiles not added into database")
		}
	}

	log.Println("All CaseFiles inserted successfully")
	return nil
}
