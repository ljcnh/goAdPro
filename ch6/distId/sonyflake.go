package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"os"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineId,
	}
	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(id)
}

func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID = readMachineIDFromLocalFile()
	if machineID == 0 {
		machineID, err = generateMachineID()
		if err != nil {
			return 0, err
		}
	}
	return machineID, err
}

func checkMachineId(machineID uint16) bool {
	saddResult, err := saddMachineIDToRedisSet()
	if err != nil || saddResult == 0 {
		return true
	}
	err = saveMachineIDToLocalFile(machineID)
	if err != nil {
		return true
	}
	return false
}

func saveMachineIDToLocalFile(id uint16) error {
	// ...
	return nil
}

func saddMachineIDToRedisSet() (uint16, error) {
	// ...
	return 0, nil
}

func generateMachineID() (uint16, error) {
	// ...
	return 0, nil
}

func readMachineIDFromLocalFile() uint16 {
	// ...
	return 0
}
