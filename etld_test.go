package etld

import (
	"fmt"
	"testing"
)

func TestParseTLD(t *testing.T) {
	tests := []struct {
		in  string
		out Host
	}{
		{in: "internal-Conte-PR-MiddleTier-LB-642703177.eu-west-1.elb.amazonaws.com", out: Host{"", "internal-Conte-PR-MiddleTier-LB-642703177", "eu-west-1.elb.amazonaws.com"}},
		{in: "pressly.com", out: Host{"", "pressly", "com"}},
		{in: "www.pressly.it", out: Host{"www", "pressly", "it"}},
		{in: "www.lixueduan.com", out: Host{"www", "lixueduan", "com"}},
		{in: "www.github.com", out: Host{"www", "github", "com"}},
	}

	for _, tt := range tests {
		h := Parse(tt.in)
		if h.Subdomain != tt.out.Subdomain || h.Domain != tt.out.Domain || h.TLD != tt.out.TLD {
			t.Errorf("expected %v, got %v", tt.out, h)
		}
	}
}

// cpu: Intel(R) Core(TM) i3-9100F CPU @ 3.60GHz
// 1696 ns/op
func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parse("www.lixueduan.com")
	}
}

func TestParse(t *testing.T) {
	parse := Parse("internal-Conte-PR-MiddleTier-LB-642703177.eu-west-1.elb.amazonaws.com")
	fmt.Println(parse)
}

func TestMatch(t *testing.T) {
	suffix, ok := match("eu-west-1.elb.amazonaws.com")
	fmt.Println(suffix, ok)
}

// 238.8 ns/op
func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = match("eu-west-1.elb.amazonaws.com")
	}
}
