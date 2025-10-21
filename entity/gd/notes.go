package gdentity

type NoteItemType int32

const (
	NoteItemTypeCustomer NoteItemType = 1
)

type NoteIsTop int32

const (
	NotIsTopNote NoteIsTop = 0
	IsTopNote    NoteIsTop = 1
)
