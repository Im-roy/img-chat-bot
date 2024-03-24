package filerepo

import (
	"context"
	"fmt"
	"img-chat-bot/model"
	"io"
	"os"
	"path/filepath"
)

type FileRepo struct {
	directory string
}

func (f *FileRepo) SetDirectory(directoryPath string) error {
	if err := os.MkdirAll(directoryPath, 0777); err != nil {
		return fmt.Errorf("directory path: {%v} doesn't exist and failed to create", directoryPath)
	}
	f.directory = directoryPath
	return nil
}

func (f *FileRepo) SaveFile(ctx context.Context, file model.FileDetailsModel) error {
	// Create a file with a unique name in the uploads directory
	filePath := filepath.Join(f.directory, file.Header.Filename)
	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Copy the uploaded file data to the newly created file
	_, err = io.Copy(newFile, file.Data)
	if err != nil {
		return err
	}
	return nil
}
