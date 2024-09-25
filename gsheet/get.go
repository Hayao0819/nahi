package gsheet

import "google.golang.org/api/sheets/v4"

func GetSheetValues(id string, readRange string) (*sheets.ValueRange, error) {

	srv, err := GetService()
	if err != nil {
		return nil, err
	}

	return srv.Spreadsheets.Values.Get(id, readRange).Do()
}

func GetSheet(id string) (*sheets.Spreadsheet, error) {
	srv, err := GetService()
	if err != nil {
		return nil, err
	}

	return srv.Spreadsheets.Get(id).Do()
}
