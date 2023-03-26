package users

import (
	"database/sql"
	"errors"

	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/humgal/art-server/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string  `json:"id"`
	Realname     string  `json:"realname"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Avatar       *string `json:"avatar"`
	Phone        *string `json:"phone"`
	Company      *string `json:"company"`
	Email        *string `json:"email"`
	Bio          *string `json:"bio"`
	Img          *string `json:"img"`
	JoinDate     *string `json:"joinDate"`
	VerifyType   *int    `json:"verifyType"`
	VerifyName   *string `json:"verifyName"`
	IsCreator    *bool   `json:"isCreator"`
	FollowerNum  *int    `json:"followerNum"`
	FollowingNum *int    `json:"followingNum"`
	Links        *string `json:"links"`
}

func (user *User) Create() (bool, error) {
	hashedPassword, _ := HashPassword(user.Password)

	_, err := db.DB.Exec("INSERT INTO User(Username,Password,Email,Phone) VALUES(?,?,?,?)", user.Username, hashedPassword, user.Email, user.Phone)
	if err != nil {
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			switch errMySQL.Number {
			case 1062:
				return false, errors.New("用户已存在")
			default:
				return false, err
			}
		}
		return false, err
	}
	return true, nil

}

func (user *User) Authenticate() bool {
	statement, err := db.DB.Prepare("select Password from User WHERE Username = ?")
	if err != nil {
		log.Println(err)
	}
	row := statement.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Println(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

// GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	statement, err := db.DB.Prepare("select ID from Users WHERE Username = ?")
	if err != nil {
		log.Println(err)
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
