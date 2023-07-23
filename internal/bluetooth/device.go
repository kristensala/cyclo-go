package bluetooth

type Device struct {
    name string
    address string
    isConnected bool
    class string
}

type DeviceType string;
const (
    HearRateMonitor DeviceType = "Hear rate monitor"
    TurboTrainer DeviceType = "Turbo trainer"
    Unknown DeviceType = "unknown"
)

func (d *Device) GetDeviceType() DeviceType {
    return HearRateMonitor;
}
