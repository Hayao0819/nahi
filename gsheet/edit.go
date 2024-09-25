package gsheet

import (
	"log/slog"

	"google.golang.org/api/sheets/v4"
)

func AppendLine(id string, writeRange string, v [][]interface{}) (*sheets.AppendValuesResponse, error) {
	srv, err := GetService()
	if err != nil {
		return nil, err
	}
	return srv.Spreadsheets.Values.Append(id, writeRange, &sheets.ValueRange{
		Values: v,
	}).ValueInputOption("RAW").Do()
}

func FindSheetId(sheet *sheets.Spreadsheet, name string) int64 {
	for _, s := range sheet.Sheets {
		slog.Debug("sheet", "title", s.Properties.Title)
		if s.Properties.Title == name {
			// slog.Debug("found", "title", name, "index", i)

			return s.Properties.SheetId
		}
	}
	return -1
}
