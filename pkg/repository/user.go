package repository

import (
	"database/sql"
	"errors"
	"fmt"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.IUserRepo {
	return &userRepository{DB: DB}
}

//user Repository

func (d *userRepository) CreateUser(userDetails *requestmodel.UserDetails) {
	query := "INSERT INTO users (name, email, phone, password) VALUES($1, $2, $3, $4)"
	d.DB.Exec(query, userDetails.Name, userDetails.Email, userDetails.Phone, userDetails.Password)
}

func (d *userRepository) IsUserExist(phone string) int {
	var userCount int

	query := "SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2"
	err := d.DB.Raw(query, phone, "delete").Row().Scan(&userCount)
	if err != nil {
		fmt.Println("Error for user exist, using same phone in signup")
	}
	return userCount
}

func (d *userRepository) ChangeUserStatusActive(phone string) error {
	fmt.Println(phone)
	query := "UPDATE users SET status = 'active' WHERE phone = ?"
	result := d.DB.Exec(query, phone)
	// count:=result.RowsAffected

	if result.Error != nil {
		return errors.New("no user Exist , phone number is wrong")
	} else {
		return nil
	}
}

func (d *userRepository) FetchUserID(phone string) (string, error) {
	var userID string

	query := "SELECT id FROM users WHERE phone=? AND status='active'"
	data := d.DB.Raw(query, phone).Row()

	if err := data.Scan(&userID); err != nil {
		return "", errors.New("fetching user id cause error")
	}
	return userID, nil
}

func (d *userRepository) FetchPasswordUsingPhone(phone string) (string, error) {
	var password string

	query := "SELECT password FROM users WHERE phone=? AND status='active'"
	row := d.DB.Raw(query, phone).Row()

	if row == nil {
		return "", errors.New("no user exist or you are blocked by admin")
	}

	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user does not exist or user get blocked")
		}
		return "", fmt.Errorf("error scanning row: %s", err)
	}
	return password, nil
}

func (d *userRepository) AllUsers(offSet int, limit int) (*[]responsemodel.UserDetails, error) {
	var users []responsemodel.UserDetails

	query := "SELECT * FROM users ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&users).Error
	if err != nil {
		return nil, errors.New("can't get user data from db")
	}

	return &users, nil
}

func (d *userRepository) UserCount(ch chan int) {
	var count int

	query := "SELECT COUNT(phone) FROM users WHERE status!='delete'"
	d.DB.Raw(query).Scan(&count)
	ch <- count
}

func (d *userRepository) BlockUser(id string) error {
	query := "UPDATE users SET status = 'block' WHERE id=? "
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("block user process , is not satisfied")
	}
	count := err.RowsAffected
	if count <= 0 {
		return errors.New("no user exist by id ")
	}
	return nil
}

func (d *userRepository) UnblockUser(id string) error {
	query := "UPDATE users SET status = 'active' WHERE id=?"
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("active user process , is not satisfied")
	}

	if err.RowsAffected <= 0 {
		return errors.New("no user exist by id ")
	}
	return nil
}
