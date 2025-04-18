package config

import 	"github.com/Swapnilgupta8585/CollabDocs/internal/database"

type ApiConfig struct{
	Db *database.Queries
	SecretToken string
}
