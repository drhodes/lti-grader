package main

import (
	// "bytes"
	// "fmt"
	"github.com/jordic/lti"
	//"io"
	// "io/ioutil"
	// "log"
	// "net/http"
	//"net/url"
	//"os"
	//"strings"
	//"text/template"
)

// https://www.imsglobal.org/specs/ltiv1p1/implementation-guide#toc-26

// The sourcedId element is the value from the lis_result_sourcedid
// parameter for a particular user_id / resource_link_id combination.
// The TP records these values as they are sent on launches and can
// then later make services calls providing the sourcedId as way to
// pick the particular cell in the TC grade book.

// For this particular service, all of the values for textString are
// decimal values numeric in the range 0.0 - 1.0.  Regardless of the
// language of the TP or TC user interface, the number format is to
// use a period as the decimal point.  Regardless of the language of
// the TP or TC user interface, the language field in the service call
// is to be “en” indicating the format of the number. While the TP is
// required to include “en” as the language, the TC will likely ignore
// the language field in this request and always assume that the
// number is formatted using “en” formatting.

// The replaceResultResponse indicates the success/failure of the
// operation in the header area of the response and as such the body
// area is empty.

// The TC must check the incoming grade for validity and must fail
// when a grade is outside the range 0.0-1.0 or if the grade is not a
// valid number.  The TC must respond to these replaceResult
// operations with a imsx_codeMajor of "failure".

// implmenetation here
// https://github.com/pylti/lti/blob/master/src/lti/outcome_request.py

type ReplaceResults struct {
	MsgId     string
	SourcedId string
	Grade     float64
}

func sendGrade(p *lti.Provider, grade string) error {
	lisResultSourcedId := p.Get("lis_result_sourcedid")
	if lisResultSourcedId == "" {
		return Err(nil, "Could not find lis_result_sourcedid in LTI request")
	}
	lisOutcomeServiceUrl := p.Get("lis_outcome_service_url")
	if lisOutcomeServiceUrl == "" {
		return Err(nil, "Could not find lis_outcome_service_url in LTI request")
	}

	sender := PyExecSendGrade{lisOutcomeServiceUrl, lisResultSourcedId, grade}
	return sender.SendGrade()
}

// // req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// xml, err := ioutil.ReadFile("./static/xml/replace-score.xml")
// if err != nil {
// 	return Err(err, "Couldn't open the replace-score.xml template for building the grade request")
// }

// tmpl, err := template.New("replace-results").Parse(string(xml))
// if err != nil {
// 	return Err(err, "Couldn't parser the replace-score.xml template for building the grade request")
// }

// buf := new(bytes.Buffer)
// rr := ReplaceResults{"12341234", lisResultSourcedid, grade}
// err = tmpl.Execute(buf, rr)
// if err != nil {
// 	return Err(err, "Could not build the grade request")
// }
// fmt.Println(buf)
// //req.Body = io.NopCloser(buf)

// req, err := http.NewRequest("POST", lisOutcomeServiceUrl, buf)
// if err != nil {
// 	return Err(err, "Couldn't build request to send to answer server")
// }

// client := &http.Client{}
// resp, err := client.Do(req)
// if err != nil {
// 	return Err(err, "Had trouble sending request to the answer server")
// }

// if resp.StatusCode != 200 {
// 	log.Println(resp)
// 	return Err(nil, fmt.Sprintf("Got a bad response code: %d, %s", resp.StatusCode, resp))
// } else {
// 	log.Println(resp)
// }
// return nil

// https://stackoverflow.com/questions/42420210
// https://www.imsglobal.org/specs/ltiv1p1/implementation-guide#toc-37

// https://edge.edx.org/courses/course-v1:MITx+18.02+T12022/xblock/block-v1:MITx+18.02+T12022+type@lti_consumer+block@638391ca6cce44869933580948a33c88/handler_noauth/outcome_service_handler

// course-v1%3AMITx%2B18.02%2BT12022:edge.edx.org-638391ca6cce44869933580948a33c88:35dd7e9124c8847ec5291a45e555f18c
