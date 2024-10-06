package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"practice/internal/config"
)

var (
	MainConfig *config.Config
)

func init() {
	projectDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("***************** CONFIG RUN *****************")
	MainConfig = config.MustLoad(filepath.Join(projectDir, "config", "config.yml"))
}
