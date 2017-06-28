package mobileapp

import (
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
	//"unicode/utf8"
)

// App type
type App struct {
	Kind Kind
	ID   string
}

const prefix = "mobileapp::"

// Kind type
type Kind byte

// Kind types
const (
	Unknown      Kind = '0'
	IOS               = '1'
	Android           = '2'
	WindowsPhone      = '3'
)

var reg = regexp.MustCompile("^" + prefix + `(\d+)-(\S+)`)

func parseRegexp(str string) (*App, bool) {
	if !strings.HasPrefix(str, prefix) {
		return nil, false
	}

	matches := reg.FindAllStringSubmatch(str, 1)
	if len(matches) == 0 || len(matches[0]) != 3 {
		return nil, false
	}

	//kind, _ := utf8.DecodeRuneInString(matches[0][1])
	kind := Kind(matches[0][1][0])
	if kind != IOS && kind != Android && kind != WindowsPhone {
		return nil, false
	}

	return &App{
		Kind: kind,
		ID:   matches[0][2],
	}, true
}

func parseSplit(str string) (*App, bool) {
	if !strings.HasPrefix(str, prefix) {
		return nil, false
	}

	matches := strings.SplitN(str, "::", 2)
	if len(matches) != 2 {
		return nil, false
	}

	matches = strings.SplitN(matches[1], "-", 2)
	if len(matches) != 2 {
		return nil, false
	}

	//kind, _ := utf8.DecodeRuneInString(matches[0])
	kind := Kind(matches[0][0])
	if kind != IOS && kind != Android && kind != WindowsPhone {
		return nil, false
	}

	matches = strings.SplitN(matches[1], " ", -1)

	return &App{
		Kind: kind,
		ID:   matches[0],
	}, true
}

func detectAtoi(str string) Kind {
	_, err := strconv.Atoi(str)
	if err == nil {
		return IOS
	}

	return Android
}

func detectScan(str string) Kind {
	s := scanner.Scanner{}
	s.Init(strings.NewReader(str))
	if s.Scan() == scanner.Int && s.Scan() == scanner.EOF {
		return IOS
	}

	return Android
}

func detectBytes(str string) Kind {
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return Android
		}
	}

	return IOS
}
