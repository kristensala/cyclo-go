import { Component, createResource, createSignal, For, JSXElement, lazy } from 'solid-js';
import { ScanDevices, FetchHeartRate } from '../wailsjs/go/main/App';

const App: Component = () => {
    const [heartRate, setHeartRate] = createSignal(0)

    const fetchDevices = async () => (await ScanDevices());
    const [test] = createResource(fetchDevices);

    setInterval(async () => {
        let t = await FetchHeartRate()
        console.log(t)

        setHeartRate(t)
    })

    return (
        <div>
            HEART RATE: {heartRate()}
        </div>
    );
};

export default App;
