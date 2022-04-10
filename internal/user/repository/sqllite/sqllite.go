package sqllite

import (
	"context"
	"gorm.io/gorm"
	"users-api-server/internal/user/model"
)

type SqlLite struct {
	db *gorm.DB
}

func NewSqlLiteRepo(db *gorm.DB) SqlLite {
	return SqlLite{
		db: db,
	}
}

func (sl SqlLite) Migrate() error {
	return sl.db.AutoMigrate(&model.Position{}, &model.User{})
}

func (sl SqlLite) CreateUser(ctx context.Context, user model.User) (string, error) {
	if err := sl.db.WithContext(ctx).Create(&user).Error; err != nil {
		return "", err
	}

	return user.ID, nil
}

func (sl SqlLite) UpdateUserByID(ctx context.Context, id string, user model.User) (model.User, error) {
	var foundUser model.User
	if err := sl.db.WithContext(ctx).
		First(&foundUser, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	foundUser.FirstName = user.FirstName
	foundUser.LastName = user.LastName
	foundUser.Age = user.Age
	foundUser.Email = user.Email
	foundUser.Position = user.Position

	if err := sl.db.WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).Updates(&foundUser).Error; err != nil {
		return model.User{}, err
	}

	return foundUser, nil
}

func (sl SqlLite) GetUserByID(ctx context.Context, id string) (model.User, error) {
	var foundUser model.User
	if err := sl.db.WithContext(ctx).Preload("Position").
		First(&foundUser, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	return foundUser, nil
}

func (sl SqlLite) SetPartially(ctx context.Context, id string, values map[string]interface{}) (model.User, error) {
	var updatedPartially model.User
	tx := sl.db.WithContext(ctx)
	if err := tx.Model(&model.User{ID: id}).
		Updates(values).Error; err != nil {
		return model.User{}, err
	}

	if err := tx.Preload("Position").First(&updatedPartially, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	return updatedPartially, nil
}

func (sl SqlLite) GetAll(
	ctx context.Context,
	pageNumber int,
	countInt int,
	sort string,
	filter map[string]interface{},
) ([]model.User, int64, error) {
	//calculate the offset
	offset := (pageNumber * countInt) - countInt
	if offset < 0 {
		offset = 0
	}

	//build the query
	queryBuilder := sl.db.WithContext(ctx).Limit(countInt).Offset(offset).Order(sort)
	//check if there is any filter
	if filter != nil && len(filter) != 0 {
		queryBuilder = queryBuilder.Where(filter)
	}

	var users []model.User
	if err := queryBuilder.Preload("Position").Find(&users).Error; err != nil {
		return users, 0, err
	}

	var totalRecords int64
	if err := sl.db.Model(&model.User{}).Count(&totalRecords).Error; err != nil {
		return users, 0, err
	}

	return users, totalRecords, nil
}
