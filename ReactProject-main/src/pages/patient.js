

import React from 'react';
import Users from './users';

const Patient = () => {
  return (
    <div>
      <h1>hello,user (user type: patient) </h1> <br></br>
      <h2 className="myappointments">my appointments</h2>
      <table style={{ border: 1 }}>
        <tr>
          <th>appointment date</th>
          <th>appointment time </th>
          <th>doctor </th>
          <th> </th>
          <th> </th>
        </tr>
        <tr>
          <th>10/11/2023 </th>
          <th>1:00 pm </th>
          <th>dr. ahmed </th>
          <th>edit </th>
          <th>cancle </th>
        </tr>
        <tr>
          <th>15/3/2023 </th>
          <th>5:00 pm </th>
          <th>dr. mostafa </th>
          <th>edit </th>
          <th>cancle </th>
        </tr>
        <tr>
          <th>10/9/2021 </th>
          <th>1:00 pm </th>
          <th>dr. ali </th>
          <th>edit </th>
          <th>cancle </th>
        </tr>
        <tr>
          <th>15/4/2021 </th>
          <th>5:00 pm </th>
          <th>dr. mohamed </th>
          <th>edit </th>
          <th>cancle </th>
        </tr>
      </table>{" "}
      <br></br>
      <h3 className="createappointment">create new appointment</h3>
      <table style={{ border: 1 }}>
        <tr>
          <th>choose doctor</th>
          <th>
            <select type="doctor" id="language">
              <option value="">Select doctor</option>
              <option value="dr. ahmed">dr. ahmed</option>
              <option value="dr. mostafa">dr. mostafa</option>
              <option value="dr. ali">dr. ali</option>
              <option value="dr. mohamed">dr. mohamed</option>
            </select>
          </th>
        </tr>
        <tr>
          <th>appointment</th>
          <th> </th>
        </tr>
      </table>{" "}
      <br></br>
      <button type="reserve">
        <b>reserve</b>
      </button>{" "}
      <br></br>
      <users />
    </div>
  );
}

export default Patient;
