import React, { useState, useEffect } from 'react';
import { getUsersForSpot, addUserToSpot, deleteUserFromSpot } from '../services/api';
import {useParams} from "react-router-dom";
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import Header from "../components/Header";


const SpotUsers = ({ spot }) => {
  const [users, setUsers] = useState([]);
  const [newBadgeID, setNewBadgeID] = useState('');
  const [newOutTime, setNewOutTime] = useState(null);
  const [error, setError] = useState(null);
  const params = useParams();
  const currentDate = new Date();
  const nextDate = new Date();
  nextDate.setDate(currentDate.getDate() + 1);
  currentDate.setHours(currentDate.getHours() + 1);

  useEffect(() => {
    // Fetch users for the spot
    fetchUsers();
  }, []);

  const fetchUsers = () => {
    getUsersForSpot(params.id)
      .then((response) => setUsers(response.data || []))
      .catch((error) => console.error(error));
  };

  const handleAddUser = () => {
    if (newBadgeID && newOutTime) {
      const newUser = {
        spot_id: params.id,
        badge_id: newBadgeID,
        out_time: newOutTime,
      };

      addUserToSpot(newUser)
        .then(() => {
          setNewBadgeID('');
          setNewOutTime(null);
          setError(null);
          fetchUsers();
        })
        .catch((error) => setError(error.response.data.error || 'An error occurred'));
        setTimeout(() => setError(null), 10000);
    } else {
      setError('Please enter badge ID and select out time');
    }
  };

  const handleDeleteUser = (userId) => {
    deleteUserFromSpot(userId)
      .then(() => {
        fetchUsers();
      })
      .catch((error) => setError(error.response?.data?.error || 'An error occurred'));
      setTimeout(() => setError(null), 10000);
  };

    return (
        <div>
            <Header />
            <h1 style={{ display: 'flex', justifyContent: 'space-between' }}>Sewa Point: {params.name}</h1>
            <div className='responsive-text'>
              <p>
                <input
                    type="text"
                    placeholder="Badge ID"
                    value={newBadgeID}
                    onChange={(e) => setNewBadgeID(e.target.value)}
                    style={{ fontSize: '1.5em' }}
                />
                <DatePicker
                    selected={newOutTime}
                    onChange={(date) => setNewOutTime(date)}
                    showTimeSelect
                    timeFormat="HH:mm"
                    timeIntervals={30}
                    timeCaption="Time"
                    dateFormat="yyyy-MM-dd HH:mm"
                    placeholderText="Select out time"
                    className='date-picker'
                    minDate={currentDate}
                    maxDate={nextDate}
                    filterTime={(time) => {
                      const currentTime = new Date();
                      // Only enable times that are later than the current time
                      return time.getTime() > currentTime.getTime();
                  }}
                />
                <button onClick={handleAddUser} style={{ fontSize: '1.5em' }}>Check In</button>
            </p></div>
            {error && <div className="error">{error}</div>}
            <table className='responsive-table'>
            <thead>
            <tr>
                <th>Name</th>
                <th>BadgeId</th>
                <th>Gender</th>
                <th>Contact</th>
                <th>In Time</th>
                <th>Out Time</th>
                <th>Check Out</th>
            </tr>
            </thead>
            <tbody>
                {users.map((user) => (
                    <tr>
                        <td>{user.Name}</td>
                        <td>{user.BadgeID}</td>
                        <td>{user.Gender}</td>
                        <td>{user.MobileNumber}</td>
                        <td>{user.InTimeString}</td>
                        <td>{user.OutTimeString}</td>
                        <td><button className='responsive-text' onClick={() => handleDeleteUser(user.ID)}>Check Out</button></td>

                    </tr>
                ))}
            </tbody>
            </table>
        </div>
    );
};

export default SpotUsers;
