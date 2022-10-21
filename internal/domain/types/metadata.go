package types

type Metadata struct {
	TotalItems   int64 `json:"totalItems"`
	ItemCount    int64 `json:"itemCount"`
	ItemsPerPage int64 `json:"itemsPerPage"`
	TotalPages   int64 `json:"totalPages"`
	CurrentPage  int64 `json:"currentPage"`
}
