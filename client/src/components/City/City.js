import React, { Component } from 'react';
import './Room.css'

class Room extends Component {

  constructor(props) {
    super(props)

    this.initialState = {
      // hit the crawler endpoints and get new info
    }
  }

  state = {}

  componentDidMount() {

  }

  render () {
    return (
      <div className="Room">
        <h1>Room</h1>
      </div>
    )
  }
}