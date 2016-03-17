package githubrss

import (
	"log"
	"testing"
)

func TestBasicConnect(t *testing.T) {
	//Input user id
	g := NewGithubRss("kkdai")
	//Get starred
	ret, err := g.GetStarred(5)

	if err != nil {
		t.Error("Not get starred response:", err)
	}

	log.Println("Starred:", ret)

	//Get follower
	ret, err = g.GetFollower(5)

	if err != nil {
		t.Error("Not get follower response:", err)
	}

	log.Println("Follower:", ret)

	//Get following
	ret, err = g.GetFollowing(5)

	if err != nil {
		t.Error("Not get following response:", err)
	}

	log.Println("Following:", ret)
}
