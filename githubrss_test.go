package githubrss

import (
	"log"
	"testing"
)

func TestBasicConnect(t *testing.T) {
	g := NewGithubRss("kkdai")
	ret, err := g.GetStarred(5)

	if err != nil {
		t.Error("Not get starred response:", err)
	}

	log.Println("Starred:", ret)

	// _, err = g.GetFollower(5)

	// if err != nil {
	// 	t.Error("Not get follower response:", err)
	// }

	// log.Println(" get follower:", err)
}
