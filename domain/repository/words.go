package repository

import "github.com/mo-zza/steader_words_write_module/domain/entity"

type WordsRepository interface {
	Save(int, *entity.Words) error
}