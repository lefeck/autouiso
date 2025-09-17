package test

import (
	"os"
	"testing"
)

func TestModifyBootConfig(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		exists   bool
	}{
		{
			name:     "NotExist",
			input:    "nonexistent.txt",
			expected: "",
			exists:   false,
		},
		{
			name:     "AlreadyModified",
			input:    "already_modified.txt",
			expected: "    append initrd=/casper/hwe-initrd quiet autoinstall ds=nocloud;s=/cdrom/ ---",
			exists:   true,
		},
		{
			name:     "AppendLine",
			input:    "append_test.txt",
			expected: "    append initrd=/casper/hwe-initrd quiet autoinstall ds=nocloud;s=/cdrom/ ---",
			exists:   true,
		},
		{
			name:     "LinuxLine",
			input:    "linux_test.txt",
			expected: "    linux /casper/vmlinuz iso-scan/filename=${iso_path} quiet autoinstall ds=nocloud;s=/cdrom/ ---",
			exists:   true,
		},
		{
			name:     "LeadingSpaces",
			input:    "leading_spaces.txt",
			expected: "\t\tlinux /casper/vmlinuz quiet autoinstall ds=nocloud;s=/cdrom/ ---",
			exists:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file if it should exist
			var _ *os.File
			if tt.exists {
				tempFile, err := os.CreateTemp("", tt.input)
				if err != nil {
					t.Fatalf("failed to create temp file: %v", err)
				}
				defer os.Remove(tempFile.Name())

				// Write initial content
				initialContent := ""
				switch tt.name {
				case "AlreadyModified":
					initialContent = "    append initrd=/casper/hwe-initrd quiet autoinstall ds=nocloud;s=/cdrom/ ---"
				case "AppendLine":
					initialContent = "    append initrd=/casper/hwe-initrd quiet ---"
				case "LinuxLine":
					initialContent = "    linux /casper/vmlinuz iso-scan/filename=${iso_path} quiet ---"
				case "LeadingSpaces":
					initialContent = "\t\tlinux /casper/vmlinuz quiet ---"
				}
				if _, err := tempFile.Write([]byte(initialContent)); err != nil {
					t.Fatalf("failed to write to temp file: %v", err)
				}
				tempFile.Close()

				// Run the function
				err = modifyBootConfig(tempFile.Name())
				if err != nil {
					t.Errorf("modifyBootConfig() error = %v", err)
					return
				}

				// Read the modified content
				modifiedContent, err := os.ReadFile(tempFile.Name())
				if err != nil {
					t.Fatalf("failed to read modified file: %v", err)
				}

				// Compare with expected
				if string(modifiedContent) != tt.expected+"\n" {
					t.Errorf("modifyBootConfig() got = %q, want %q", string(modifiedContent), tt.expected+"\n")
				}
			} else {
				// Test non-existent file
				err := modifyBootConfig(tt.input)
				if err != nil {
					t.Errorf("modifyBootConfig() error = %v, want nil", err)
				}
			}
		})
	}
}
