package main

import (
	"testing"
)

func TestFileExtensionEnum(t *testing.T) {
	webm := Webm.String()
	jpg := Jpg.String()
	gif := Gif.String()

	if webm != "webm" || jpg != "jpg" || gif != "gif" {
		t.Fatalf("Enum doesn't match its value")
	}
}

func TestStringToFileExtension(t *testing.T) {
	webm, _ := ToFileExtension("webm")
	jpg, _ := ToFileExtension("jpg")
	gif, _ := ToFileExtension("gif")

	if webm != Webm || jpg != Jpg || gif != Gif {
		t.Fatalf("Raw value doesn't match Enum")
	}
}

func TestStringToFileExtension2(t *testing.T) {
	webm, _ := ToFileExtension("Webm")
	jpg, _ := ToFileExtension("jPg")
	gif, _ := ToFileExtension("giF")

	if webm != Webm || jpg != Jpg || gif != Gif {
		t.Fatalf("Raw value doesn't match Enum")
	}
}

func TestUnknownKeyFileExtension(t *testing.T) {
	enumMap := map[string]FileExtension{"webm": Webm, "jpg": Jpg, "gif": Gif}

	_, ok := enumMap["unknownKey"]

	if ok {
		t.Fatalf("This behavior isn't right")
	}
}
