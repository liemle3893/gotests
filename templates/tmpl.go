// Code generated by "esc -include=.*\.tmpl -o=tmpl.go -pkg=templates ./"; DO NOT EDIT.

package templates

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/complex/call.tmpl": {
		name:    "call.tmpl",
		local:   "complex/call.tmpl",
		size:    235,
		modtime: 1634504572,
		compressed: `
H4sIAAAAAAAC/0ROTYrCUAze9xShdDEDJQcYmAN0Mwwqug6vqQbap6RRkZC7S2u1qw++f/eWO8kMZaK+
LyPc72InwA0nlhvrxEgH+WyAzbg1vSaLMEN3zu2kvp2AESuLfzRwxJe7Uj4yVFJDxT38/AL+k9LAxjou
7ZVE1PDJrnsHFXt9WKQ5Oi3NLtw9LozNuCcVaiVFIK4XZvhesCieAQAA//8OtTkD6wAAAA==
`,
	},

	"/complex/function.tmpl": {
		name:    "function.tmpl",
		local:   "complex/function.tmpl",
		size:    2096,
		modtime: 1635780102,
		compressed: `
H4sIAAAAAAAC/8xUTWvbQBC961dMhSlSUOS7wYeSfpyalDSQQylBSCN7qbQysyOnYdn/Xnb15ZVlmhwK
XbAZMTO7M/PePK0LLIVECMtW5iwaGRoTaH0NqxI2W0iNCQLrAq3TB1R8m9VoTMRwxahYyF36EIMOAAD4
5YDwVDf5LwWKqc25dxj3f8wIImfZI6Tgry50C/b+yObdMFVwtWusnd40kqmpKqQYrvpr9ZhvDyG3JOF9
5/R9Jji3FHJ7GN6L/fj1Gh7uPt4tZDFmVDTP8k2J8TQRzv1p2COzGoeimITcjZ68kaXYtYTUj2XoPYbI
wvIseA/pPeYojkjGgNbDhwXLwfRyQGddA8rCmMSZlMkdgsPwW0ZZjYykXJj79JOTKTseSzsiiVIgAXSl
nXLg7I17VG3F3QOPmeSL9ztLlLYni6b6RNTYvpDI/hqal9INmJWl54+fnE9T7aGAa/hQFGBrgyLjDPZI
OAb56FkgNhB+3zdtVYBq8xyVChMvZoJk879gMutiiYf+kvi74fc34LpZhvWfofr2XrReujNTConTW1FF
nCBRPFx9qene7PxlQ/CUALPlU9cqnwoNp/etjJhTS5XkfEQzORhlbLOFXshu8XnSsohjL77AEmnMSj8L
KdQ+8mOcbkVLeYM2zZzrNXwRR/RRt/o7MtfziPJEjt9tQYpqAZQatlPYKNfxBdXt4fr7agwLoXU3/JWQ
Bf5OYIUV1igdLEsb4ujVBbvsgXbd5gzZp0/Ym5jTaZ2j+mxsj/vZ1IayLtL/laT3qtCasT5UGSOEQlZC
YgirckZYuwTzapjTYVsjfpXmvrW8aSKj3Jog6McbBH8CAAD///AEoo8wCAAA
`,
	},

	"/complex/header.tmpl": {
		name:    "header.tmpl",
		local:   "complex/header.tmpl",
		size:    171,
		modtime: 1635822239,
		compressed: `
H4sIAAAAAAAC/1SMMQ7CMBAEe79ilQoK/AkqGsQXTnhJImQnMu5O+3ekOCDodm504574mAsxTLTEOkgB
ANyrlZGI5yVnlvaS3ONXsiScdlrt/rSRcI+3PqWwmTmvS204/Bcv27UHr5Yp9dc2/eb3fQwfCu8AAAD/
/143JTOrAAAA
`,
	},

	"/complex/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "complex/inline.tmpl",
		size:    51,
		modtime: 1634501989,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NZycQEC
AAD//7mD0zgzAAAA
`,
	},

	"/complex/input.tmpl": {
		name:    "input.tmpl",
		local:   "complex/input.tmpl",
		size:    178,
		modtime: 1634502006,
		compressed: `
H4sIAAAAAAAC/0yNMQoCQQxF+z1FWLaUOYDgAWxE8ASRzSxTTJQkW4XcXSZYTJX8x+d9951qY4K18fc0
XSPctwrXG5Txtgr8MSiv822kpjN7YKc9wqy4E+8RjJ0u8A/Z22p5SmO7p3xAQT4oOQp2MhJNBcqhxT3p
mJ4881mWXwAAAP//ZOzibLIAAAA=
`,
	},

	"/complex/message.tmpl": {
		name:    "message.tmpl",
		local:   "complex/message.tmpl",
		size:    203,
		modtime: 1634502018,
		compressed: `
H4sIAAAAAAAC/zyNUWqEQBBE/+cUjSgkoH2AQA6QnyBJyP8klqZBJ+7M6LI0fffFgfWriqrileqAUQKo
WpCSn1BRZ+ZUZaTwn4k/t5+MlJNZc2FSRRjMVK+S/4g/8AvZEY+Ev24r+NvPG8z4HPK7X2D2VIjcRwn5
LazbAVSNPkygWlqqMdPLK3Hvo1+QEUsvI9Vi1j5+m/3kFnl2qh0V69w9AAD//wfGdPbLAAAA
`,
	},

	"/complex/result.tmpl": {
		name:    "result.tmpl",
		local:   "complex/result.tmpl",
		size:    170,
		modtime: 1634502027,
		compressed: `
H4sIAAAAAAAC/1yNPQrDMAyF95ziETyGHKDQsXTvDQqRiyDY8OxMQncvSk0LnfTzfdIz2yRrEcyUduy9
ze4w47O8BEkXJNlxuWJ9fLC7mWYkdV9gJmWLzb12rNGMWXMc9IOl3cjKkIUcHKdQ+X36L0fgzz3rNL0D
AAD//x2ARVmqAAAA
`,
	},

	"/convey/call.tmpl": {
		name:    "call.tmpl",
		local:   "convey/call.tmpl",
		size:    235,
		modtime: 1635822239,
		compressed: `
H4sIAAAAAAAC/0ROTYrCUAze9xShdDEDJQcYmAN0Mwwqug6vqQbap6RRkZC7S2u1qw++f/eWO8kMZaK+
LyPc72InwA0nlhvrxEgH+WyAzbg1vSaLMEN3zu2kvp2AESuLfzRwxJe7Uj4yVFJDxT38/AL+k9LAxjou
7ZVE1PDJrnsHFXt9WKQ5Oi3NLtw9LozNuCcVaiVFIK4XZvhesCieAQAA//8OtTkD6wAAAA==
`,
	},

	"/convey/function.tmpl": {
		name:    "function.tmpl",
		local:   "convey/function.tmpl",
		size:    1109,
		modtime: 1636889261,
		compressed: `
H4sIAAAAAAAC/5xTwY6bMBC98xUjlANUife+Ug9tmvaWrJKo2yuCR2LVjCPbZBtZ/vfKZKGETaXVcsL4
zZt57w3eV6glg9K65dJJzWkIifcLmtX0+JlECEkSr8h7sYd166JBCJmjTw7WST6IfU4+ISJaaj7jkqXe
0wClENI5uTlFjqxHxqfR5e+lMyp2Oeh4Emu8LDU7o5WCyVw+YCvUMEOF+C5Z2mOWJwPgXBjKhpP39CLd
kcQWJeQZJgTyvj9ETZ2aywnd24LAVQjjclPwAVcVT4UpGjgYO4JcYd3VLd2IZEo7Grd36vkIJtYEY7Qh
XZatSXur/E2zhwfatWUJa6ksLKZ3P+QZPP2433zb0JeqIqsbEP6cDKyVmukIA0E4PHaW2mldnOqjXGL1
62m13Ge5mHi1IFnTrBYbVpctXGvYbhg/C9Vi4uu9ELawrXK2c/q5YPdqemdwXCDvHZqTKhwolawkI6VZ
/SYwgrJxJT/Ub/4v016PeFWyivmFEHMcgd4/1+2i9JZPY/jPiDeYnc7GE++OulXVWruvWEuVRwfetPL+
jpSdzmBMT3CtvrPT4fqPhjwJSdLrSP4GAAD//4RU5ldVBAAA
`,
	},

	"/convey/header.tmpl": {
		name:    "header.tmpl",
		local:   "convey/header.tmpl",
		size:    236,
		modtime: 1635822239,
		compressed: `
H4sIAAAAAAAC/1SMQY7DIAxF9z6FxWpmMXCIWc1m1CtQcEmUgiOglSLLd68a0kpd/S/7/ScS6TIXQjOR
j1SNKohUXxKh/eWcqfSmKmL3B5WIP6qw+rD4RChiT6OqAsx55drxCxAR35a//Tok/z6T6pj1SfUgqcSj
mzT36Xa2gbNLfPUlucxhcYmfYXbGflAt+9q31itRby5x4HKnzY0w8A0vPTwCAAD//xRdGXTsAAAA
`,
	},

	"/convey/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "convey/inline.tmpl",
		size:    51,
		modtime: 1635822239,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NZycQEC
AAD//7mD0zgzAAAA
`,
	},

	"/convey/input.tmpl": {
		name:    "input.tmpl",
		local:   "convey/input.tmpl",
		size:    178,
		modtime: 1635822239,
		compressed: `
H4sIAAAAAAAC/0yNMQoCQQxF+z1FWLaUOYDgAWxE8ASRzSxTTJQkW4XcXSZYTJX8x+d9951qY4K18fc0
XSPctwrXG5Txtgr8MSiv822kpjN7YKc9wqy4E+8RjJ0u8A/Z22p5SmO7p3xAQT4oOQp2MhJNBcqhxT3p
mJ4881mWXwAAAP//ZOzibLIAAAA=
`,
	},

	"/convey/message.tmpl": {
		name:    "message.tmpl",
		local:   "convey/message.tmpl",
		size:    203,
		modtime: 1635822239,
		compressed: `
H4sIAAAAAAAC/zyNUWqEQBBE/+cUjSgkoH2AQA6QnyBJyP8klqZBJ+7M6LI0fffFgfWriqrileqAUQKo
WpCSn1BRZ+ZUZaTwn4k/t5+MlJNZc2FSRRjMVK+S/4g/8AvZEY+Ev24r+NvPG8z4HPK7X2D2VIjcRwn5
LazbAVSNPkygWlqqMdPLK3Hvo1+QEUsvI9Vi1j5+m/3kFnl2qh0V69w9AAD//wfGdPbLAAAA
`,
	},

	"/convey/result.tmpl": {
		name:    "result.tmpl",
		local:   "convey/result.tmpl",
		size:    170,
		modtime: 1635822239,
		compressed: `
H4sIAAAAAAAC/1yNPQrDMAyF95ziETyGHKDQsXTvDQqRiyDY8OxMQncvSk0LnfTzfdIz2yRrEcyUduy9
ze4w47O8BEkXJNlxuWJ9fLC7mWYkdV9gJmWLzb12rNGMWXMc9IOl3cjKkIUcHKdQ+X36L0fgzz3rNL0D
AAD//x2ARVmqAAAA
`,
	},

	"/test/call.tmpl": {
		name:    "call.tmpl",
		local:   "test/call.tmpl",
		size:    0,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/function.tmpl": {
		name:    "function.tmpl",
		local:   "test/function.tmpl",
		size:    19,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/6qu1tdSSFRIzs/NTc0rUdDSr+UCBAAA//+6o7WcEwAAAA==
`,
	},

	"/test/header.tmpl": {
		name:    "header.tmpl",
		local:   "test/header.tmpl",
		size:    0,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "test/inline.tmpl",
		size:    0,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "test/inputs.tmpl",
		size:    0,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/message.tmpl": {
		name:    "message.tmpl",
		local:   "test/message.tmpl",
		size:    0,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/results.tmpl": {
		name:    "results.tmpl",
		local:   "test/results.tmpl",
		size:    0,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/testify/call.tmpl": {
		name:    "call.tmpl",
		local:   "testify/call.tmpl",
		size:    241,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/0SOQWrDQAxFryKMFy0YHaDQA3hTSlvatRjLrsCeFo2SEITuHsY4mdWHP2/el/vEs2SG
LtG6dhHuF7FfwA9OLGfW2sgM+c8Ax/JpekoWYYbunKf6eicBI1qLb7RxxJO7Ul4Yehmg5xVeXgHfSWlj
Yy2HvZeIAR5/296PitUbzJB0KU2/K+riTuPX9Z9xLN+kQpOkCMTG7vF85C0AAP//ZQi8iPEAAAA=
`,
	},

	"/testify/function.tmpl": {
		name:    "function.tmpl",
		local:   "testify/function.tmpl",
		size:    2798,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/5xWUW/cNgx+tn8FYQTFebiq7yn60GHtkJemSIL1oRgG9UzfjPl0N4necBD43wdRlu3z
2VnTPOQsyiQ/8qM+y/sK68YgFNo5tNQcTQGvmXMiNVg2tAXvmxrUvWnPD0idNe6DtUfL7D3h4dRqQih2
um0LUMGIrUNmtNZ7NBVz7v1raGo4WtiYI4F67L4ROnIlqM+2MXRnTh05SQ0AsIX6QOrxFLbqTTHNckDn
9B4lURGAjVuNBJGdUjKiqSTksGLO83wsGv/udFswx1LVh7CUSv8HJ3PdF5bqm0Y9uH0Rdr4nDmzhB6p7
JnndmZ3Q2Df9pobbd8Enz8MWeK+e0NEnfUDmDcFPAVZj9uqpBJ9nweXfhv4E9YA7bP5By5xnWU+funOP
ZLsdiXGwfmywrVy0ZXQ+IdRiAScvh7j921abPc4cMu9lHUAKvPMJ+62RtbAacibT5Hn2GFCFMj9rqw9I
aCWZQNN2fwFsAuvaQxKK6QrdJONlfuE5NN17wRFaXTEf9OmrI9uY/e8xezolX8d1CNCDCq5xdMS/ZDb6
gBAjJOKzNb5SE7SpRtJmfe85ij9DawVSX3gKec3MCgvP9D7LpPHh34LPhIAHdF1LLuX5og091/sh5aUu
ScIkYJDOd9h7n6wfO7Ob8Rha/+YNPN3/cn8L76sKApew0w6dEprro41SuDna8VD3RN25T/ovrMpS0E6o
D8wltv/oudsCURySyHrsQAznJ2WlHBe1hta2LbbMMQjR28uRSNSnFy+gBI/w+9Z7UUjxIvXQmc0qbCLV
LyP4oCXX6rGIcHjelJcgh5O9NsQrsnM1mtJGGbvzCeVlbZlfDSoZ7b/ptkNmn0KsqFHmvYrqeAtEKh4P
NdGo7Rhg1KZsQbCuFn2+BYlJZX6xDQ3VX0jP7Tt49e1M6NTPXV2j9d+TsB+DOJ7zj3d5bb83KF0qYUA2
fn9sPJgF3NQy4PMPfzCvoZifTpj9TaNNbiKrIdf0YqmVsPA3zRcvAZIr3HSI1Kg64Qs8cKAeRXzDGE8+
2G4ffcuRj6ihy3kF3k291PUVn5find4Y2sZgfHUF8xrIZ2p4KZ5fjxTvLC8DMByt1RGfSyQwA5cwvfhl
nHO4HUWX/wIAAP//vN7yje4KAAA=
`,
	},

	"/testify/header.tmpl": {
		name:    "header.tmpl",
		local:   "testify/header.tmpl",
		size:    142,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/0TMMQ7CMAyF4d2nsDrBQC7BxIK4gkUebYXiViGb9e6OlAi6/bL1voiM1+rQaYFl1ImU
iGo+Q9N1KwXePmRE6g941gspuz3fNkMj0mMkKbKWfatNT4dw65cB3K2AHJO2/DhSzv/6BgAA///GzMM9
jgAAAA==
`,
	},

	"/testify/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "testify/inline.tmpl",
		size:    49,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NYCAgAA
//+q60H/MQAAAA==
`,
	},

	"/testify/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "testify/inputs.tmpl",
		size:    177,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/0yNMaoDMQxE+38KsWz58QECOUCaEMgJFCwvLqwESVsJ3T1YpHAlzWN4416pdSbYOn9O
0y3CfW9wuUKZb2/Ab4PyPF9GarqyOw6qEWbFnbhGMA76h1/I3t7KQzrbLeUTCvJByVFwkJFoKlAOLe5J
5/TiWc/fNwAA//94+RPrsQAAAA==
`,
	},

	"/testify/message.tmpl": {
		name:    "message.tmpl",
		local:   "testify/message.tmpl",
		size:    201,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/zyN4WqDQBCE//sUiyi0oPsAhT5A/xRpS/9f4mgW9GLuTkNY9t2DB/HXDDPDN6o9BvGg
ckaMbkRJrVmhKgP5ayL+XU8JMUWz+sakCt+bqd4lXYh/cIZsCHvCf48F/O+mFWZ8DPnbzTB7y0Tugvj0
5Zd1B6oG50dQJQ1VmOjjk7hzwc1ICLmXgSoxa16/9XZws7wXqi1l+wwAAP//kC65UskAAAA=
`,
	},

	"/testify/results.tmpl": {
		name:    "results.tmpl",
		local:   "testify/results.tmpl",
		size:    168,
		modtime: 1636814132,
		compressed: `
H4sIAAAAAAAC/1yNTQrCQAyFr/Iosyw9gOBS3HsDoRkJlAy8ma5C7i6pRcFVfr4vee6rVDXBROn7NvoU
AXc+7SUoOqPIhssVy+ODI9y1omjEDHexNTf3NrBkc85a82DstH4jG1MW8uQ4hMbv0385A3/uUd8BAAD/
/7BPz2GoAAAA
`,
	},

	"/": {
		name:  "/",
		local: `./`,
		isDir: true,
	},

	"/complex": {
		name:  "complex",
		local: `complex`,
		isDir: true,
	},

	"/convey": {
		name:  "convey",
		local: `convey`,
		isDir: true,
	},

	"/test": {
		name:  "test",
		local: `test`,
		isDir: true,
	},

	"/test_empty": {
		name:  "test_empty",
		local: `test_empty`,
		isDir: true,
	},

	"/testify": {
		name:  "testify",
		local: `testify`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"./": {},

	"complex": {
		_escData["/complex/call.tmpl"],
		_escData["/complex/function.tmpl"],
		_escData["/complex/header.tmpl"],
		_escData["/complex/inline.tmpl"],
		_escData["/complex/input.tmpl"],
		_escData["/complex/message.tmpl"],
		_escData["/complex/result.tmpl"],
	},

	"convey": {
		_escData["/convey/call.tmpl"],
		_escData["/convey/function.tmpl"],
		_escData["/convey/header.tmpl"],
		_escData["/convey/inline.tmpl"],
		_escData["/convey/input.tmpl"],
		_escData["/convey/message.tmpl"],
		_escData["/convey/result.tmpl"],
	},

	"test": {
		_escData["/test/call.tmpl"],
		_escData["/test/function.tmpl"],
		_escData["/test/header.tmpl"],
		_escData["/test/inline.tmpl"],
		_escData["/test/inputs.tmpl"],
		_escData["/test/message.tmpl"],
		_escData["/test/results.tmpl"],
	},

	"test_empty": {},

	"testify": {
		_escData["/testify/call.tmpl"],
		_escData["/testify/function.tmpl"],
		_escData["/testify/header.tmpl"],
		_escData["/testify/inline.tmpl"],
		_escData["/testify/inputs.tmpl"],
		_escData["/testify/message.tmpl"],
		_escData["/testify/results.tmpl"],
	},
}
