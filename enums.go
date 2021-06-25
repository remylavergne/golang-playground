package main

import "strings"

type FileExtension int
type RawExtension string

const (
	Webm FileExtension = iota
	Jpg
	Gif
)

func (fe FileExtension) String() string {
	return []string{"webm", "jpg", "gif"}[fe]
}

func ToFileExtension(rawExt string) (FileExtension, bool) {
	re := strings.ToLower(rawExt)
	fileExtensionMap := map[string]FileExtension{"webm": Webm, "jpg": Jpg, "gif": Gif}

	value, ok := fileExtensionMap[re]

	return value, ok
}
