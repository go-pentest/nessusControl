package nessus

type ListPolicyResponse struct {
	Policies []struct {
		CreationDate         int    `json:"creation_date"`
		Description          string `json:"description"`
		ID                   int    `json:"id"`
		LastModificationDate int    `json:"last_modification_date"`
		Name                 string `json:"name"`
		NoTarget             string `json:"no_target"`
		Owner                string `json:"owner"`
		OwnerID              int    `json:"owner_id"`
		Shared               int    `json:"shared"`
		TemplateUUID         string `json:"template_uuid"`
		UserPermissions      int    `json:"user_permissions"`
		Visibility           string `json:"visibility"`
	} `json:"policies"`
}