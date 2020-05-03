import { Component } from 'react';
import React from 'react';
import './App.css';

export default class App extends Component {
  state = {
    title: 'hamsterapps.net',
  };

  render() {
    return (
      <div className="App">
        <h1>{this.state.title} is currently under development!</h1>
        <p>Please stay tuned for a release!</p>
      </div>
    );
  }
}
