package core

type Note struct {
	Id     uint32
	Decree string
	Priest int
}

func NewNote(id uint32, decree string, priest int) *Note {
	return &Note{
		Id:     id,
		Decree: decree,
		Priest: priest,
	}
}

func InitNote() (*[]Note, error) {
	return FindAllNote()
}

func ContainsNote(notes []Note, decree string) bool {
	for _, v := range notes {
		if v.Decree == decree {
			return true
		}
	}
	return false
}
