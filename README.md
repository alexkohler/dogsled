# dogsled [![Build Status](https://travis-ci.com/alexkohler/dogsled.svg?branch=master)](https://travis-ci.com/alexkohler/dogsled)

dogsled is a Go static analysis tool to find assignments/declarations with too many blank identifiers (e.g. `x, _, _, _, := f()`). Its name was inspired from [this reddit post](https://www.reddit.com/r/golang/comments/9syjj8/what_are_some_red_flags_for_you_in_go_code_reviews/e8sgygf/)

## Installation

    go get -u github.com/alexkohler/dogsled/cmd/dogsled

## Usage

Similar to other Go static analysis tools (such as golint, go vet), dogsled can be invoked with one or more filenames, directories, or packages named by its import path. dogsled also supports the `...` wildcard. By default, it will search for assignment with more than two blank identifiers.

    dogsled [flags] files/directories/packages


### Flags
- **-tests** (default true) - Include test files in analysis
- **-n** (default 2) - Include test files in analysis
- **-set_exit_status** (default false) - Set exit status to 1 if any issues are found.

NOTE: by default, dogsled will check for typos in every identifier (functions, function calls, variables, constants, type declarations, packages, labels). In this case, no flag needs specified. Due to a lack of frequency, there are currently no flags to find only type declarations, packages, or labels.

## Example uses in popular Go repos


Some examples from the [Go standard library](https://github.com/golang/go)

```Bash
$ dogsled go/src/...
go/src/crypto/elliptic/elliptic_test.go:553: declaration has 3 blank identifiers: priv, _, _, _ := GenerateKey(p256,go/src/image/names.go:46: declaration has 3 blank identifiers: _, _, _, a := c.C.RGBA()
go/src/image/color/color.go:232: declaration has 3 blank identifiers: _, _, _, a := c.RGBA()
go/src/image/color/color.go:240: declaration has 3 blank identifiers: _, _, _, a := c.RGBA()
go/src/internal/cpu/cpu_x86.go:67: declaration has 3 blank identifiers: maxID, _, _, _ := cpuid(0, 0)
go/src/internal/cpu/cpu_x86.go:100: declaration has 3 blank identifiers: _, ebx7, _, _ := cpuid(7, 0)
go/src/math/big/natconv_test.go:286: declaration has 3 blank identifiers: x, _, _, _ = x.scan(strings.NewReader(pi),go/src/reflect/value.go:172: declaration has 3 blank identifiers: pc, _, _, _ := runtime.Caller(2)
go/src/reflect/makefunc.go:62: declaration has 4 blank identifiers: _, _, _, stack, _ := funcLayout(t, nil)
go/src/reflect/makefunc.go:113: declaration has 4 blank identifiers: _, _, _, stack, _ := funcLayout(funcType, nil)
go/src/runtime/softfloat64.go:309: declaration has 4 blank identifiers: _, _, _, _ = fi, fn, gi, gn
go/src/runtime/symtab_test.go:54: declaration has 3 blank identifiers: _, _, line, _ := runtime.Caller(1)
go/src/syscall/syscall_unix_test.go:182: declaration has 3 blank identifiers: _, oobn, _, _, err := uc.ReadMsgUnixgo/src/time/time.go:498: declaration has 3 blank identifiers: year, _, _, _ := t.date(false)
go/src/time/time.go:504: declaration has 3 blank identifiers: _, month, _, _ := t.date(true)
go/src/time/time.go:510: declaration has 3 blank identifiers: _, _, day, _ := t.date(true)
go/src/time/time.go:624: declaration has 3 blank identifiers: _, _, _, yday := t.date(false)
```

## Contributing

Please open an issue and/or a PR for any features/bugs. 


## Other static analysis tools

If you've enjoyed dogsled, take a look at my other static anaylsis tools!
- [prealloc](https://github.com/alexkohler/prealloc) - Finds slice declarations that could potentially be preallocated.
- [nakedret](https://github.com/alexkohler/nakedret) - Finds naked returns.
- [identypo](https://github.com/alexkohler/unimport) - Finds typos in variable names, function names, constants, and more!
