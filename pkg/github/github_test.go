package github

import "testing"

func TestOne(t *testing.T) {
	c := NewClient(true)
	c.CleanupPackages("grimtapi", 5)

	defer t.Fail()
}
