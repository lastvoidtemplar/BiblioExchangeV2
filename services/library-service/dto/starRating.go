package dto

type StarredDTO struct {
	Starred bool `json:"starred"`
}

func CreateStarDTO(starred bool) StarredDTO {
	return StarredDTO{
		Starred: starred,
	}
}
