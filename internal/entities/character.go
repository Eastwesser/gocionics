package entities

// Character represents a socionics character type
// @Description Socionics character type with traits and description
type Character struct {
	ID          int      `json:"id" example:"1"`
	Type        string   `json:"type" example:"Дон Кихот"`
	Description string   `json:"description" example:"Искатель, интуитивно-логический экстраверт"`
	Traits      []string `json:"traits" example:"изобретательный,энтузиаст,непредсказуемый"`
}
