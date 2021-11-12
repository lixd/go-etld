package etld

import (
	"strings"
)

var TLDs = make(map[string]struct{}, 10000)

type Host struct {
	Subdomain string
	Domain    string
	TLD       string
}

func Parse(host string) Host {
	var h Host

	parts := strings.Split(host, ".")

	if len(parts) == 0 {
		h.Domain = host
		return h
	}

	var tld string
	for i := len(parts) - 1; i >= 0; i-- {
		p := parts[i]

		if tld == "" {
			tld = p
		} else {
			tld = p + "." + tld
		}

		if _, ok := match(strings.ToLower(tld)); ok {
			h.TLD = tld
			h.Domain = "" // 匹配到后缀之后需要把domain清空
		} else if h.Domain == "" {
			h.Domain = p
			h.Subdomain = "" // 同理 匹配到主域名之后需要把二级域名清空
		} else {
			h.Subdomain = p
		}
	}

	return h
}

func match(suffix string) (string, bool) {
	// 1.精确匹配
	_, ok := TLDs[suffix]
	if ok {
		return suffix, true
	}
	// 2.通配符 匹配
	// 例如: a.b.c 匹配 *.b.c 当前实现为 直接把c替换成*再去匹配
	split := strings.Split(suffix, ".")
	if len(split) < 2 { // 至少要用两段域名时才进行通配符匹配
		return "", false
	}
	wildcard := strings.Replace(suffix, split[0], "*", 1)
	_, ok = TLDs[wildcard]
	if ok {
		return suffix, true
	}
	return "", false
}
