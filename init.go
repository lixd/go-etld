package etld

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
)

// eTLDList.dat datasource https://publicsuffix.org/list/effective_tld_names.dat

//go:embed eTLDList.dat
var eTLDList string

func init() {
	reader := strings.NewReader(eTLDList)
	br := bufio.NewReader(reader)
	for {
		l, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(l)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		TLDs[line] = struct{}{}
	}
}
