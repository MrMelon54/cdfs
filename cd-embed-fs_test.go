package cdfs

import (
	"bytes"
	"embed"
	"testing"
)

var (
	//go:embed assets
	assetsDir embed.FS

	testAssetsContent = []byte("Hello World!\n")
)

func TestCD(t *testing.T) {
	f, err := assetsDir.ReadFile("assets/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(f, testAssetsContent) != 0 {
		t.Fatalf("File '%s' doesn't match expected content", f)
	}

	cd := CD(assetsDir, "assets")
	f2, err := cd.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(f2, testAssetsContent) != 0 {
		t.Fatalf("File '%s' doesn't match expected content", f)
	}
}
