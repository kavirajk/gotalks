package models

import "fmt"

//ALBUMMODELSTART OMIT
type Album struct {
	Model              // HL
	Title    string    `json:"title" gorm:"type:varchar(50);index"`
	Story    string    `json:"title"`
	Slug     string    `json:"slug"`
	Pictures []Picture `json:"-"`
}

//ALBUMMODELEND OMIT

func (a Album) String() string {
	return fmt.Sprintf("<%T %d: %s>", a, a.Id, a.Title)
}

//ALBUMDBSTART OMIT
func (a *Album) Save() error {
	if err := db.Save(a).Error; err != nil {
		return fmt.Errorf("album.save: %s", err)
	}
	return nil
}

func GetAlbum(id uint) (*Album, error) {
	var a Album
	if err := db.First(&a, "id=?", id).Error; err != nil {
		return nil, fmt.Errorf("album.get.id %d: %s", id, err)
	}
	return &a, nil
}

func DeleteAlbum(id uint) error {
	album, err := GetAlbum(id)
	if err != nil {
		return err
	}
	if err := db.Delete(album).Error; err != nil {
		return fmt.Errorf("album.delete.id %d: %s", id, err)
	}
	return nil
}

// ALBUMDBEND OMIT
func GetAllAlbums() ([]Album, error) {
	var albums []Album
	if err := db.Find(&albums).Error; err != nil {
		return nil, fmt.Errorf("album.get_all: %s", err)
	}
	return albums, nil
}
