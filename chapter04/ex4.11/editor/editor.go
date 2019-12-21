package editor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

func Edit(value map[string]string) error {
	editor := getEditorName()

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}

	tempFileName := tempFile.Name()
	defer os.Remove(tempFileName)

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "   ")

	if err = encoder.Encode(value); err != nil {
		panic(err)
	}

	if err = tempFile.Close(); err != nil {
		panic(err)
	}

	cmd := exec.Command(editor, tempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		panic(err)
	}
	edited, err := ioutil.ReadFile(tempFileName)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(removeUTF8BOM(edited), &value); err != nil {
		panic(err)
	}
	return nil

}

// https://git-scm.com/book/en/v2/Git-Internals-Environment-Variables
func getEditorName() string {
	editor := os.Getenv("GIT_EDITOR")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}

	if editor == "" {
		editor = "emacs"
	}
	return editor
}

// removeUTF8BOM
func removeUTF8BOM(s []byte) []byte {
	utf8Bom := []byte{239, 187, 191}
	return bytes.TrimPrefix(s, utf8Bom)
}
