package gen

import (
	"bufio"
	"fmt"
	"ogcat/src/enums"
	"os"
	"strings"
)

const CHECKED_TXT = "checked.txt"

func AddUsernameToTxt(username *string, result enums.UsernameCheckReturnType) error {
	file, err := os.OpenFile(CHECKED_TXT, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintf(writer, "%s    %s\n", *username, result.String())
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// filters out usernames that have already been checked according to checked.txt.
// iterates through every line in and checks if the username is in the line.
func FilterUsernamesFrom(target *[]string) error {
	file, err := os.Open(CHECKED_TXT)
	if err != nil {
		if os.IsNotExist(err) {
			// If checked.txt does not exist, no filtering is needed
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, username := range *target {
			actual := strings.Split(line, "    ")
			if len(actual) != 2 {
				continue
			}

			if actual[0] == username {
				*target = append((*target)[:i], (*target)[i+1:]...)
			}
		}
	}

	return nil
}
