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

package main

import (
	"log"
	"net/http"
	"os"
)

type httpAPI struct {
}

func (h *httpAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//key := r.RequestURI
	switch {
	case r.Method == "GET":
		//call api
	default:
		w.Header().Add("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func serveHttpAPI(port string, existC chan bool) {

	// exit when receivee error channel
	go func() {
		if err, ok := <-existC; ok {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	srv := http.Server{
		Addr:    ":" + port,
		Handler: &httpAPI{},
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
