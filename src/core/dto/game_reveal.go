package dto

import "hexagonal/src/core/domain"

type BodyGameRevealCell struct {
	Row uint `json:"row"`
	Col uint `json:"col"`
}

type ResponseGameRevealCell domain.Game

func BuildResponseGameRevealCell(model domain.Game) ResponseGameRevealCell {
	return ResponseGameRevealCell(model)
}
