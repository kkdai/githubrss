githubrss:
==============

 [![GoDoc](https://godoc.org/github.com/kkdai/nfa?status.svg)](https://godoc.org/github.com/kkdai/nfa)  [![Build Status](https://travis-ci.org/kkdai/nfa.svg?branch=master)](https://travis-ci.org/kkdai/nfa)
![](https://validator.w3.org/feed/images/valid-rss-rogers.png)


How to use it?
=============




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

	fmt.Println("Starred:", ret)


	//Get latest 5 follower RSS
	_, err = g.GetFollower(5)

	if err != nil {
		fmt.Println("Not get follower response:", err)
	}

	fmt.Println(" get follower:", err)
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

This package is licensed under MIT license. See LICENSE for details.
