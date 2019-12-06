import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import LoginComponent from './Login';
import LogoutComponent from './LogoutComponent';
import MenuComponent from './menu';
import Homepage from './UserHomepage';
import AddDogComponent from './AddDogComponent';
import Register from './UserSignup';
import Reccomendations from './Reccomendations';
import getMatches from './GetMatches'
import userProfile from './UserProfile'

class DinderApp extends Component {
  render() {
    return (
      <Router>
        <>
          <MenuComponent />
          <div class="text-center">
            <br></br>
            <br></br>
            <br></br>
            <br></br>
            <h1>Dinder </h1>
          </div>
          <Switch>
            <Route exact path="/" component={LoginComponent} />
            <Route exact path="/login" component={LoginComponent} />
            <Route exact path="/logout" component={LogoutComponent} />
            <Route exact path="/user/:id" component={Homepage} />
            <Route exact path="/userSignup" component={Register} />
            <Route path="/user/create/:id" exact component={AddDogComponent} />
            <Route exact path="/user/getRec/:id1/:id2" component={Reccomendations} />
            <Route exact path="/user/getMatches/:id1/:id2" component={getMatches} />
            <Route exact path="/userProfile/:id" component={userProfile} />
          </Switch>
        </>
      </Router>
    );
  }
}
export default DinderApp;
