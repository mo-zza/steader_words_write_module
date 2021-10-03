package entity

type Words struct {
	Key string `firestore:"key, omitempty"`
	Value string `firestore:"value, omitempty"`
}