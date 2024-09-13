package models

type Patient struct {
    PatientName   string `json:"patient_name"`
    Age           uint   `json:"age"`
    Gender        string `json:"gender"`
    Disease       string `json:"disease"`
    ContactNumber string `json:"contact_number"`
    BloodGroup    string `json:"blood_group"`
    IPFSHash      string `json:"ipfs_hash"`
}

