package models

import (
	"github.com/jinzhu/gorm"
)

type Pair struct {
	BaseModel
	Round         Round
	RoundId       int
	UnitOne       Unit
	UnitOneId     int
	UnitTwo       Unit
	UnitTwoId     int
	UserPick      *Unit
	UserPickId    *int
	UserId        int
	User          User
	UnitOnePicked Picked `gorm:"-"`
	UnitTwoPicked Picked `gorm:"-"`
}

func (pair *Pair) LoadPicked(db *gorm.DB) {

	var pairs []*Pair
	var friendPairs []*Pair
	var friends []*Friend
	var unitOnePickedCount, unitTwoPickedCount, unitOneFriendPickedCount, unitTwoFriendPickedCount int
	var friendUserIds []int

	var unitOnePickedUsers []*User
	var unitOnePickedUserIds []int
	var unitTwoPickedUsers []*User
	var unitTwoPickedUserIds []int

	db.Where("user_one_id = ?", pair.UserId).Find(&friends)

	for _, friend := range friends {
		friendUserIds = append(friendUserIds, friend.UserTwoId)
	}

	db.Where("((unit_one_id = ? and unit_two_id = ?) or (unit_one_id = ? and unit_two_id = ?)) and user_pick_id IS NOT NULL",
		pair.UnitOneId, pair.UnitTwoId, pair.UnitTwoId, pair.UnitOneId).
		Find(&pairs)

	db.Where("((unit_one_id = ? and unit_two_id = ?) or (unit_one_id = ? and unit_two_id = ?)) and user_id in (?) and user_pick_id IS NOT NULL",
		pair.UnitOneId, pair.UnitTwoId, pair.UnitTwoId, pair.UnitOneId, friendUserIds).
		Find(&friendPairs)

	for _, p := range friendPairs {
		if *p.UserPickId == pair.UnitOneId {
			unitOneFriendPickedCount++

			if len(unitOnePickedUserIds) < 3 {
				unitOnePickedUserIds = append(unitOnePickedUserIds, p.UserId)
			}
		}

		if *p.UserPickId == pair.UnitTwoId {
			unitTwoFriendPickedCount++

			if len(unitTwoPickedUserIds) < 3 {
				unitTwoPickedUserIds = append(unitTwoPickedUserIds, p.UserId)
			}
		}
	}

	for _, p := range pairs {
		if *p.UserPickId == pair.UnitOneId {
			unitOnePickedCount++

			if len(unitOnePickedUserIds) < 3 {
				unitOnePickedUserIds = append(unitOnePickedUserIds, p.UserId)
			}
		}

		if *p.UserPickId == pair.UnitTwoId {
			unitTwoPickedCount++

			if len(unitTwoPickedUserIds) < 3 {
				unitTwoPickedUserIds = append(unitTwoPickedUserIds, p.UserId)
			}
		}
	}

	db.Where("id in (?)", unitOnePickedUserIds).Find(&unitOnePickedUsers)
	db.Where("id in (?)", unitTwoPickedUserIds).Find(&unitTwoPickedUsers)

	pair.UnitOnePicked = Picked{
		Count:        unitOnePickedCount,
		FriendsCount: unitOneFriendPickedCount,
		Users:        unitOnePickedUsers,
	}

	pair.UnitTwoPicked = Picked{
		Count:        unitTwoPickedCount,
		FriendsCount: unitTwoFriendPickedCount,
		Users:        unitTwoPickedUsers,
	}
}
