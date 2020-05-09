package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const VERSION = "v1.0"

var whitelistedFiles = []string{"site/", ".git/", ".gitignore", ".nojekyll", "CNAME"}
var separator = string(filepath.Separator)
var siteDir = "site" + separator

func main() {
	// Initialize
	workingDirectory, err := os.Getwd()
	if err != nil {
		handleError(err)
	}
	log.Println("Burst Wiki Site Updater " + VERSION + " working in " + workingDirectory)

	// Checkout master, get contents of 404 file (redirect script)
	runCommandToCompletion("git", "checkout", "master")
	redirectScript, err := ioutil.ReadFile("docs/404.html")
	if err != nil {
		handleError(err)
	}

	// Build site
	if _, err := os.Stat(siteDir); !os.IsNotExist(err) {
		log.Println("Found old \"site\" directory - removing")
		if err := os.RemoveAll(siteDir); err != nil {
			handleError(err)
		}
	}
	runCommandToCompletion("mkdocs", "build")

	// Checkout gh-pages branch and remove old files
	log.Println("Switching to gh-pages branch")
	runCommandToCompletion("git", "checkout", "gh-pages")
	log.Println("Deleting old site files")
	var filesToDelete []string
	if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			handleError(err)
		}
		if path == "." {
			return nil
		}
		if !isWhitelisted(path, info) {
			filesToDelete = append(filesToDelete, path)
		}
		return nil
	}); err != nil {
		handleError(err)
	}
	deleteAll(filesToDelete)

	// Update 404 file with redirect script
	notFoundFileContents, err := ioutil.ReadFile("site/404.html")
	if err != nil {
		handleError(err)
	}
	updatedNotFoundContents := strings.Replace(string(notFoundFileContents), "<head>", "<head>"+string(redirectScript), 1)
	notFoundFile, err := os.Create("site/404.html")
	if err != nil {
		handleError(err)
	}
	if _, err := notFoundFile.WriteString(updatedNotFoundContents); err != nil {
		handleError(err)
	}
	if err := notFoundFile.Close(); err != nil {
		handleError(err)
	}

	// Copy new files
	log.Println("Copying new site files")
	if err := filepath.Walk(siteDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			handleError(err)
		}
		if err := copyFile(path, strings.Replace(path, siteDir, "", 1)); err != nil {
			handleError(err)
		}
		return nil
	}); err != nil {
		handleError(err)
	}

	// Delete site directory
	log.Println("Deleting site/ directory")
	if err := os.RemoveAll(siteDir); err != nil {
		handleError(err)
	}

	// Commit
	log.Println("Committing to git")
	runCommandToCompletion("git", "add", ".")
	runCommandToCompletion("git", "commit", "--author=\"Auto Update Script <autoupdate@burstappsteam.org>", "-m", "Auto-Update Site at "+time.Now().String())
	runCommandToCompletion("git", "push")
}

func isWhitelisted(path string, info os.FileInfo) bool {
	// Basic check
	for _, whitelistedFile := range whitelistedFiles {
		if whitelistedFile == info.Name() {
			return true
		}
	}

	// Check directories
	for _, whitelistedFile := range whitelistedFiles {
		if strings.HasSuffix(whitelistedFile, "/") { // It's a directory
			if strings.ReplaceAll(whitelistedFile, "/", "") == info.Name() {
				return true
			}
			if strings.Contains(path, strings.ReplaceAll(whitelistedFile, "/", separator)) {
				return true
			}
		}
	}

	return false
}

func handleError(err error) {
	panic(err)
}

func deleteAll(paths []string) {
	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				handleError(err)
			}
		}
		if fileInfo.IsDir() {
			if err := os.RemoveAll(path); err != nil {
				handleError(err)
			}
		} else {
			if err := os.Remove(path); err != nil {
				handleError(err)
			}
		}
	}
}

func copyFile(src, dst string) error {
	log.Println("copy " + src + " to " + dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return nil
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func runCommandToCompletion(command ...string) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		handleError(err)
	}
}
