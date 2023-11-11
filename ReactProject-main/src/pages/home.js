import React from 'react';
import Users from './users';

const Home = () => {
  return (
    <div>
      <h1 className="signup">sign up</h1>
      <label for="Email">Email :</label>
      <input type="text" id="Email"></input> <br></br>
      <label for="password">password :</label>
      <input type="text" id="password"></input>
      <br></br>
      <label for="type">user type : </label>
      <label>patient:</label>
      <input name="user type" value="patient" type="radio"></input>
      <label>doctor:</label>
      <input name="user type" value="doctor" type="radio"></input> <br></br>
      <button type="sign up">
        <b>sign up</b>
      </button>
      <h2 className="login">log in</h2>
      <label for="Email">Email :</label>
      <input type="text" id="Email"></input> <br></br>
      <label for="password">password :</label>
      <input type="text" id="password"></input>
      <br></br>
      <button type="sign up">
        <b>login</b>
      </button>
      <users />
    </div>
  );
}

export default Home;
