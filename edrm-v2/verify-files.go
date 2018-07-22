package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/grokify/gotilla/crypto/sha1util"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

type Files struct {
	Files []File `xml:"file"`
}

type File struct {
	Name          string `json:"name" xml:"name,attr"`
	Source        string `json:"source" xml:"source,attr"`
	Mtime         int64  `json:"mtime" xml:"mtime"`
	Size          int64  `json:"size" xml:"size"`
	MD5           string `json:"md5" xml:"md5"`
	CRC32         string `json:"crc32" xml:"crc32"`
	SHA1          string `json:"sha1" xml:"sha1"`
	Format        string `json:"format" xml:"format"`
	LocalVerified bool   `json:"localVerified"`
}

func readFilesInfo(file string) (Files, error) {
	files := Files{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return files, err
	}

	return files, xml.Unmarshal(bytes, &files)
}

type FileStatuses struct {
	Verified []string
	Corrupt  []string
	Missing  []string
}

func verifyFiles(files Files, dir string) FileStatuses {
	fileStatuses := FileStatuses{
		Verified: []string{},
		Corrupt:  []string{},
		Missing:  []string{},
	}

	for _, f := range files.Files {
		if f.Format != "ZIP" {
			continue
		}
		path := filepath.Join(dir, f.Name)
		sha1, err := sha1util.HashFile(path)

		if err != nil {
			if strings.Index(err.Error(), "no such file or directory") > -1 {
				fileStatuses.Missing = append(fileStatuses.Missing, f.Name)
			} else {
				log.Fatal(err)
			}
		}
		if sha1 == f.SHA1 {
			fileStatuses.Verified = append(fileStatuses.Verified, f.Name)
			fmtutil.PrintJSON(fileStatuses)
		} else {
			fileStatuses.Corrupt = append(fileStatuses.Corrupt, f.Name)
			fmtutil.PrintJSON(fileStatuses)
		}
	}
	return fileStatuses
}

func main() {
	file := "edrm.enron.email.data.set.v2.xml_files.xml"
	dir := "path/to/files"

	files, err := readFilesInfo(file)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(files)

	sha1, err := sha1util.HashFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(sha1))

	fileStatuses := verifyFiles(files, dir)
	fmtutil.PrintJSON(fileStatuses)

	fmt.Println("DONE")
}
