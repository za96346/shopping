package service

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type MemberApp struct {
	MemberRepository repository.MemberRepository
}

type MemberAppInterface interface {
	GetMember(memberEntity *entities.Member) (*entities.Member, error)
	Login(memberEntity *entities.Member) (*entities.Member, error)
}

func (app *MemberApp) GetMember(memberEntity *entities.Member) (*entities.Member, error) {
	return (*app).MemberRepository.GetMember(memberEntity)
}

func (app  *MemberApp) Login(memberEntity *entities.Member) (*entities.Member, error) {
	return (*app).MemberRepository.GetMember(memberEntity)
}