package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mo-zza/steader_words_write_module/domain/entity"
	"github.com/mo-zza/steader_words_write_module/domain/interactor"
	"github.com/xuri/excelize/v2"
)

type Word struct {
	words entity.Words
}

func Save() error {
	file, err := excelize.OpenFile(os.Getenv("FILE_PATH")+os.Getenv("FILE_NAME")+".xlsx")
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows, err := file.GetRows(os.Getenv("SHEET_NAME"))
	if err != nil {
		fmt.Println(err)
	}
	for index, _ := range rows {
		strIndex := strconv.Itoa(index+1)
		searchBCell := "B"+strIndex
		bCell, err := file.GetCellValue(os.Getenv("SHEET_NAME"), searchBCell)
		if err != nil {
			fmt.Println(err)
			return err
		}
		searchACell := "A"+strIndex
		aCell, err := file.GetCellValue(os.Getenv("SHEET_NAME"), searchACell)
		if err != nil {
			fmt.Println(err)
			return err
		}

		var w Word
		w.words.Key = aCell
		w.words.Value = bCell
		interactor.WordsInteractor(index, w.words)
	}
	return err
}