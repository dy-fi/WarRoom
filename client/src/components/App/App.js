import React, { Component } from 'react';
import './App.css';
import RoomForm from '../Room-form/Room-form'

class App extends Component {

  constructor(props) {
    super(props)
    this.state = { cities: [] }
  }

  AllCities() {
    fetch('http://localhost:8000/cities/all')
      .then(res => res.text)
        .then(res => this.setState({cities: res}));
  }

  componentDidMount() {
    this.AllCities();
  }

  
  render() {
    return (<div className="App"> 
      {this.state.cities}

      {RoomForm}
    </div>);
  }

}

export default App;
