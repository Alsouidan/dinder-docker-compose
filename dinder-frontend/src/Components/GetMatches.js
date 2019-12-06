import React, { Component } from 'react';
import UserDataService from '../Services/UserDataService';
import DogDataService from '../Services/DogDataService';

class GetMatches extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userId: this.props.match.params.id1,
      dogId: this.props.match.params.id2
    };
    this.goToUser = this.goToUser.bind(this);
  }

  componentDidMount() {
    DogDataService.getDogMatches(this.state.dogId).then(res => {
      console.log(res);
      this.setState({ matches: res });
    });
  }
  goToUser(id) {
    this.props.history.push(`/userProfile/${id}`);
  }
  render() {
    console.log(this.state.matches);
    return (
      <div className="container">
        <table className="table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Breed</th>
              <th>Gender</th>
              <th>DOB</th>
              <th>Color</th>
              <th>Owner</th>
            </tr>
          </thead>
          <tbody>
            {this.state.matches &&
              this.state.matches.map(dog => (
                <tr key={dog._id}>
                  <td>{dog.Name}</td>
                  <td>{dog.Breed}</td>
                  <td>{dog.Gender}</td>
                  <td>{dog.DateOfBirth}</td>
                  <td>{dog.Colour}</td>
                  <td>
                    {' '}
                    <button
                      className="btn btn-success"
                      onClick={() => this.goToUser(dog.Owner_id)}
                    >
                      Get User Info
                    </button>
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      </div>
    );
  }
}
export default GetMatches;
