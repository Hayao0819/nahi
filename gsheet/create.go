package gsheet

import "google.golang.org/api/sheets/v4"

func CreateNewSheet(id string, title string) (*sheets.AddSheetResponse, error) {
	res, err := BatchUpdate(id, &sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Title: title,
			},
		},
	})
	return res.Replies[0].AddSheet, err
}
