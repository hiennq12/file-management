package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Manager struct {
	logger Logger
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func New(logger Logger) *Manager {
	return &Manager{
		logger: logger,
	}
}

func (m *Manager) Copy(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		m.logger.Error("Failed to open source file:", err)
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		m.logger.Error("Failed to create destination file:", err)
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		m.logger.Error("Failed to copy file:", err)
		return err
	}

	m.logger.Info("Successfully copied file from", src, "to", dst)
	return nil
}
func (m *Manager) List(path string, detailed bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		m.logger.Error("Failed to read directory:", err)
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, f := range files {
		if detailed {
			m.logger.Info(fmt.Sprintf("%s  %8d  %s  %s",
				f.Mode(),
				f.Size(),
				f.ModTime().Format("2006-01-02 15:04:05"),
				f.Name()))
		} else {
			m.logger.Info(f.Name())
		}
	}

	return nil
}
