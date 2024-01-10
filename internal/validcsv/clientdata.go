package validcsv

import "errors"

var clientFiles = map[string]ClientFile{}

func SaveClientFile(hash string, file *ClientFile) error {
	if _, ok := clientFiles[hash]; ok {
		return errors.New("file already exists")
	}
	clientFiles[hash] = *file
	return nil
}

func GetClientFile(hash string) (ClientFile, error) {
	if _, ok := clientFiles[hash]; !ok {
		return ClientFile{}, errors.New("file not found")
	}
	return clientFiles[hash], nil
}

func GetClientFileRows(hash string) (int, error) {
	cf, err := GetClientFile(hash)
	if err != nil {
		return 0, err
	}
	return cf.GetNumberOfRows(), nil
}

func GetValidatedRow(hash string, row int) ([]string, error) {
	cf, err := GetClientFile(hash)
	if err != nil {
		return []string{}, err
	}
	return *cf.ValidateRow(row), nil
}
