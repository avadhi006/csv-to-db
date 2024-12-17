package models

//"gorm.io/gorm"

// Define the Record model
type Record struct {
	//gorm.Model
	SiteID                string
	FixletID              string
	Name                  string
	Crtiticality          string
	RelevantComputerCount string
}
