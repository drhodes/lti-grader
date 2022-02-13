package main

import (
	"log"
	"os"
	"os/exec"
)

// run the python grader program
type PyExecSendGrade struct {
	LisOutcomeServiceUrl string
	LisResultSourcedId   string
	Grade                string
}

func (ex PyExecSendGrade) SendGrade() error {
	app := "./send_grade.py"
	arg0 := "--consumer"
	arg1 := "--secret"
	arg2 := "--lis-outcome-service-url"
	arg3 := "--lis-result-sourcedid"
	arg4 := "--grade"

	cmd := exec.Command(app,
		arg0, os.Getenv("GRADER_CONSUMER"),
		arg1, os.Getenv("GRADER_SECRET"),
		arg2, ex.LisOutcomeServiceUrl,
		arg3, ex.LisResultSourcedId,
		arg4, ex.Grade)

	stdout, err := cmd.Output()

	// TODO clean this up.
	if err != nil {
		log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
		log.Println(err)
		log.Println(stdout)
		buf := []byte{}
		cmd.Stderr.Write(buf)
		log.Println(string(buf))
		log.Println("consumer    ", os.Getenv("GRADER_CONSUMER"))
		log.Println("secret      ", os.Getenv("GRADER_SECRET"))
		log.Println("service_url ", ex.LisOutcomeServiceUrl)
		log.Println("sourcdid    ", ex.LisResultSourcedId)
		log.Println("grade       ", ex.Grade)
		return Err(err, "could not send grade because could not run send_grade.py python script")
	}
	return nil
}
