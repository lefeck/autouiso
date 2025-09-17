package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Constants
const (
	InsertParams = "autoinstall ds=nocloud;s=/cdrom/" // Autoinstall parameter string
	FilePerm     = 0644                               // Default file permission (rw-r--r--)
)

// modifyBootConfig modifies the boot configuration file by inserting autoinstall parameters after "quiet" in "linux" or "append" lines.
// It skips if the file does not exist or if the parameter is already present.
func modifyBootConfig(path string) error {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil // Skip if file does not exist
	} else if err != nil {
		return fmt.Errorf("failed to stat file %s: %w", filepath.Base(path), err)
	}

	// Read file content
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filepath.Base(path), err)
	}

	// Split content into lines
	lines := strings.Split(string(content), "\n")
	modified := false

	for i, line := range lines {
		// Trim whitespace from the beginning to handle leading spaces
		trimmedLine := strings.TrimLeft(line, " \t")
		if trimmedLine == "" {
			continue
		}

		// Handle "append" or "linux" lines ending with "---"
		if (strings.HasPrefix(trimmedLine, "append") || strings.HasPrefix(trimmedLine, "linux")) && strings.HasSuffix(line, "---") {
			// Check if autoinstall is already present
			if !strings.Contains(line, "autoinstall") {
				// Remove trailing "---" temporarily
				baseLine := strings.TrimSuffix(line, "---")
				// Find the position after "quiet"
				parts := strings.Fields(baseLine) // Split into words
				var newParts []string
				quietFound := false

				for _, part := range parts {
					newParts = append(newParts, part)
					if part == "quiet" {
						quietFound = true
						// Insert autoinstall params after "quiet"
						newParts = append(newParts, InsertParams)
					}
				}

				// If "quiet" is not found, append at the end
				if !quietFound && len(parts) > 0 {
					newParts = append(newParts, InsertParams)
				}

				// Reconstruct the line with "---"
				newLine := strings.Join(newParts, " ") + " ---"
				lines[i] = newLine
				modified = true
			}
		}
	}

	// If no modification, return early
	if !modified {
		return nil
	}

	// Join lines back and write to file
	newContent := []byte(strings.Join(lines, "\n"))
	if err := os.WriteFile(path, newContent, FilePerm); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filepath.Base(path), err)
	}

	return nil
}
