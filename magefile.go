//go:build mage
// +build mage

package main

import (
	"io"
	"net/http"
	"os"

	"github.com/aserto-dev/mage-loot/common"
	"github.com/aserto-dev/mage-loot/deps"
	"github.com/magefile/mage/mg"
)

// Deps installs dependency tools for the project
func Deps() {
	deps.GetAllDeps()
}

// Builds the binary
func Build() error {
	mg.SerialDeps(Shoe)

	return common.Build()
}

// Gets a random shoe image and makes ascii art
// It writes it to shoe.txt
func Shoe() error {
	toAscii := deps.BinDep("ascii-image-converter")

	err := downloadFile("shoe.jpg", "https://api.lorem.space/image/shoes?w=400&h=400")
	if err != nil {
		return err
	}

	return toAscii("./shoe.jpg", "--save-txt", "./cmd/shoe/", "--only-save")
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
