package dal

import (
	"github.com/TremblingV5/DouTok/applications/user/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
