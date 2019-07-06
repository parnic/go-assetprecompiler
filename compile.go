package precompiler

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
)

// FileType represents a supported file extension (without the .)
type FileType string

// supported file types
const (
	CSS FileType = "css"
	JS  FileType = "js"
)

func initMinify() *minify.M {
	minifier := minify.New()
	minifier.AddFunc(string(CSS), css.Minify)
	minifier.AddFunc(string(JS), js.Minify)
	return minifier
}

func supportedFileType(t FileType) bool {
	switch t {
	case CSS, JS:
		return true
	}

	return false
}

func getBytes(config Config, minifier *minify.M) (map[FileType]*bytes.Buffer, error) {
	buf := make(map[FileType]*bytes.Buffer)

	for _, pattern := range config.Files {
		files, _ := filepath.Glob(pattern)
		for _, file := range files {
			if file, err := os.Open(file); err == nil {
				ext := FileType(filepath.Ext(file.Name())[1:])
				if !supportedFileType(ext) {
					fmt.Println("Unsupported file type:", file.Name())
					continue
				}

				if buf[ext] == nil {
					buf[ext] = &bytes.Buffer{}
				}

				if config.Minify {
					if err = minifier.Minify(string(ext), buf[ext], file); err != nil {
						return nil, err
					}
				} else {
					if _, err = buf[ext].ReadFrom(file); err != nil {
						return nil, err
					}
				}
			} else {
				return nil, err
			}
		}
	}

	return buf, nil
}

// CompileResult holds the results of compilation
type CompileResult struct {
	Bytes      []byte
	Hash       string
	OutputPath string
}

func finalize(config Config, buf map[FileType]*bytes.Buffer) (map[FileType]*CompileResult, error) {
	ret := make(map[FileType]*CompileResult, len(buf))

	for key, b := range buf {
		if b.Len() > 0 {
			bytes := b.Bytes()
			hash := sha256.Sum256(bytes)
			ret[key] = &CompileResult{
				Bytes: bytes,
				Hash:  hex.EncodeToString(hash[:]),
			}

			if len(config.OutputDir) > 0 {
				ext := "." + string(key)
				if config.Minify {
					ext = ".min" + ext
				}

				dir := filepath.Join(config.OutputDir, string(key))
				os.MkdirAll(dir, 0755)

				destFile := filepath.Join(dir, "app-"+ret[key].Hash+ext)
				ioutil.WriteFile(destFile, bytes, 0644)
				ret[key].OutputPath = destFile
			}
		}
	}

	return ret, nil
}

// Compile compiles files indicated by the config into a single file per type
func Compile(config Config) (map[FileType]*CompileResult, error) {
	var buf map[FileType]*bytes.Buffer
	var err error
	if buf, err = getBytes(config, initMinify()); err != nil {
		return nil, err
	}

	return finalize(config, buf)
}
