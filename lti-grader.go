// Copyright <2021> <see contributers file>

// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use, copy,
// modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// this source is derived from an example file located at
// https://github.com/jordic/lti/blob/master/cmd/main.go
// it has been changed to use TLS encryption

package main

import (
	"flag"
	"fmt"
	"github.com/jordic/lti"
	"log"
	"net/http"
	"os"
)

// When run on multi-user linux, this server needs read access to all
// home directories with the "jupyter-" prefix.

// It also needs access to a TLS cert and key files, which are, as far
// as I (rhodesd) know, these files are not included as seperate files
// on the jupyterhub install by default, rather they need to be
// extracted from the acme.json that Traefik uses.

var (
	secret      = flag.String("secret", "", "Default secret for use during testing")
	consumer    = flag.String("consumer", "", "Def consumer")
	httpAddress = os.Getenv("WEBHOST")
)

type LtiHandler struct{}

func main() {
	flag.Parse()
	log.Printf("Lis %s, waiting POST request.", httpAddress)

	// open connection to sqlite3
	dbpath := os.Getenv("DB_FILE")
	err := initGlobalStore(dbpath) //"../answer-server/answers.db")
	if err != nil {
		log.Fatal(Err(err, "couldn't not acquire database"))
	}
	log.Printf("Loaded db: %s\n", dbpath)

	// https://gist.github.com/denji/12b3a568f092ab951456
	keyFile := "./key.pem"
	certFile := "./certificate.pem"
	h := LtiHandler{}
	log.Fatal(http.ListenAndServeTLS(httpAddress, certFile, keyFile, h))
	//log.Fatal(http.ListenAndServe(httpAddress, h))
}

func (h LtiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("got something...")

	if r.Method != "POST" {
		http.Error(w, "Only post", 500)
		return
	}

	p := lti.NewProvider(*secret, "https://"+httpAddress+"/")
	p.ConsumerKey = *consumer

	ok, err := p.IsValid(r)
	if err != nil {
		log.Printf("Invalid request %s", err)
		return
	}

	if ok {
		fmt.Println(p)
		userId := p.Get("user_id")
		labName := p.Get("custom_labname")
		// custom_staff_answers is a base64 encoded json string.
		staffAnswers64 := p.Get("custom_staff_answers")

		if userId == "student" {
			// handle the case in studio mode
			fmt.Fprintf(w, "Psst, the grader doesn't work in studio mode, please 'View live version'")
			return
		}

		// for now, shell out to the existing python rendering code
		// TODO what happens where there is no grade?

		proc, err := NewPyExec(userId, labName, staffAnswers64)

		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		} else {
			page, grade, err := proc.BuildPage()
			if err != nil {
				fmt.Fprintf(w, err.Error())
				return
			} else {
				err = sendGrade(p, grade)
				if err != nil {
					log.Println(Err(err, "Couldn't send grades back to openedx"))
					return
				}
				fmt.Fprintf(w, page)
				return
			}
		}
	} else {
		fmt.Fprintf(w, "Invalid request...")
	}
}
