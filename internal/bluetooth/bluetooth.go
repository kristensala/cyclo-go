package bluetooth

import (
	bt "tinygo.org/x/bluetooth"
)

type Btle struct {
    HeartRate uint8
}

var (
    adapter = bt.DefaultAdapter

    heartRateServiceUUID = bt.ServiceUUIDHeartRate
    heartRateCharacteristicUUID = bt.CharacteristicUUIDHeartRateMeasurement

    cyclingPowerServiceUUID = bt.ServiceUUIDCyclingPower
    cyclingSpeedAndCadenceUUID = bt.ServiceUUIDCyclingSpeedAndCadence
)

func (btle *Btle) Start() {
    adapter.Enable()

    ch := make(chan bt.ScanResult, 1)

    println("scanning")
    err := adapter.Scan(func(a *bt.Adapter, sr bt.ScanResult) {
        if sr.Address.String() == "EA:FE:62:36:C4:83" {
            adapter.StopScan()
            ch <-sr
        }
    })

    var device *bt.Device
    select {
    case result := <-ch:
        device, err = adapter.Connect(result.Address, bt.ConnectionParams{})
        if err != nil {
            println(err.Error())
        }

        println("connected to ", result.Address.String())
    }

    btle.getHeartRate(*device)
}

func (btle *Btle) getHeartRate(device bt.Device) {
    services, _ := device.DiscoverServices([]bt.UUID {heartRateServiceUUID})

    if len(services) == 0 {
        panic("hmm...")
    }

    service := services[0]
    println("Found service", service.UUID)
    chars, err := service.DiscoverCharacteristics([]bt.UUID{heartRateCharacteristicUUID})

    if err != nil {
        println("Error discovering:", err)
    }

    if len(chars) == 0 {
        println("could not read heartrate data")
    }

    char := chars[0]
    char.EnableNotifications(func(buf []byte) {
        btle.HeartRate = uint8(buf[1])
        println("heart", uint8(buf[1]))
    })

    select {}
}


