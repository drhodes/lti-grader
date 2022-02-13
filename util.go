package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

func Err(err error, args ...interface{}) error {
	_, file, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf(" @ %s : %d\n", file, line)

	msg := loc + fmt.Sprint(args...)
	if err == nil {
		err = errors.New(msg)
	}
	return errors.New(err.Error() + "\n!!-- " + msg)
}

// https://stackoverflow.com/questions/10510691
// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func logif(err error) {
	if err != nil {
		log.Println(err)
	}
}

func generate_jupyterhub_userid(anonymous_student_id string) string {
	if len(anonymous_student_id) < 32 {
		anonymous_student_id += "--------------------------------"
	}
	anon_id := "jupyter-" + anonymous_student_id
	bs := sha256.Sum256([]byte(anon_id))
	userhash := hex.EncodeToString(bs[:])
	return fmt.Sprintf("%s-%s", anon_id[:26], userhash[:5])
}
