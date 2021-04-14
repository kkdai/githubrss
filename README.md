githubrss:A github notification (starred, follower, followed) RSS feed in Golang
==============

 [![GoDoc](https://godoc.org/github.com/kkdai/githubrss?status.svg)](https://godoc.org/github.com/kkdai/githubrss) 
![Go](https://github.com/kkdai/githubrss/workflows/Go/badge.svg)


Github RSS will help you to generate RSS feed string of specific user ID with  (Starred, Following, Followed).

You can use this to write a web service to host specific RSS for other services like [IFTTT](http://ifttt.com)


Installation and Usage
=============


Install
---------------

```go
go get github.com/kkdai/githubrss
```

Usage
---------------

```go

package main

import (
    "github.com/kkdai/githubrss"
    "fmt"
)

func main() {
	g := NewGithubRss("kkdai")
	//Get latest 5 starred RSS
	ret, err := g.GetStarred(5)

	if err != nil {
		fmt.Println("Not get starred response:", err)
	}

	fmt.Println("Starred RSS:", ret)


	//Get latest 5 follower RSS
	ret, err = g.GetFollower(5)

	if err != nil {
		fmt.Println("Not get follower response:", err)
	}

	fmt.Println("Follower RSS:", ret)


	//Get latest 5 following RSS
	ret, err = g.GetFollowing(5)

	if err != nil {
		fmt.Println("Not get following response:", err)
	}

	fmt.Println("Following RSS:", ret)
}

```

Inspired By
=============

- [Github API doc](https://developer.github.com/v3/users/)
- [https://github.com/lukeforeman/fever-google-starred-importer/blob/master/import.php](https://github.com/lukeforeman/fever-google-starred-importer/blob/master/import.php)
- [http://stackoverflow.com/questions/14893287/creating-an-rss-feed-for-github-stars](http://stackoverflow.com/questions/14893287/creating-an-rss-feed-for-github-stars)
- [https://github.com/thomasyung/twitter-json-to-rss/blob/master/twitter_json_to_rss.php](https://github.com/thomasyung/twitter-json-to-rss/blob/master/twitter_json_to_rss.php)
- [https://gist.github.com/Sanjo/4ed367c68acc27fd9a18](https://gist.github.com/Sanjo/4ed367c68acc27fd9a18)


Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

