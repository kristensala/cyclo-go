package bluetooth

type Device struct {
    Name string
    Address string
    IsConnected bool
    Class string
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
