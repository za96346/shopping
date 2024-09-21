package repository

import "backend/domain/entities"

type MemberRepository interface {
	GetMember(memberEntity *entities.Member) (*entities.Member, error)
}
