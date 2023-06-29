package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(startTime string, machineID uint16) (err error) {
	sonyMachineID = machineID
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

func GenID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sonyflake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}

func main() {
	if err := Init("2020-07-01", 1); err != nil {
		fmt.Println("init failed err:", err)
		return
	}
	ID, err := GenID()
	if err != nil {
		fmt.Println("GenID failed")
		return
	}
	fmt.Println(ID)

}
