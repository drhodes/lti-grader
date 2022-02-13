package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	// "io/ioutil"
	// "strings"
	"os/exec"
)

// run the python grader program
type PyExec struct {
	EdxAnonId    string
	LabName      string
	StaffAnswers string
}

func NewPyExec(userId, labName, staffAnswers64 string) (PyExec, error) {
	edxAnonId := generate_jupyterhub_userid(userId)

	staffAnswers, err := b64.StdEncoding.DecodeString(staffAnswers64)
	if err != nil {
		msg := `
Could not decode staffs_answers from lti request for lab:%s,
it is likely there is a problem with the formatting in the lti_consumer
`
		return PyExec{}, Err(err, fmt.Sprintf(msg, labName))
	}

	return PyExec{edxAnonId, labName, string(staffAnswers)}, nil
}

func (ex PyExec) getUserAnswer() (string, error) {
	answers, err := globalStore.GetAnswers(ex.EdxAnonId, ex.LabName)
	if err != nil {
		log.Println(Err(err, "Could not find your answers, make sure you have submitted the lab"))
		return "{}", nil
	}
	return answers.LabAnswers, nil
}

func (ex PyExec) BuildPage() (string, string, error) {
	app := "./grade_answers.py"

	arg0 := "--student-answers"
	arg1, err := ex.getUserAnswer()

	if err != nil {
		return "", "", Err(err, "Couldn't get user answers from the server-answer database")
	}

	arg2 := "--staff-answers"
	arg3 := ex.StaffAnswers

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return "", "", Err(err, "answer-server could not build page because could not run python program")
	}

	// stdout has json encoded result of the python script.
	// if output has errors, then the form is {"Ok": False, "Error": <<msg>>, "Html": "", "Grade": 0}
	// if the output is ok,  then the form is {"Ok": False, "Error": "", "Html": <<webpage>>, "Grade": int}

	out := struct {
		Ok    bool
		Grade float64
		Error string
		Html  string
	}{}

	err = json.Unmarshal([]byte(stdout), &out)
	if err != nil {
		fmt.Println(err.Error())
		return "", "", Err(err, "Couldn't decode json output to generate webpage for lab answers")
	}

	if !out.Ok {
		return "", "", Err(nil, "Couldn't build answer page because: "+out.Error)
	}

	// Print the output
	gradeStr := fmt.Sprintf("%f", out.Grade)
	return string(out.Html), gradeStr, nil

}
