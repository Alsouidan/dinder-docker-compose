import React, { Component } from 'react';
import './App.css';
import DinderApp from '../src/Components/DinderApp';
//import Homepage from './Components/UserHomepage'



class App extends Component {
  render() {
    return (
      <div className="container">
        <DinderApp />
      </div>
    );
  }
}

export default App;
