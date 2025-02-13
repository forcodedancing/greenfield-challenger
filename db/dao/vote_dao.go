package dao

import (
	"github.com/bnb-chain/greenfield-challenger/db/model"
	"gorm.io/gorm"
)

type VoteDao struct {
	DB *gorm.DB
}

func NewVoteDao(db *gorm.DB) *VoteDao {
	return &VoteDao{
		DB: db,
	}
}

func (d *VoteDao) SaveVote(vote *model.Vote) error {
	err := d.DB.Create(vote).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *VoteDao) GetVotesByChallengeId(challengeId uint64) ([]*model.Vote, error) {
	votes := make([]*model.Vote, 0)
	err := d.DB.
		Where("challenge_id = ?", challengeId).
		Find(&votes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return votes, nil
}

func (d *VoteDao) IsVoteExists(challengeId uint64, pubKey string) (bool, error) {
	exists := false
	if err := d.DB.Raw(
		"SELECT EXISTS(SELECT id FROM votes WHERE challenge_id = ? and pub_key = ?)",
		challengeId, pubKey).Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}
