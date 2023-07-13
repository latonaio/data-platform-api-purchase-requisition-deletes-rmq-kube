package requests

type Item struct {
	PurchaseRequisition      	int   `json:"PurchaseRequisition"`
	PurchaseRequisitionItem     int   `json:"PurchaseRequisitionItem"`
	IsMarkedForDeletion 	 	*bool `json:"IsMarkedForDeletion"`
}
