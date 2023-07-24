package bluetooth

import (
    "time"

    bt "tinygo.org/x/bluetooth"
)

type Btle struct {
    scanResult []Device
}

var (
    adapter = bt.DefaultAdapter

    heartRateServiceUUID = bt.ServiceUUIDHeartRate
    heartRateCharacteristicUUID = bt.CharacteristicUUIDHeartRateMeasurement
)


// So this is an async function
// <- means we are awaiting something
func (btle *Btle) Scan() <- chan []Device {
    adapter.Enable()

    r := make(chan []Device)

    go func() {
        adapter.Scan(func(a *bt.Adapter, sr bt.ScanResult) {
            if !contains(btle.scanResult, sr.Address.String()) {
                btle.scanResult = append(btle.scanResult, Device {
                    Name: sr.LocalName(),
                    Address: sr.Address.String(),
                })
            }
        })

        r <-btle.scanResult
    }()

    /*
This also works

time.Sleep(10 * time.Second)
println("Stopping scanning")
*/

    select {
        case <-time.After(10 * time.Second):
            println("Stopping scan!")
            adapter.StopScan()
    }

    return r
}

func contains(array []Device, address string) bool {
    for _, device := range array {
        if device.Address == address {
            return true
        }
    }

    return false
}
func (btle *Btle) Connect() {

}

func (btle *Btle) Disconnect() {

}
