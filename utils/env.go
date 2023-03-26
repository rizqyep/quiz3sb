package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
)

var envMap map[string]string
var once sync.Once

func InitEnv() map[string]string {
	_, b, _, _ := runtime.Caller(0)
	root_dir := path.Join(path.Dir(b))

	env, envErr := godotenv.Read(os.ExpandEnv(fmt.Sprintf("%s/config/.env", filepath.Dir(root_dir))))
	if envErr != nil {
		fmt.Println(envErr)
		fmt.Println("Error reading env")
		os.Exit(1)
	}
	return env
}

func GetEnv() map[string]string {
	once.Do(func() {
		envMap = InitEnv()
	})
	return envMap
}
