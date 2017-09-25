package core

type Note struct {
	Id     uint32
	Decree string
	Priest int
}

func newNote(id uint32, decree string, priest int) *Note {
	return &Note{
		Id:     id,
		Decree: decree,
		Priest: priest,
	}
}

func initNote() ([]Note, error) {
	return findAllNote()
}

func containsNote(notes []Note, decree string) bool {
	for _, v := range notes {
		if v.Decree == decree {
			return true
		}
	}
	return false
}

func addNote(notes []Note, note Note) ([]Note, error) {
	err := insertNote(note)
	if err != nil {
		return notes, err
	}
	if cap(notes) == 0 {
		notes = make([]Note, 5)
	}
	var c_notes []Note
	if len(notes) == cap(notes) {
		c_notes = make([]Note, len(notes)*2)
		copy(c_notes, notes)
	}
	c_notes[len(notes)] = note
	return c_notes, nil
}
