package usecase

import "sample/domain"

type SampleInteractor struct {
	SampleRepository SampleRepository
}

func (interactor *SampleInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.SampleRepository.FindById(id)
	return
}

func (interactor *SampleInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.SampleRepository.FindAll()
	return
}

func (interactor *SampleInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.SampleRepository.Store(u)
	return
}

func (interactor *SampleInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.SampleRepository.Update(u)
	return
}

func (interactor *SampleInteractor) DeleteById(u domain.User) (err error) {
	err = interactor.SampleRepository.DeleteById(u)
	return
}