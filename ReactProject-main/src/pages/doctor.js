import React from 'react';
import Users from './users';

const Doctor = () => {
  return (
    <div>
      <h1>hello,user (user type: doctor) </h1> <br></br>
      <h2 className="myslots"> my slots</h2>
      <table style={{ border: 1 }}>
        <tr>
          <th>date</th>
          <th>hour</th>
          <th> </th>
          <th> </th>
        </tr>
        <tr>
          <td>10/11/2023</td>
          <td>1:00 pm</td>
          <td>edit</td>
          <td>cancle</td>
        </tr>
        <tr>
          <td>15/3/2023</td>
          <td>5:00 pm</td>
          <td>edit</td>
          <td>canscl</td>
        </tr>
      </table>
      <br></br>
      <h3 className="createslot">create slot</h3>
      <table style={{ border: 1 }}>
        <tr>
          <th>date</th>
          <th>10/11/2023</th>
        </tr>
        <tr>
          <th>hour</th>
          <th>1:00 pm</th>
        </tr>
      </table>
      <br></br>
      <button type="add slot">
        <b>add slot</b>
      </button>
      <br></br>
      <br></br>
      <users />
    </div>
  );
}

export default Doctor;
