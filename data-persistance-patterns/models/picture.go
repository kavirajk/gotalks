package models

import "fmt"

// PICMODELSTART OMIT
type Picture struct {
	Model          // HL
	Caption string `json:"caption" gorm:"type:varchar(144)"`
	AlbumId uint   `json:"album_id" gorm:"index"`
	Album   *Album `json:"-"`
	Path    string `json:"path"`
}

// PICMODELEND OMIT
func (p Picture) String() string {
	return fmt.Sprintf("<%T %d: %s>", p, p.Id, p.Caption)
}

// PICDBSTART OMIT
func (p *Picture) Save() error {
	if err := db.Save(p).Error; err != nil {
		return fmt.Errorf("picture.save: %s", err)
	}
	return nil
}

func GetPicture(id uint) (*Picture, error) {
	var p Picture
	if err := db.First(&p, "id=?", id).Error; err != nil {
		return nil, fmt.Errorf("picture.get: %s", err)
	}
	return &p, nil
}

func GetAllPictures() ([]Picture, error) {
	var pics []Picture
	if err := db.Find(&pics).Error; err != nil {
		return nil, fmt.Errorf("picture.get_all: %s", err)
	}
	return pics, nil
}

// PICDBEND OMIT

func DeletePicture(id uint) error {
	pic, err := GetPicture(id)
	if err != nil {
		return err
	}
	if err := db.Delete(pic).Error; err != nil {
		return fmt.Errorf("picture.delete: %s", err)
	}
	return nil
}
