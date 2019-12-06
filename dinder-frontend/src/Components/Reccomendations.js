import React, { Component } from 'react';
import UserDataService from '../Services/UserDataService';
import DogDataService from '../Services/DogDataService';

class Reccomendations extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userId: this.props.match.params.id1,
      dogId: this.props.match.params.id2
    };
    this.like = this.like.bind(this);
    this.dislike = this.dislike.bind(this);
  }

  componentDidMount() {
    let res = DogDataService.retrieveDog(this.state.dogId).then(e => {
      console.log(e);
      this.setState({ dog: e });
      DogDataService.recomendDogs(e).then(res => {
        this.setState({ reccomendations: res });
        console.log(this.state);
      });
    });
  }
  like(id) {
    console.log('like');
    DogDataService.approveDog(this.state.dogId, id).then(res => {
      console.log(res);
      this.setState({
        reccomendations: this.state.reccomendations.filter(e => {
          return e._id !== id;
        })
      });
    });
  }
  dislike(id) {
    console.log('dislike');
    DogDataService.rejectDog(this.state.dogId, id).then(res => {
      console.log(res);
      this.setState({
        reccomendations: this.state.reccomendations.filter(e => {
          return e._id !== id;
        })
      });
    });
  }
  render() {
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
              <th>Like/Dislike</th>
            </tr>
          </thead>
          <tbody>
            {this.state.reccomendations &&
              this.state.reccomendations.map(dog => (
                <tr key={dog._id}>
                  <td>{dog.Name}</td>
                  <td>{dog.Breed}</td>
                  <td>{dog.Gender}</td>
                  <td>{dog.DateOfBirth}</td>
                  <td>{dog.Colour}</td>
                  <td>
                    <button
                      className="btn btn-success"
                      onClick={() => this.like(dog._id)}
                    >
                      {' '}
                      Like
                    </button>
                    <button
                      className="btn btn-success"
                      onClick={() => this.dislike(dog._id)}
                    >
                      Dislike
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
export default Reccomendations;
