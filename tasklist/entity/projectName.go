package entity

type ProjectName struct {
	value string
}

func NewProjectName(value string) *ProjectName {
	return &ProjectName{value: value}
}

func (p *ProjectName) ToString() string {
	return p.value
}
