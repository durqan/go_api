package contact

type CreateContactRequest struct {
	Type  string `json:"type" binding:"required,oneof=email phone telegram whatsapp"`
	Value string `json:"value" binding:"required,min=1,max=255"`
}
