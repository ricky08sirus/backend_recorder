package api

import (
    "encoding/json"
    "net/http"
    "go-client/blockchain"
    "go-client/ipfs"
    "go-client/models"
    "go-client/encryption"
)

type AddPatientRequest struct {
    PatientName   string `json:"patient_name"`
    Age           uint   `json:"age"`
    Gender        string `json:"gender"`
    Disease       string `json:"disease"`
    ContactNumber string `json:"contact_number"`
    BloodGroup    string `json:"blood_group"`
    RecordData    string `json:"record_data"`  // Additional data to store in IPFS
}

func AddPatientHandler(w http.ResponseWriter, r *http.Request) {
    var req AddPatientRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Encrypt patient record data
    encryptedData, err := encryption.Encrypt(req.RecordData)
    if err != nil {
        http.Error(w, "Failed to encrypt data", http.StatusInternalServerError)
        return
    }

    // Upload encrypted record to IPFS
    ipfsHash, err := ipfs.UploadToIPFS(encryptedData)
    if err != nil {
        http.Error(w, "Failed to upload to IPFS", http.StatusInternalServerError)
        return
    }

    // Interact with the smart contract to store patient details
    err = blockchain.AddPatientRecord(req.PatientName, req.Age, req.Gender, req.Disease, req.ContactNumber, req.BloodGroup, ipfsHash)
    if err != nil {
        http.Error(w, "Failed to add patient record to blockchain", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Patient record added successfully"))
}

