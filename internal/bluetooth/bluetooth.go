package bluetooth

import (
	"time"

	bt "tinygo.org/x/bluetooth"
)

var (
    adapter = bt.DefaultAdapter

    heartRateServiceUUID = bt.ServiceUUIDHeartRate
    heartRateCharacteristicUUID = bt.CharacteristicUUIDHeartRateMeasurement
)


func Scan() {
    adapter.Enable()

    go func() {
        adapter.Scan(func(a *bt.Adapter, sr bt.ScanResult) {
            println("Device:", sr.Address.String(), sr.RSSI, sr.LocalName())
        })
    }()

    select {
        case <-time.After(10 * time.Second):
            println("Stopping scanning")
            adapter.StopScan()
    }
}

func Connect() {

}

func Disconnect() {

}
