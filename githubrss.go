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

import "encoding/json"

type GithubRss struct {
	Id string

	httpClient *client
}

func NewGithubRss(id string) *GithubRss {
	g := new(GithubRss)
	g.Id = id
	g.httpClient = NewClient(g.Id)

	return g
}

func (g *GithubRss) GetStarred(count int) (string, error) {
	//get response
	body, err := g.httpClient.GetStarredObj(count)
	// log.Println("body:", string(body))

	var result StarredList
	err = json.Unmarshal(body, &result)

	if err != nil {
		//error
		return "", err
	}

	// log.Println("Obj=", len(result), result[1])
	r := NewRssRederWithStarred(result)
	return r.render(), nil
}

func (g *GithubRss) GetFollower(count int) (string, error) {
	body, err := g.httpClient.GetFollowerObj(count)

	var result FollowerList
	err = json.Unmarshal(body, &result)

	// log.Println("Body:", string(body), err, result)

	if err != nil {
		//error
		return "", err
	}

	r := NewRssRederWithFollower(result)
	return r.render(), nil
}

func (g *GithubRss) GetFollowing(count int) (string, error) {
	body, err := g.httpClient.GetFollowingObj(count)

	var result FollowingList
	err = json.Unmarshal(body, &result)

	if err != nil {
		//error
		return "", err
	}

	r := NewRssRederWithFollowing(result)
	return r.render(), nil
}
