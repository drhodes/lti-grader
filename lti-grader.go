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
	"errors"
	"flag"
	"fmt"
	"github.com/jordic/lti"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// This package allows to test the lib, acting as a webserver, and
// responding to a / endpoint... that should receive POST requests..

var (
	secret      = flag.String("secret", "", "Default secret for use during testing")
	consumer    = flag.String("consumer", "", "Def consumer")
	httpAddress = flag.String("https", "www.mathtech.org:5001", "Listen to")
)

type LtiHandler struct{}

func main() {
	flag.Parse()
	log.Printf("Lis %s, waiting POST request.", *httpAddress)

	// https://gist.github.com/denji/12b3a568f092ab951456
	keyFile := "./key.pem"
	certFile := "./certificate.pem"
	h := LtiHandler{}
	log.Fatal(http.ListenAndServeTLS("www.mathtech.org:5001", certFile, keyFile, h))
}

func (h LtiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only post", 500)
		return
	}

	p := lti.NewProvider(*secret, "https://www.mathtech.org:5001/")
	p.ConsumerKey = *consumer

	ok, err := p.IsValid(r)
	if err != nil {
		log.Printf("Invalid request %s", err)
		return
	}

	if ok {
		fmt.Fprintf(w, "<html> Request Ok <br><br> ")
		//data := fmt.Sprintf("User %s", p.Get("user_id"))
		username := fmt.Sprintf("%s", p.Get("lis_person_sourcedid"))
		labname := fmt.Sprintf("%s", p.Get("custom_component_display_name"))
		notebook, err := GetNotebookData(username, labname)
		fmtstr := "username: %s, labname: %s <br><br><br> notebook data \\o/ <br> %s </html>"
		var data = ""
		if err != nil {
			data = fmt.Sprintf(fmtstr, username, labname, err)
		} else {
			data = fmt.Sprintf(fmtstr, username, labname, notebook)
		}
		fmt.Fprint(w, data)

	} else {
		fmt.Fprintf(w, "Invalid request...")
	}
}

// <problem>
//     <customresponse cfn="check_function">
//         <script type="loncapa/python">
//             <![CDATA[ template-file(check_function.py) ]]>
//         </script>
//         <p>This is paragraph text displayed before the iframe.</p>
//         <jsinput
//             gradefn="JSInputDemo.getGrade"
//             get_statefn="JSInputDemo.getState"
//             set_statefn="JSInputDemo.setState"
//             initial_state='{"selectedChoice": "incorrect1", "availableChoices":
//             ["incorrect1", "correct", "incorrect2"]}'
//             width="600"
//             height="100"
//             html_file="https://files.edx.org/custom-js-example/jsinput_example.html"
//             title="Dropdown with Dynamic Text"
//             sop="false"/>
//     </customresponse>
// </problem>

func GetNotebookData(username string, labName string) (string, error) {
	fmtStr := "/home/jupyter-%s/%s.ipynb"
	filename := fmt.Sprintf(fmtStr, strings.ToLower(username), labName)
	bs, err := ioutil.ReadFile(filename)
	notebookData := string(bs)

	if err != nil {
		return "", errors.New("Couldn't find notebook data" + err.Error())
	} else {
		return notebookData, nil
	}

}
