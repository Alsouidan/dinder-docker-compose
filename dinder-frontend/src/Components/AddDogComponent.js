import React, { Component } from 'react';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import DogDataService from '../Services/DogDataService';
import UserDataService from '../Services/UserDataService';

class AddDogComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      id: this.props.match.params.id,
      Breed: '',
      Gender: '',
      DateOfBirth: '',
      Name: '',
      Colour: '',
      Weight: '',
      image_id: '',
      Owner_id: '',
      dog_id: ''
    };
    this.onSubmit = this.onSubmit.bind(this);
    this.validate = this.validate.bind(this);
            this.handleChange = this.handleChange.bind(this);

  }
  handleChange(event) {
    this.setState({
      [event.target.name]: event.target.value
    });
  }
  onSubmit = async values => {
        let { Name, Breed, Gender, id, DateOfBirth, Weight, Colour } = this.state;

    console.log('submiting');
    console.log(values);
    let Dog = {
      Breed: Breed,
      Gender: Gender,
      DateOfBirth: DateOfBirth,
      Name: Name,
      Colour: Colour,
      Weight: Weight,
      image_id: 0,
      Matched_IDs: [],
      Matched_by_IDs: [],
      Rejected_IDs: [],
      Owner_id: this.state.id
    };

    await DogDataService.createDog(Dog)
      .then(response => {
        this.setState({ dog_id: response.InsertedID });
        console.log(this.state.id);
        console.log(this.state.dog_id);
        UserDataService.addDog(this.state.id, { _id: this.state.dog_id }).then(
          response => {
            console.log(response);
          }
        );

        console.log(this.state.dog_id);
        this.props.history.push(`/user/${this.state.id}`);
      })
      .catch(() => {
        console.log('fail');
      });
  };

  componentDidMount() {
    console.log(this.state.id);
  }
  validate(values) {
    let errors = {};
    if (!values.name) {
      errors.name = 'Enter a Name';
    } else if (values.name.length < 5) {
      errors.name = 'Enter First and last Name Please';
    }

    return errors;
  }

  render() {
    let { Name, Breed, Gender, id, DateOfBirth, Weight, Colour } = this.state;
console.log(this.state)
    return (
      <div>
        <h3>Employee</h3>
        <Formik
          initialValues={{
            id,
            Name,
            Breed,
            Gender,
            DateOfBirth,
            Weight,
            Colour,
          }}
          onSubmit={this.onSubmit}
          validateOnChange={false}
          validateOnBlur={false}
          validate={this.validate}
          enableReinitialize={true}
        >
          {props => (
            <Form>
              <ErrorMessage name="name" component="div" className="alert alert-warning" />
              <fieldset className="form-group">
                <label>Name</label>
                <Field
                  className="form-control"
                  type="text"
                  name="Name"
                  value={this.state.Name}
                  onChange={this.handleChange}
                />
              </fieldset>
              <fieldset className="form-group">
                <label>Breed</label>
                <Field
                  className="form-control"
                  type="text"
                  name="Breed"
                  value={this.state.Breed}
                  onChange={this.handleChange}
                />
              </fieldset>
              <fieldset className="form-group">
                <label>Gender</label>
                <Field
                  className="form-control"
                  type="text"
                  name="Gender"
                  value={this.state.Gender}
                  onChange={this.handleChange}
                />
              </fieldset>
              <fieldset className="form-group">
                <label>DateOfBirth</label>
                <Field
                  className="form-control"
                  type="text"
                  name="DateOfBirth"
                  value={this.state.DateOfBirth}
                  onChange={this.handleChange}
                />
              </fieldset>
              <fieldset className="form-group">
                <label>Colour</label>
                <Field
                  className="form-control"
                  type="text"
                  name="Colour"
                  value={this.state.Colour}
                  onChange={this.handleChange}
                />
              </fieldset>
              <fieldset className="form-group">
                <label>Weight</label>
                <Field
                  className="form-control"
                  type="text"
                  name="Weight"
                  value={this.state.Weight}
                  onChange={this.handleChange}
                />
              </fieldset>
              {/* <fieldset className="form-group">
                                    <label>image_id</label>
                                    <Field className="form-control" type="text" name="image_id" />
                                </fieldset> */}
              <button className="btn btn-success" onClick={this.onSubmit} type="submit">
                Save
              </button>
            </Form>
          )}
        </Formik>
      </div>
    );
  }
}

export default AddDogComponent;
