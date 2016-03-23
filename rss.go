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
	"bytes"
	"html/template"
	"log"
	"time"
)

type Item struct {
	Title   string
	Desc    string
	PubDate template.HTML
	Link    string
}

type RssRender struct {
	RssTitle string
	RssLink  string
	RssDesc  string
	RssNow   template.HTML
	Items    []Item
}

func getNow() string {
	// zone, _ := time.LoadLocation("UTC")
	return time.Now().Format(time.RFC822Z)
}

func NewRssRederWithStarred(in StarredList) *RssRender {
	r := new(RssRender)
	r.RssTitle = "Github Starred List"
	r.RssLink = "http://github.com"
	r.RssDesc = "RSS for Github Starred List"
	r.RssNow = template.HTML(getNow())

	// log.Println("Len:", len(in))
	for _, v := range in {
		var item Item
		item.Title = v.FullName
		item.Desc = v.Description
		item.Link = v.URL
		item.PubDate = template.HTML(v.CreatedAt.Format(time.RFC822Z))
		r.Items = append(r.Items, item)
	}

	return r
}

func NewRssRederWithFollower(in FollowerList) *RssRender {
	r := new(RssRender)
	r.RssTitle = "Github Follower List"
	r.RssLink = "http://github.com"
	r.RssDesc = "RSS for Github Follower List"
	r.RssNow = template.HTML(getNow())

	for _, v := range in {
		// log.Println("Follower:", v)
		var item Item
		item.Title = v.Login
		item.Desc = v.Login
		item.Link = v.URL
		item.PubDate = template.HTML(getNow())
		r.Items = append(r.Items, item)
	}

	return r
}

func NewRssRederWithFollowing(in FollowingList) *RssRender {
	r := new(RssRender)
	r.RssTitle = "Github Following List"
	r.RssLink = "http://github.com"
	r.RssDesc = "RSS for Github Following List"
	r.RssNow = template.HTML(getNow())

	for _, v := range in {
		var item Item
		item.Title = v.Login
		item.Desc = v.Login
		item.Link = v.URL
		item.PubDate = template.HTML(getNow())
		r.Items = append(r.Items, item)
	}

	return r
}

func (r *RssRender) render() string {
	t, err := template.New("atom").Parse(AtomTmpl)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("parse template")

	buf := new(bytes.Buffer)
	err = t.Execute(buf, r)
	// log.Println("execute it")
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("get rss:", buf.String())
	return buf.String()
}
