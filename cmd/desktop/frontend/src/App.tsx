import type { Component, onMount } from 'solid-js';
import {createSignal, For, onCleanup} from "solid-js";
import { render } from "solid-js/web";

import logo from './logo.svg';
import styles from './App.module.css';
import {Greet} from "../wailsjs/go/main/App";
import {CalculateMentions, CurrentOrder, SingleEvent, EventList} from "../wailsjs/go/eventbucket/EventBucket";

const App: Component = () => {
  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt="logo" />
        <p>
          <EventCounterComponent />
            <EventCounterComponent2 />
            <CalcMentions />
            <OneEvent />
        </p>
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

const EventCounterComponent2 = () => {
    const [g, setString] = createSignal("x");
    Greet("sadfsda").then(function (e) {
        setString(e)
    })
    return <div>Number of events that have been indexed: {g()}</div>;
};

const CalcMentions = () => {
  return (
      <>
        <button onClick={calculate}>Calculate</button>
      </>
  );
};

const OneEvent = () => {
    const [ev, setev] = createSignal("");
    SingleEvent().then(function (e){
        setev(e.Event.content)
        console.log(e.Event.content)
    })
    return (
        <>
            <p>
                {ev()}
            </p>
        </>
    )
}



function calculate() {
    console.log(59)
    CalculateMentions().then(function (result) {
        console.log(result)
        EventList().then(function (e) {
            const [elist, seteList] = createSignal(e);

            e.forEach(function (e2){
                console.log(e2)
            })
        })
        //currentOrder()
    })
}

