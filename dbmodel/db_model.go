package dbmodel

type Vehicle struct {
	ID                   uint   `gorm:"primaryKey;autoIncrement"`
	CargoCapacity        string `gorm:"type:varchar(255)"`
	Consumables          string `gorm:"type:varchar(255)"`
	CostInCredits        string `gorm:"type:varchar(255)"`
	Crew                 string `gorm:"type:varchar(255)"`
	Length               string `gorm:"type:varchar(255)"`
	Manufacturer         string `gorm:"type:varchar(255)"`
	MaxAtmospheringSpeed string `gorm:"type:varchar(255)"`
	Model                string `gorm:"type:varchar(255)"`
	Name                 string `gorm:"type:varchar(255)"`
	Passengers           string `gorm:"type:varchar(255)"`
	VehicleClass         string `gorm:"type:varchar(255)"`
	URL                  string `gorm:"type:varchar(255)"`

	// Pilots []model.Character `gorm:"many2many:vehicle_pilots;"`
	// Films  []Film  `gorm:"many2many:vehicle_films;"`
}
