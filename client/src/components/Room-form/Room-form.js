import React, { Component } from 'react';
import './Room-form.css';

class Roomform extends Component {
  constructor () {
    this.state = {
      values: [{
        name: '',

        urls: [{
          url: '',

          locations: [{
            label: '',
            xpath: "",
          }],
        }],

        updateInterval: 5,
      }],    
    }
  }

  changeNameHandler = event => {

    name = event.target.name;
    value = event.target.value;

    this.setState({
      values: {
        [name]: value
      }
    });
  }

  changeLocationHandler(i, event) {
    var newArray  = this.state.locations.slice();
    newArray.push(i.key, i.xpath);
    this.setState({locations: newArray});
  }

  submitHandler(event) {
    fetch('localhost:8000/cities', {
      method: 'POST',
      headers: {
        'Acccept': 'applications/json',
        'Content-Type': 'applications/json',
      },
        body: JSON.stringify({
          name: this.state.name,
          locations: this.state.locations,
          updateInterval: this.state.updateInterval,
        })
      })
  }

  addClick() {
    this.setState(prevState => ({
      locations: [...prevState, { location: null }]
    }));
  }

  removeClick(i) {
    let values = [...this.state.values];
    values.splice(i, 1);
    this.setState({ values });
  }
    
  render() {
    return (
      <form onSubmit={this.handleSubmit}>

        <input
          type="text"
          value={this.state.values.name || ""}
          name="name"
          onChange={e => this.changeNameHandler(e)}
        />

        <input 
          type="number"
          name="updateInterval"
          min="3"
          value={this.state.name || "5"}
        />

        {this.state.urls.locations.map((location) => (
          <div key={location.label.toString()}>

            <input
              type="text"
              value={location.label || ""}
              onChange={e => this.handleChange(location)}
            />

            <input
              type="button"
              value="remove"
              onClick={() => this.removeClick(location)}
            />
          </div>
        ))}

        <input type="button" value="add more" onClick={() => this.addClick()} />
        <input type="submit" value="Submit" />
      </form>
    );
  }
}

export default Roomform;