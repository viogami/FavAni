package repos

import (
	"DigBGM/database"
	"errors"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserRepository(db *gorm.DB, rdb *redis.Client) UserRepository {
	return &userRepository{
		db:  db,
		rdb: rdb,
	}
}

// 登陆方法
func (u *userRepository) Login(username, password string) (*database.User, error) {
	var user database.User
	result := u.db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// 注册方法
func (u *userRepository) Register(newUser database.User) error {
	var existingUser database.User
	result := u.db.Where("username = ?", newUser.Username).First(&existingUser)
	if result.Error == nil {
		// 用户名已存在，返回错误
		return errors.New("username already exists")
	}
	// 在数据库中插入新用户
	result = u.db.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 获取全部用户
func (u *userRepository) List() (database.Users, error) {
	var users database.Users
	result := u.db.Find(&users)
	if result.Error != nil {
		// 获取用户列表失败，返回错误
		return nil, result.Error
	}

	return users, nil
}

func (u *userRepository) Create(user *database.User) (*database.User, error) {
	// TODO
	return user, nil
}

func (u *userRepository) Update(user *database.User) (*database.User, error) {
	// TODO
	return user, nil
}

// 删除用户
func (u *userRepository) Delete(user database.User) error {
	// 在数据库中查找对应用户名的用户
	result := u.db.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		// 用户不存在，返回错误
		return result.Error
	}

	// 执行删除操作
	result = u.db.Delete(&user)
	if result.Error != nil {
		// 删除失败，返回错误
		return result.Error
	}
	return nil
}

func (u *userRepository) GetUserByName(name string) (*database.User, error) {
	user := new(database.User)
	// TODO
	return user, nil
}
