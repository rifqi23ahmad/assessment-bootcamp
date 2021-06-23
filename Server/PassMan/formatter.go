package password

type DeleteFormat struct {
	Data string `json:"data"`
}

func FormatDelete(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Data: msg,
	}
	return deleteFormat
}
