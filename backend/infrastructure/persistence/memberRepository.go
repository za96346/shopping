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

	searchQuery := r.db.
	Debug().
	Table(r.tableName)

	if memberEntity.Account != "" {
		searchQuery = searchQuery.Where("account = ?", memberEntity.Account)
	}

	if memberEntity.Password != "" {
		searchQuery = searchQuery.Where("password = ?", memberEntity.Password)
	}

	if memberEntity.ID != 0 {
		searchQuery = searchQuery.Where("id = ?", memberEntity.ID)
	}

	err := searchQuery.
		Find(&member).
		Error

	return &member, err
}
