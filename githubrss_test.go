// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
