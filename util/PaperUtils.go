package util

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
)

func GetCacheDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, "Library", "Caches"), nil
}

func DownloadImage(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", errors.New("non-200 status code")
	}

	cacheDir, err := GetCacheDir()
	if err != nil {
		return "", err
	}

	file, err := os.Create(filepath.Join(cacheDir, "wallpaper"))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

// SetFromFile uses AppleScript to tell Finder to set the desktop wallpaper to specified file.
func SetFromFile(file string) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}

func SetFromURL(url string) error {
	file, err := DownloadImage(url)
	if err != nil {
		return err
	}

	return SetFromFile(file)
}

func DownloadImageToFile(url string, file string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", errors.New("non-200 status code")
	}

	createfile, err := os.Create(file)
	_, err = io.Copy(createfile, res.Body)
	if err != nil {
		return "", err
	}

	err = createfile.Close()
	if err != nil {
		return "", err
	}

	return createfile.Name(), nil
}
