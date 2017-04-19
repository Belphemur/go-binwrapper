# Golang Binary Wrapper

Inspired by and partially ported from npm package [bin-wrapper](https://github.com/kevva/bin-wrapper)

## Install

```go get -u github.com/nickalie/go-binwrapper```

## Example of usage

Create wrapper for [cwebp](https://developers.google.com/speed/webp/docs/cwebp)

```
package main

import (
	"github.com/nickalie/go-binwrapper"
	"fmt"
)

func main() {
  base := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/"

  bin := binwrapper.NewBinWrapper().
    Src(
    binwrapper.NewSrc().
      Url(base + "libwebp-0.6.0-mac-10.12.tar.gz").
      Os("darwin")).
    Src(
    binwrapper.NewSrc().
      Url(base + "libwebp-0.6.0-linux-x86-32.tar.gz").
      Os("linux").
      Arch("x86")).
    Src(
    binwrapper.NewSrc().
      Url(base + "libwebp-0.6.0-linux-x86-64.tar.gz").
      Os("linux").
      Arch("x64")).
    Src(
    binwrapper.NewSrc().
      Url(base + "libwebp-0.6.0-windows-x64.zip").
      Os("win32").
      Arch("x64").
      ExecPath("cwebp.exe")).
    Src(
    binwrapper.NewSrc().
      Url(base + "libwebp-0.6.0-windows-x86.zip").
      Os("win32").
      Arch("x86").
      ExecPath("cwebp.exe")).
    Strip(2).
    Dest("vendor/cwebp").
    ExecPath("cwebp")

  err := bin.Run("-version")

  fmt.Printf("stdout: %s\n", string(bin.StdOut))
  fmt.Printf("stderr: %s\n", string(bin.StdErr))
  fmt.Printf("err: %v\n", err)
}
```

It downloads cwebp distribution according to current platform and runs *cwebp* with *-version* argument.

**Important note**: Many vendors don't provide binaries for some specific platforms. For instance, common linux binaries won't work on alpine linux or arm-based linux. In that case you need to have prebuilt binaries on target platform and use SkipDownload. The example above will look like:

```
bin = binwrapper.NewBinWrapper().
		SkipDownload().
		ExecPath("cwebp")
```

Now binwrapper will run *cwebp* located in **PATH**

Use Dest to specify directory with binary:

```
bin = binwrapper.NewBinWrapper().
    SkipDownload().
    Dest("/path/to/directory").
    ExecPath("cwebp")
```
