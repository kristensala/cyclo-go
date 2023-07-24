import { Component, createResource, For, lazy } from 'solid-js';

import { ScanDevices } from '../wailsjs/go/main/App';

const App: Component = () => {
    const fetchDevices = async () => (await ScanDevices());
    const [devices] = createResource(fetchDevices);

    return (
        <div>
            {devices.loading &&
                <div>scanning for devices</div>
            }

            {devices() && (
                <For each={devices()}>
                    {(device) => (
                        <div>{device.Name} {device.Address}</div>
                    )}
                </For>
            )}

            
        </div>
    );
};

export default App;
