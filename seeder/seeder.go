package seeder

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedUserSuperAdmin(contract *contract.Contract) error {
	log.Println("[INFO] Running seeder.....")
	db := contract.Common.DB

	roleId, err := SeedRole(db)
	if err != nil {
		return err
	}

	userName := os.Getenv("USER_NAME_SEEDER")
	if userName == "" {
		return fmt.Errorf("USER_NAME_SEEDER not set")
	}

	userEmail := os.Getenv("USER_EMAIL_SEEDER")
	if userEmail == "" {
		return fmt.Errorf("USER_EMAIL_SEEDER not set")
	}

	userPassword := os.Getenv("USER_PASSWORD_SEEDER")
	if userPassword == "" {
		return fmt.Errorf("USER_PASSWORD_SEEDER not set")
	}

	var count int64
	db.Model(&model.User{}).Where("email = ?", userEmail).Count(&count)
	if count > 0 {
		log.Println("Super admin already exists, skipping user seed.")
		return nil
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	admin := model.User{
		Name:     userName,
		Email:    userEmail,
		Password: string(hashedPwd),
		RoleID:   roleId,
	}

	if err := db.Create(&admin).Error; err != nil {
		return fmt.Errorf("failed to create super admin: %w", err)
	}

	log.Println("[INFO] Seeder completed.")
	return nil
}

func SeedRole(db *gorm.DB) (uuid.UUID, error) {
	roleName := os.Getenv("ROLE_NAME_SEEDER")
	if roleName == "" {
		return uuid.UUID{}, fmt.Errorf("ROLE_NAME_SEEDER not set")
	}

	var role model.Role
	err := db.Where("name = ?", roleName).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRole := model.Role{
				Name: roleName,
			}
			if err := db.Create(&newRole).Error; err != nil {
				return uuid.UUID{}, fmt.Errorf("failed to create super admin role: %w", err)
			}
			role = newRole // assign ke variabel role untuk dipakai nanti
			log.Println("Role Super Admin created successfully.")
		} else {
			return uuid.UUID{}, fmt.Errorf("failed to check role: %w", err)
		}
	} else {
		log.Println("Role Super Admin already exists, skipping seed role.")
	}

	// add new default role
	err = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&model.Role{
		Name: constants.DEFAULT_ROLE,
	}).Error
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to create default role: %w", err)
	}

	return role.ID, nil
}
