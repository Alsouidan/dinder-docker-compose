import React, { Component } from 'react'
import { Link, withRouter } from 'react-router-dom'
import UserDataService from '../Services/UserDataService'


class MenuComponent extends Component {

    render() {
        //const isUserLoggedIn = LoginComponent.isUserLoggedIn();
        

        return (
            <header>
                <nav className="navbar navbar-expand-lg navbar-light bg-light">
                    <div><a href="https://www.google.com" className="navbar-brand">Dinder</a></div>
                    <ul className="navbar-nav">
                    </ul>
                    <ul className="navbar-nav navbar-collapse justify-content-end">
                        {<li><Link className="nav-link" to="/login">Login</Link></li>}
                        {<li><Link className="nav-link" to="/logout" onClick={UserDataService.authUser('')}>Logout</Link></li>}
                    </ul>
                </nav>
            </header>
        )
    }
}

export default withRouter(MenuComponent)