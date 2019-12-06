import React, { Component } from 'react';
import UserDataService from '../Services/UserDataService';
import DogDataService from '../Services/DogDataService';

class UserProfile extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userId: this.props.match.params.id
    };
  }

  componentDidMount() {
    UserDataService.retrieveUser(this.state.userId).then(e => {
      this.setState({ user: e });
    });
  }

  render() {
    let user = this.state.user;
    return (
      <div className="container">
        <table className="table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Email</th>
              <th>Gender</th>
              <th>DOB</th>
            </tr>
          </thead>
          <tbody>
            {this.state.user && (
              <tr key={user._id}>
                <td>{user.Name}</td>
                <td>{user.Email}</td>
                <td>{user.Gender}</td>
                <td>{user.DateOfBirth}</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    );
  }
}
export default UserProfile;
