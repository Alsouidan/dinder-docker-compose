import React, { Component } from 'react'
import UserDataService from '../Services/UserDataService';
import jwt from 'jsonwebtoken'

class LoginComponent extends Component {

    constructor(props) {
        super(props)

        this.state = {
            email:'',
            password:'',
            token:'',
            hasLoginFailed: false,
            showSuccessMessage: false
        }

        this.handleChange = this.handleChange.bind(this)
        this.loginClicked = this.loginClicked.bind(this)
        this.signupRedirect=this.signupRedirect.bind(this)
    }

    handleChange(event) {
        this.setState(
            {
                [event.target.name]
                    : event.target.value
            }
        )
    }

    loginClicked= async()=> {
        let creds= {
            email:this.state.email,
            password:this.state.password
        }
        await
        UserDataService.login(creds).then(response => { if(response.data!=null){
            console.log(creds);
            console.log(creds.email);
            console.log(creds.password);
            console.log(response);
            this.setState({ token: response.data.token })
            console.log(this.state.token);
            UserDataService.authUser(this.state.token)
            var decoded=jwt.decode(this.state.token)
            console.log(decoded);
            console.log(decoded.ID)
            this.props.history.push(`/user/${decoded.ID}`)}
            else{}
        }).catch(() => {
            this.setState({ showSuccessMessage: false })
            this.setState({ hasLoginFailed: true })
        })               
    }
    signupRedirect(){
        this.props.history.push(`/userSignup`)
    }
    isUserLoggedIn(){
        let token=this.state.token
        if(token===''){
            return false;
        }
        else{
            return true;
        }
    }
    logout(){
        this.setState.token =''
        this.isUserLoggedIn();
    }


    render() {
        return (
            <div class="text-center">
                
                <div class="outer">
                <div class="middle">
                <div className="container text-center" >
                <h1>Login</h1>
                <br></br>
                    {this.state.hasLoginFailed && <div className="alert alert-warning">Invalid Credentials</div>}
                    {this.state.showSuccessMessage && <div>Login Sucessful</div>}
                    Email: <input type="text" name="email" value={this.state.email} onChange={this.handleChange} />
                    <br></br>
                    <br></br>
                    Password: <input type="password" name="password" value={this.state.password} onChange={this.handleChange} />
                    <br></br>
                    <br></br>
                    <button className="btn btn-success" onClick={this.loginClicked}>Login</button>
                    </div>
                    <button className="btn btn-success" onClick={this.signupRedirect}>Signup</button>
                    </div>
                </div>
            </div>
        )
    }
}

export default LoginComponent