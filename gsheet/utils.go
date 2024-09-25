package gsheet

import "google.golang.org/api/sheets/v4"

func BatchUpdate(spreadsheetID string, requests ...*sheets.Request) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	service, err := GetService()
	if err != nil {
		return nil, err
	}
	if len(requests) == 0 {
		return nil, nil
	}

	return service.Spreadsheets.BatchUpdate(spreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}).Do()

}
