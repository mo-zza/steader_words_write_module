package dependency

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"github.com/mo-zza/steader_words_write_module/config"
	"google.golang.org/api/option"
)


func NewConnection() (*firebase.App, error) {
	serviceAccountKey, keyPath := config.FirebaseConfig()

	opt := option.WithCredentialsFile(keyPath+serviceAccountKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}