package UserInterfacerface

import "forum-project/backend/internal/models"

type UserInterface interface {
	Register(*models.User) error
	createUser()
}
