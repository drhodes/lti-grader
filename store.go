package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// https://gorm.io/docs/

type SqlStore struct {
	DB     *gorm.DB
	DBPath string
}

type LabSubmission struct {
	gorm.Model
	EdxAnonId  string
	LabName    string
	LabAnswers string
}

// =============================================================================
// global mutable state, be careful!

var globalStore = SqlStore{}

// -----------------------------------------------------------------------------

func OpenSqlStore(path string) (SqlStore, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return SqlStore{}, Err(err, "Could not open the answer database")
	}
	db.AutoMigrate(&LabSubmission{})
	return SqlStore{db, path}, nil
}

func initGlobalStore(path string) error {
	store, err := OpenSqlStore(path)
	if err != nil {
		return Err(err, "could not initialize global store database")
	}
	globalStore = store
	return nil
}

func (store SqlStore) InsertAnswer(edxAnonId, labName, labAnswers string) error {
	store.DB.Create(&LabSubmission{
		EdxAnonId:  edxAnonId,
		LabName:    labName,
		LabAnswers: labAnswers,
	})
	return store.DB.Error
}

func (store SqlStore) GetAnswers(edxAnonId, labName string) (LabSubmission, error) {
	// get the latest answers submitted by user: edxAnonId with labName.
	var labSub LabSubmission
	err := store.DB.
		Where("edx_anon_id = ?", edxAnonId).
		Where("lab_name = ?", labName).
		Last(&labSub).Error

	log.Printf("getting answers for: user: %s, labName: %s\n", edxAnonId, labName)
	log.Println("This is what the database got: ", labSub)

	if err != nil {
		return labSub, err
	} else {
		return labSub, nil
	}
}
