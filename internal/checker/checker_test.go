package checker

import "testing"

func TestDefaultSites(t *testing.T) {
	sites := DefaultSites()
	if len(sites) == 0 {
		t.Error("DefaultSites should return at least one site")
	}
}