import type { Component } from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';
import { Greet } from '../wailsjs/go/main/App';

const App: Component = () => {

    async function test() {
        var t = await Greet("test");
        console.log(t);
    }
    return (
        <div class={styles.App}>
            <header class={styles.header}>
                <img src={logo} class={styles.logo} alt="logo" />
                <p>
                    Edit <code>src/App.tsx</code> and save to reload.
                </p>
                <a
                    class={styles.link}
                    href="https://github.com/solidjs/solid"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Learn Solid
                </a>
            </header>

            <button onclick={test}>test</button>
        </div>
    );
};

export default App;
