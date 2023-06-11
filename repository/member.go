package repository

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"

	"github.com/ramrodo/facebook-members/model"
)

func GetFilePaths() []string {
	paths := []string{}
	err := filepath.Walk("/home/ramrodo/git/liberarte-creaciones/facebook-members/data", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		paths = append(paths, path)

		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return paths
}

// GetAll - reads CSV file and returns an array of all the Members in CSV files
func GetAll() ([]model.Member, error) {
	allMembers := []model.Member{}

	files := GetFilePaths()
	for _, file := range files {
		csvFile, err := os.Open(file)
		if err != nil {
			log.Fatalln(err)
		}
		defer csvFile.Close()

		reader := csv.NewReader(bufio.NewReader(csvFile))

		data, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		for i, row := range data {
			if i == 0 {
				continue
			}

			allMembers = append(allMembers, model.Member{
				ProfileId:        row[0],
				Fullname:         row[1],
				ProfileLink:      row[2],
				Bio:              row[3],
				ImageSrc:         row[4],
				GroupId:          row[5],
				GroupJoiningText: row[6],
				ProfileType:      row[7],
			})
		}
	}

	return allMembers, nil
}
