package editor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Edit opens the desired text in the desired editor with the arguments
func Edit(text string, editorSetting string, args string) (string, error) {

	editor := getPreferredEditor(editorSetting)

	textBytes, err := captureInputFromEditor(text, editor, args)
	if err != nil {
		return "", err
	}
	return string(textBytes), nil
}

// getPreferredEditor returns the user's editor as defined by the
// `$EDITOR` environment variable, or the editorSetting
func getPreferredEditor(editorSetting string) string {
	editor := editorSetting
	if editorSetting == "$EDITOR" {
		// Get it from the $EDITOR env setting
		editor = os.Getenv("EDITOR")
	}

	return editor
}

// buildEditorArguments builds the command string
func buildEditorArguments(args string, filename string) string {
	if args != "" {
		return strings.TrimSpace(args) + " " + filename
	}
	return filename
}

// openFileInEditor opens filename in a text editor.
func openFileInEditor(filename string, editor string, args string) error {
	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	editorArguments := buildEditorArguments(args, filename)

	fmt.Println("openFileInEditor editor " + editorArguments)
	cmd := exec.Command(executable, editorArguments)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// captureInputFromEditor opens a temporary file in a text editor and returns
// the written bytes on success or an error on failure. It handles deletion
// of the temporary file behind the scenes.
func captureInputFromEditor(text string, editor string, args string) ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	// Defer removal of the temporary file in case any of the next steps fail.
	defer os.Remove(filename)

	if _, err := file.Write([]byte(text)); err != nil {
		return []byte{}, err
	}

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = openFileInEditor(filename, editor, args); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
