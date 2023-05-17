package filereader_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	freader "github.com/aniljaiswalcs/exif-reader/app/reader"
)

// TestFilePath tests the FilePath function.
func TestFilePath(t *testing.T) {
	// Create a fake directory for testing
	fakeDir, err := ioutil.TempDir("", "test_dir")
	if err != nil {
		t.Fatalf("Error creating fake directory: %v", err)
	}
	defer os.RemoveAll(fakeDir)

	// Create a subdirectory within the fake directory
	subDir := filepath.Join(fakeDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Error creating subdirectory: %v", err)
	}
	// Create fake files inside the fake directory
	fakeFiles := []string{
		filepath.Join(fakeDir, "file1.txt"),
		filepath.Join(subDir, "file2.txt"),
	}
	for _, file := range fakeFiles {
		if err := ioutil.WriteFile(file, []byte{}, 0644); err != nil {
			t.Fatalf("Error creating fake file: %v", err)
		}
	}

	testCase := struct {
		inputPath    string
		expectedPath []string
	}{
		inputPath: fakeDir,
		expectedPath: []string{
			fakeFiles[0],
			fakeFiles[1],
		},
	}

	testfreader := freader.NewReader()
	actualPath, err := testfreader.FilePath(testCase.inputPath)
	if err != nil {
		t.Errorf("Error retrieving file paths: %v", err)
	}

	// Check if the number of file paths matches
	if len(actualPath) != len(testCase.expectedPath) {
		t.Errorf("Expected %d file paths, but got %d", len(testCase.expectedPath), len(actualPath))
	}

	// Check if each file path is present in the result
	for _, path := range testCase.expectedPath {
		if !containsPath(actualPath, path) {
			t.Errorf("File path not found: %s", path)
		}
	}

}

// Helper function to check if a file path exists in the slice.
func containsPath(paths []string, path string) bool {
	for _, p := range paths {
		if p == path {
			return true
		}
	}
	return false
}
