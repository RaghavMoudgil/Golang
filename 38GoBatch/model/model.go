package model

type StudentData struct {
	StudName   string `order:"0" header:"studName"`
	StudRollno int64  `order:"1" header:"studRollno"`
	StudBranch string `order:"2" header:"studBranch"`
}
