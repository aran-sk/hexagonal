package dto

import "hexagonal/src/core/domain"

type BodyGameCreate struct {
	Name  string `json:"name"`
	Size  uint   `json:"size"`
	Bombs uint   `json:"bombs"`
}

type ResponseGameCreate domain.Game

func BuildResponseGameCreate(model domain.Game) ResponseGameCreate {
	return ResponseGameCreate(model)
}
