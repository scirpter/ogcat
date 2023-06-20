package tests

import (
	"fmt"
	"io"
	"ogcat/src/common"
	"os"
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	file, err := os.Open("../../accounts.txt")
	if err != nil {
		t.Error("Failed to open file.")
	}
	defer file.Close()

	body, _ := io.ReadAll(file)
	str := string(body)
	authData := fmt.Sprintf(str, common.VERSION)
	encrypted, _ := common.Encrypt(&authData)
	fmt.Println(strings.ReplaceAll(*encrypted, " ", ""))
}
