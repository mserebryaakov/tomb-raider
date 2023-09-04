package model

type ITypeSchema interface {
	GetSearch()
}

type TypeSchema struct {
	Type     string
	Multiple bool
	Deleted  bool
}

func (ts *TypeSchema) GetSearch() {

}
