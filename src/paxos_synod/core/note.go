package core

type Note struct {
	Id     uint
	Decree string
	Priest int
}

func NewNote(id uint, decree string, priest int) *Note {
	return &Note{
		Id:     id,
		Decree: decree,
		Priest: priest,
	}
}

func InitNote() (*[]Note, error) {
	return FindAllNote()
}
