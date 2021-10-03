package config

import "os"

func FirebaseConfig() (string, string) {
	var serviceAccountKey, keyPath string

	if serviceAccountKey = os.Getenv("SERVICE_ACCOUNT_KEY"); len(serviceAccountKey) == 0 {
		serviceAccountKey = ""
	}
	if keyPath = os.Getenv("KEY_PATH"); len(keyPath) == 0 {
		keyPath = ""
	}
	return serviceAccountKey, keyPath
}
