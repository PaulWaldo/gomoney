package db

// func ConnectToDatabase() (*gorm.DB, error) {
// 	// In-memory sqlite if no database name is specified
// 	dsn := "file::memory:?cache=shared"
// 	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.AutoMigrate(&models.Account{})
// 	return db, nil
// }
