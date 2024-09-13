// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MedicalRecord {
    struct Patient {
        string patientName;
        uint age;
        string gender;
        string disease;
        string contactNumber;
        string bloodGroup;
        string ipfsHash;
    }

    mapping(address => Patient[]) private patients;
    mapping(address => mapping(address => bool)) private accessPermissions;

    event RecordAdded(address indexed owner, string ipfsHash);
    event AccessGranted(address indexed owner, address indexed to);
    event AccessRevoked(address indexed owner, address indexed to);

    function addPatientRecord(
        string memory _patientName,
        uint _age,
        string memory _gender,
        string memory _disease,
        string memory _contactNumber,
        string memory _bloodGroup,
        string memory _ipfsHash
    ) public {
        patients[msg.sender].push(Patient(_patientName, _age, _gender, _disease, _contactNumber, _bloodGroup, _ipfsHash));
        emit RecordAdded(msg.sender, _ipfsHash);
    }

    function grantAccess(address _to) public {
        accessPermissions[msg.sender][_to] = true;
        emit AccessGranted(msg.sender, _to);
    }

    function revokeAccess(address _to) public {
        accessPermissions[msg.sender][_to] = false;
        emit AccessRevoked(msg.sender, _to);
    }

    function getPatientRecords(address _owner) public view returns (Patient[] memory) {
        require(accessPermissions[_owner][msg.sender] || _owner == msg.sender, "No access to records");
        return patients[_owner];
    }
}

