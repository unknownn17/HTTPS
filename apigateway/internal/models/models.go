package models

type CreateItemRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Amount   int32  `json:"amount"`
}

type GeneralItem struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Amount   int32  `json:"amount"`
}

type GetItemRequest struct {
	ID int32 `json:"id"`
}

type GetItemsRequest struct{}

type DeleteResponse struct {
	Message string `json:"message"`
}

type Error struct {
	Message string `json:"message"`
}
