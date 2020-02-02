package seed

import (
	"log"

	"github.com/LuD1161/posts_api/fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Asim Shere",
		Email:    "asimshere@cryptosports.com",
		Password: "asim_password",
	},
	models.User{
		Nickname: "Moodit Khand",
		Email:    "mooditkhand@cryptosports.com",
		Password: "moodit_password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

// Load test data in DB
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot DROP table : %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table : %v", err)
	}
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
