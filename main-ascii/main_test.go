package main
package main

import (
    "strings"
    "testing"
)

// helper — loads font once and fails the test if it can't
func mustLoadBanner(t *testing.T, name string) map[rune][]string {
    t.Helper() // marks this as a helper so error lines point to the caller
    font, err := loadBanner(name + ".txt")
    if err != nil {
        t.Fatalf("could not load banner %q: %v", name, err)
    }
    return font
}
func TestLoadBanner_CharacterCount(t *testing.T) {
    font := mustLoadBanner(t, "standard")

    // ASCII printable characters go from 32 (space) to 126 (~)
    // That's 95 characters total
    want := 95
    got := len(font)

    if got != want {
        t.Errorf("expected %d characters in font, got %d", want, got)
    }
}