import React, { Component } from 'react';
import UserDataService from '../Services/UserDataService';

export default class Homepage extends Component {
  constructor(props) {
    super(props);

    this.state = { dogs: [], id: this.props.match.params.id };
  }
  componentDidMount = async () => {
    //why is this async
    console.log('mounted');
    var id = this.state.id;
    //await is already in api call no need to have it here
    await UserDataService.retrieveDogs(id).then(res => {
      console.log(res);
      if (!res) {
        this.setState({ dogs: [] });
      } else {
        this.setState({ dogs: res });
      }
    });
  };
  addDogClicked = async () => {
    let id = this.state.id;
    await this.props.history.push(`/user/create/${id}`);
  };
  render() {
    console.log(this.state);
    console.log('rendering');
    const ids = this.state.dogs.map(e => {
      return <li> {'Dog:' + e._id + '  '} </li>;
    });
    return (
      <div className="container">
        <h3>My Dogs</h3>
        {this.state.message && (
          <div class="alert alert-success">{this.state.message}</div>
        )}
        <div className="container">
          <table className="table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Breed</th>
                <th>Gender</th>
                <th>DOB</th>
                <th>Color</th>
                <th>Get Recomendations</th>
                <th>Get Matches</th>
              </tr>
            </thead>
            <tbody>
              {this.state.dogs.map(dog => (
                <tr key={dog._id}>
                  <td>{dog.Name}</td>
                  <td>{dog.Breed}</td>
                  <td>{dog.Gender}</td>
                  <td>{dog.DateOfBirth}</td>
                  <td>{dog.Colour}</td>
                  <td>
                    <button
                      className="btn btn-success"
                      onClick={() => {
                        console.log(dog);
                        this.props.history.push(
                          `/user/getRec/${this.props.match.params.id}/${dog._id}`
                        );
                      }}
                    >
                      Get Recommendations
                    </button>
                  </td>
                  <td>
                    <button
                      className="btn btn-success"
                      onClick={() => {
                        console.log(dog);
                        this.props.history.push(
                          `/user/getMatches/${this.props.match.params.id}/${dog._id}`
                        );
                      }}
                    >
                      Get Matches
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <div class="row">
          <div class="col-12 text-center">
            <button className="btn btn-success" onClick={this.addDogClicked}>
              Add
            </button>
          </div>
        </div>
      </div>
    );
  }
}
