import React, { Component } from 'react'
import UserDataService from '../Services/UserDataService'
import jwt from 'jsonwebtoken'

export default class Register extends Component {
  
    constructor(props) {
        super(props)

    this.state = {
        Name:'',
        Email:'',
        Password:'',
        Gender:'',
        DateOfBirth:'',
        DogArray:[],
        successMessage: false,
        failMessage: false
        
    }
    this.handleChange = this.handleChange.bind(this)
    this.signupClicked = this.signupClicked.bind(this)
}
    

    handleChange(event) {
        this.setState(
            {
                [event.target.name]
                    : event.target.value
            }
        )
    }
    loginClicked= async(values)=> {
        let creds= {
            email:this.state.Email,
            password:this.state.Password
        }
        await
        UserDataService.login(creds).then(response => { if(response.data!=null){
            this.setState({ token: response.data.token })
            console.log(this.state.token);
            UserDataService.authUser(this.state.token)
            var decoded=jwt.decode(this.state.token)
            this.props.history.push(`/user/${decoded.ID}`)}
            else{}
        }).catch(() => {
        })               
    }
    signupClicked= async(values)=> {
        let creds= {
            Name:this.state.Name,
            Email:this.state.Email,
            Password:this.state.Password,
            Gender:this.state.Gender,
            DateOfBirth:this.state.DateOfBirth,
            DogArray:[]
        }
        
            UserDataService.createUser(creds).then(()=>{this.loginClicked()
                        this.setState({ showSuccessMessage: true });
})
            
        
    }

    render() {
        return (
            <div class="text-center">
                
                <div class="outer">
                <div class="middle">
                <div className="container text-center" >
                <h1>Sign Up</h1>
                <br></br>
                    Name: <input type="text" name="Name" value={this.state.Name} onChange={this.handleChange} />
                    <br></br>
                    {this.state.successMessage && <div className="alert alert-warning">Signup Succesful</div>}
                    {this.state.failMessage && <div className="alert alert-warning">Signup Unsuccesful</div>}
                    <br></br>
                    Email: <input type="text" name="Email" value={this.state.Email} onChange={this.handleChange} />
                    <br></br>
                    <br></br>
                    Password: <input type="password" name="Password" value={this.state.Password} onChange={this.handleChange} />
                    <br></br>
                    <br></br>
                    Gender: <input type="text" name="Gender" value={this.state.Gender} onChange={this.handleChange} />
                    <br></br>
                    <br></br>
                    DateOfBirth: <input type="text" name="DateOfBirth" value={this.state.DateOfBirth} onChange={this.handleChange} />
                    <br></br>
                    <br></br>
                    </div>
                    <button className="btn btn-success" onClick={this.signupClicked}>Signup</button>
                    </div>
                </div>
            </div>
        )
    }
}
