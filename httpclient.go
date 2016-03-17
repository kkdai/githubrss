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
	"io/ioutil"
	"log"
	"net/http"
)

type client struct {
	userId string
}

func NewClient(userId string) *client {
	c := new(client)
	c.userId = userId
	return c
}

func (c *client) GetStarredObj(resultCount int) ([]byte, error) {
	url := getStarredURL(c.userId, resultCount)
	log.Println("URL=", url)
	return getHttpResponse(url)
}

func (c *client) GetFollowerObj(resultCount int) ([]byte, error) {
	url := getFollowerURL(c.userId, resultCount)
	log.Println("URL=", url)
	return getHttpResponse(url)
}

func (c *client) GetFollowingObj(resultCount int) ([]byte, error) {
	url := getFollowingURL(c.userId, resultCount)
	log.Println("URL=", url)
	return getHttpResponse(url)
}

func getHttpResponse(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
