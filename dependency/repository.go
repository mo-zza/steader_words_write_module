package dependency

import (
	firebase "firebase.google.com/go"
	"github.com/mo-zza/steader_words_write_module/domain/repository"
	"golang.org/x/text/cases"
	"google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1"
)

func NewWordsRepository(db *firebase.App) repository.WordsRepository {
	switch connection := db.(type) {
		case db.Firestore
	}
}