package binwrapper_test

import (
	"fmt"
	"github.com/belphemur/go-binwrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Example of wrapping cwebp command line tool
func ExampleNewBinWrapper() {
	libwebpVersion := "1.6.0"
	base := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/"

	bin := binwrapper.NewBinWrapper()
	bin.Src(
		binwrapper.NewSrc().
			URL(base + "libwebp-" + libwebpVersion + "-mac-arm64.tar.gz").
			Os("darwin").
			Arch("arm64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-mac-x86-64.tar.gz").
				Os("darwin").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-x86-32.tar.gz").
				Os("linux").
				Arch("x86")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-x86-64.tar.gz").
				Os("linux").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-aarch64.tar.gz").
				Os("linux").
				Arch("arm64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-aarch64.tar.gz").
				Os("linux").
				Arch("aarch64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-windows-x64.zip").
				Os("win32").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-windows-x86.zip").
				Os("win32").
				Arch("x86")).
		Strip(2).
		Dest("vendor/cwebp").
		ExecPath("cwebp")

	err := bin.Run("-version")

	fmt.Printf("stdout: %s\n", string(bin.StdOut()))
	fmt.Printf("stderr: %s\n", string(bin.StdErr()))
	fmt.Printf("err: %v\n", err)
}

func TestNewBinWrapperNoError(t *testing.T) {
	libwebpVersion := "1.6.0"
	base := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/"

	bin := binwrapper.NewBinWrapper()
	bin.Src(
		binwrapper.NewSrc().
			URL(base + "libwebp-" + libwebpVersion + "-mac-arm64.tar.gz").
			Os("darwin").
			Arch("arm64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-mac-x86-64.tar.gz").
				Os("darwin").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-x86-32.tar.gz").
				Os("linux").
				Arch("x86")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-x86-64.tar.gz").
				Os("linux").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-aarch64.tar.gz").
				Os("linux").
				Arch("arm64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-linux-aarch64.tar.gz").
				Os("linux").
				Arch("aarch64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-windows-x64.zip").
				Os("win32").
				Arch("x64")).
		Src(
			binwrapper.NewSrc().
				URL(base + "libwebp-" + libwebpVersion + "-windows-x86.zip").
				Os("win32").
				Arch("x86")).
		Strip(2).
		Dest("test/cwebp").
		ExecPath("cwebp")

	err := bin.Run("-version")
	assert.Nil(t, err)
}

func TestNewBinWrapperError(t *testing.T) {
	bin := binwrapper.NewBinWrapper().
		ExecPath("cwebp")

	err := bin.Run("-version")
	assert.NotNil(t, err)
}
