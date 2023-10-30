package models

func (s *Conn) AutoMigrate() error {
	//if s.db.Migrator().HasTable(&User{}) {
	//	return nil
	//}


	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err := s.db.Migrator().AutoMigrate(&User{},&Company{},&Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	s.db.AutoMigrate(&Company{}, &Job{})

	// Add foreign key constraint

	return nil
}
