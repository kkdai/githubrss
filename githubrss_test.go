package githubrss

import (
	"log"
	"testing"
)

func TestBasicConnect(t *testing.T) {
	g := NewGithubRss("kkdai")
	_, err := g.GetStarred(5)

	if err != nil {
		t.Error("Not get response:", err)
	}

	log.Println(" get starred:", err)
}
