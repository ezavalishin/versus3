package models

import (
	"github.com/ezavalishin/versus3/pkg/utils"
	"github.com/jinzhu/gorm"
	vkapi "github.com/leidruid/golang-vk-api"
)

type User struct {
	BaseModel
	VkUserId  int `gorm:"NOT NULL"`
	FirstName *string
	LastName  *string
	Avatar    *string
}

func (u *User) AfterCreate(scope *gorm.Scope) (err error) {

	go fillUserFromVk(u, scope)
	go fillUserFriends(u, scope)

	return
}

func fillUserFriends(u *User, scope *gorm.Scope) {

	db := scope.DB()

	client, err := vkapi.NewVKClientWithToken(utils.MustGet("VK_APP_SERVICE_KEY"), &vkapi.TokenOptions{
		ServiceToken:    true,
		TokenLanguage:   "ru",
		ValidateOnStart: true,
	})

	if err != nil {
		return
	}

	_, friends, err := client.FriendsGet(u.VkUserId, 100)

	var friendVkUserIds []int

	for _, vkUser := range friends {
		friendVkUserIds = append(friendVkUserIds, vkUser.UID)
	}

	var friendModels []*User

	db.Where("vk_user_id in (?)", friendVkUserIds).Find(&friendModels)

	for _, friendModel := range friendModels {

		friendUserModel := Friend{
			UserOneId: u.ID,
			UserTwoId: friendModel.ID,
		}

		friendUserModelReversed := Friend{
			UserOneId: friendModel.ID,
			UserTwoId: u.ID,
		}

		db.New().FirstOrCreate(&friendUserModel, friendUserModel)
		db.New().FirstOrCreate(&friendUserModelReversed, friendUserModelReversed)
	}
}

func fillUserFromVk(u *User, scope *gorm.Scope) {
	client, err := vkapi.NewVKClientWithToken(utils.MustGet("VK_APP_SERVICE_KEY"), &vkapi.TokenOptions{
		ServiceToken:    true,
		TokenLanguage:   "ru",
		ValidateOnStart: true,
	})

	if err != nil {
		return
	}

	users, err := client.UsersGet([]int{u.VkUserId})

	if err != nil {
		return
	}

	vkUser := users[0]

	u.FirstName = &vkUser.FirstName
	u.LastName = &vkUser.LastName
	u.Avatar = &vkUser.PhotoMedium

	scope.DB().Save(&u)
}
