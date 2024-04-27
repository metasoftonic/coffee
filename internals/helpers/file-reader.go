package helper

import "os"

func ReadHTMLFile(filename string) (string, error) {
	// Read the content of the HTML file
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert content to a string
	htmlContent := string(content)

	return htmlContent, nil
}
