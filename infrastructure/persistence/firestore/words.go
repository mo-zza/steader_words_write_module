package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/mo-zza/steader_words_write_module/domain/entity"
)

type wordsRepository struct {
	app *firebase.App
}

func (w *wordsRepository) Save(index int, word *entity.Words) error {
	ctx := context.Background()
	client, err := w.app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	doc := make(map[int]interface{})
	doc[index] = map[string]interface{}{
		"key":   word.Key,
		"value": word.Value,
	}
	_, firestoreErr := client.Collection("words").Doc("highSchool").Set(ctx, doc, firestore.MergeAll)

	if firestoreErr != nil {
		fmt.Println(firestoreErr)
	}
	return nil
}
