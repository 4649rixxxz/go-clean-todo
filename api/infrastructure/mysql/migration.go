package mysql

func RunMigration() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Todo{})
}
