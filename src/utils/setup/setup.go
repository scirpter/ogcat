package setup

import (
	"ogcat/src/utils/gen"
	"os"
)

const CONFIG_FILE_NAME = "config.json"

type Config map[string]string

// returns either "exists" as bool or error only (1: config, 2: checked).
// includes config.json & checked.txt.
func CreateFilesIfNotExists() (bool, bool, error, error) {
	var didConfigExist bool
	var didCheckedExist bool
	var errConfig error
	var errChecked error

	if _, err := os.Stat(CONFIG_FILE_NAME); os.IsNotExist(err) {
		_, errConfig = os.Create(CONFIG_FILE_NAME)
	} else {
		didConfigExist = true
	}

	if _, err := os.Stat(gen.CHECKED_TXT); os.IsNotExist(err) {
		_, errChecked = os.Create(gen.CHECKED_TXT)
	} else {
		didCheckedExist = true
	}

	return didConfigExist, didCheckedExist, errConfig, errChecked
}
