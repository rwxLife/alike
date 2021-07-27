package backup

import (
	"log"
	"os"
	"testing"
)

// TestExtractName will run a bunch of test cases for the function: extractName
func TestExtractName(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"../word/", "word/"},
		{"/a/b/c/word", "word"},
		{"/a/b/c/file.txt", "file.txt"},
	}

	for _, testCase := range testCases {

		got := extractName(testCase.input)
		if got != testCase.expected {
			log.Fatalf("Input: %s, Expected: %s, Got: %s", testCase.input, testCase.expected, got)
		}
	}
}

// TestExtractParent will test the function: extractParent
func TestExtractParent(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"../word/", "../"},
		{"/a/b/c/word", "/a/b/c/"},
		{"/a/b/c/file.txt", "/a/b/c/"},
	}

	for _, testCase := range testCases {

		got := extractParent(testCase.input)
		if got != testCase.expected {
			log.Fatalf("Input: %s, Expected: %s, Got: %s", testCase.input, testCase.expected, got)
		}
	}
}

// TestIsDirectory will run a bunch of cases for the function: isDirectory
func TestIsDirectory(t *testing.T) {

	testCases := []struct {
		input    string
		expected bool
	}{
		{"/usr", true},
		{"./file_test.go", false},
		{"../", true},
	}

	for _, testCase := range testCases {

		got := isDirectory(testCase.input)

		if got != testCase.expected {
			log.Fatalf("Input: %s, Expected: %t, Got: %t", testCase.input, testCase.expected, got)
		}
	}
}

// TestCreateMetaIfNotFound will test the function: createMetaIfNotFound
func TestCreateMetaIfNotFound(t *testing.T) {

	testCases := []struct {
		input string
	}{
		{"./"},
		{"../"},
	}

	for _, testCase := range testCases {

		createMetaIfNotFound(testCase.input)

		_, err := os.Stat(testCase.input + ".meta")

		if os.IsNotExist(err) {
			log.Fatalf("No file found for input: %s", testCase.input)
		}

		os.Remove(testCase.input + ".meta")
	}
}

func TestMakeDirectory(t *testing.T) {

	testCases := []struct {
		path string
		name string
	}{
		{"./", "root"},
		{"../", "random"},
	}

	for _, testCase := range testCases {

		makeDirectory(testCase.path, testCase.name)

		if !isDirectory(testCase.path + testCase.name) {
			log.Fatalf("Path: %s, Name: %s", testCase.path, testCase.name)
		}

		os.Remove(testCase.path + testCase.name)
	}
}

// TestGetDirectoryListing will test the function: getDirectoryListing
func TestGetDirectoryListing(t *testing.T) {

	sample := []string{".git", ".gitignore", "Makefile", "backup", "bin", "go.mod", "main.go"}
	result := getDirectoryListing("../")

	if len(sample) != len(result) {
		log.Fatal(result)
		log.Fatal("Dir listing is incorrect (length)")
	}

	i := 0
	j := 0
	for i < len(sample) && j < len(result) {

		if sample[i] != result[j] {
			log.Fatal("Dir listing is incorrect")
		}
		i++
		j++
	}
}
