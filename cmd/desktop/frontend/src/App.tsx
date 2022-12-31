import type { Component } from 'solid-js';
import { createSignal, onCleanup } from "solid-js";
import { render } from "solid-js/web";

import logo from './logo.svg';
import styles from './App.module.css';
import {Greet} from "../wailsjs/go/main/App";
import {CalculateMentions} from "../wailsjs/go/eventbucket/EventBucket";

const App: Component = () => {
  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt="logo" />
        <p>
          <EventCounterComponent />
          <CalcMentions />
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
    </div>
  );
};

export default App;

const CountingComponent = () => {
  const [count, setCount] = createSignal(0);
  const interval = setInterval(
      () => setCount(c => c + 1),
      1000
  );
  onCleanup(() => clearInterval(interval));
  return <div>Count value is {count()}</div>;
};

const EventCounterComponent = () => {
  const [g, setString] = createSignal("x");
  const interval2 = setInterval(
      function () {
        Greet("sadfsda").then(function (e) {
          setString(e)
        })
      },
      1000
  )
  onCleanup(() => clearInterval(interval2));
  return <div>Number of events that have been indexed: {g()}</div>;
};

const CalcMentions = () => {
  return (
      <>
        <button onClick={calculate}>Calculate</button>
      </>
  );
};

function calculate() {
    console.log()
    CalculateMentions().then(function (result) {
        console.log(result)
    })
}