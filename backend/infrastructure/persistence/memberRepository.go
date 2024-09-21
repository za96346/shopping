package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type MemberRepo struct {
	db        *gorm.DB
	tableName string
}

func NewMemberRepository(db *gorm.DB) *MemberRepo {
	return &MemberRepo{db, "member"}
}

var _ repository.MemberRepository = &MemberRepo{}

func (r *MemberRepo) GetMember(memberEntity *entities.Member) (*entities.Member, error) {
	var member entities.Member

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("id = ?", memberEntity.ID).
		Find(&member).
		Error

	return &member, err
}
