package interactor

import (
	"github.com/mo-zza/steader_words_write_module/domain/entity"
	"github.com/mo-zza/steader_words_write_module/domain/repository"
)

type WordsInteractor interface {
	Save(int, words *entity.Words) error
}

type wordsInteractor struct {
	wordsRepository repository.WordsRepository
}

func (wi *wordsInteractor) Save(index int, words *entity.Words) error {
	err := wi.wordsRepository.Save(index, words);

	if err != nil {
		return err
	}
	return nil
}