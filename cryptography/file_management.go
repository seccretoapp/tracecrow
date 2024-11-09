package cryptography

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadEncryptedData(filename string) (EncryptedData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return EncryptedData{}, fmt.Errorf("erro ao abrir o arquivo %s: %w", filename, err)
	}
	defer file.Close()

	var ed EncryptedData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ed); err != nil {
		return EncryptedData{}, fmt.Errorf("erro ao carregar dados criptografados: %w", err)
	}
	return ed, nil
}

func SaveEncryptedData(filename string, ed EncryptedData) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo %s: %w", filename, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(ed); err != nil {
		return fmt.Errorf("erro ao salvar dados criptografados: %w", err)
	}
	return nil
}
