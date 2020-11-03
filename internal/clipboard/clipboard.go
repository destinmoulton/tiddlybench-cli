package clipboard

import (
	"github.com/atotto/clipboard"
	"tiddlybench-cli/internal/logger"
)

// Paste gets the contents of the clipboard
func Paste(log logger.Logger) string {
	text, err := clipboard.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return text
}
