package gsheet

import "google.golang.org/api/sheets/v4"

func RenameSheetByName(id string, source string, target string) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	sheet, err := GetSheet(id)
	if err != nil {
		return nil, err
	}
	sourceSheetId := FindSheetId(sheet, source)

	// req := sheets.Request{
	// 	UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
	// 		Properties: &sheets.SheetProperties{
	// 			SheetId: sourceSheetId,
	// 			Title:   "Sheet2", // 新しいシート名
	// 		},
	// 		Fields: "title",
	// 	},
	// }

	// res, err := BatchUpdate(id, &req)

	return RenameSheetById(id, sourceSheetId, target)
}

func RenameSheetById(id string, source int64, target string) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	req := sheets.Request{
		UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
			Properties: &sheets.SheetProperties{
				SheetId: source,
				Title:   target, // 新しいシート名
			},
			Fields: "title",
		},
	}

	res, err := BatchUpdate(id, &req)
	return res, err
}
