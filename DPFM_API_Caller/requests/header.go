package requests

type Header struct {
	PurchaseRequisition      int   `json:"PurchaseRequisition"`
	IsMarkedForDeletion 	 *bool `json:"IsMarkedForDeletion"`
}
