package gsheet

import (
	"log"
	"log/slog"

	"github.com/cockroachdb/errors"
	"google.golang.org/api/sheets/v4"
)

func CopySheet1(id string, source string, dest string) (*sheets.SheetProperties, error) {
	srv, err := GetService()
	if err != nil {
		return nil, err
	}

	sheet, err := srv.Spreadsheets.Get(id).Do()
	if err != nil {
		return nil, err
	}
	slog.Info("sheet", "title", sheet.Properties.Title)

	sourceId := FindSheetId(sheet, source)
	if sourceId == -1 {
		return nil, errors.Newf("sheet not found: %s", source)
	}

	// Sheet1の範囲をコピー
	sourceValues, err := srv.Spreadsheets.Values.Get(id, source).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Sheet2を作成
	if _, err := CreateNewSheet(id, dest); err != nil {
		return nil, err
	}

	// Sheet2にデータを書き込む
	_, err = srv.Spreadsheets.Values.Update(id, dest, &sheets.ValueRange{
		Values: sourceValues.Values,
	}).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to write data to sheet: %v", err)
	}
	return nil, nil
}

func CopySheet2(id string, source string, dest string) (*sheets.SheetProperties, error) {
	srv, err := GetService()
	if err != nil {
		return nil, err
	}

	sheet, err := srv.Spreadsheets.Get(id).Do()
	if err != nil {
		return nil, err
	}
	slog.Info("sheet", "title", sheet.Properties.Title)

	sourceId := FindSheetId(sheet, source)
	if sourceId == -1 {
		return nil, errors.Newf("sheet not found: %s", source)
	}

	// var createdCopySheetName string
	var destSheetId int64
	if res, err := srv.Spreadsheets.Sheets.CopyTo(id, sourceId, &sheets.CopySheetToAnotherSpreadsheetRequest{
		DestinationSpreadsheetId: id,
	}).Do(); err != nil {
		slog.Error("failed to copy", "title", source)
		return nil, err
	} else {
		slog.Info("copied", "id", res.SheetId)
		destSheetId = res.SheetId
	}

	if _, err := RenameSheetById(id, destSheetId, dest); err != nil {
		return nil, err
	}
	return nil, nil
}
