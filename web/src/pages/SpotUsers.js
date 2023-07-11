import React, { useState, useEffect } from 'react';
import { getUsersForSpot, addUserToSpot, deleteUserFromSpot } from '../services/api';
import {useParams} from "react-router-dom";
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

const SpotUsers = ({ spot }) => {
  const [users, setUsers] = useState([]);
  const [newBadgeID, setNewBadgeID] = useState('');
  const [newOutTime, setNewOutTime] = useState(null);
  const [error, setError] = useState(null);
  const params = useParams();

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
            <h2>Users for Spot: {params.name}</h2>
            <div>
                <input
                    type="text"
                    placeholder="Enter badge ID"
                    value={newBadgeID}
                    onChange={(e) => setNewBadgeID(e.target.value)}
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
                />
                <button onClick={handleAddUser}>Add User</button>
            </div>
            {error && <div className="error">{error}</div>}
            <table>
            <thead>
            <tr>
                <th>Name</th>
                <th>BadgeId</th>
                <th>Gender</th>
                <th>Initiated</th>
                <th>Action</th>
            </tr>
            </thead>
            <tbody>
                {users.map((user) => (
                    <tr>
                        <td>{user.Name}</td>
                        <td>{user.BadgeID}</td>
                        <td>{user.Gender}</td>
                        <td>{user.Initiated ? "Yes" : "No"}</td>
                        <td><button onClick={() => handleDeleteUser(user.ID)}>Delete User</button></td>

                    </tr>
                ))}
            </tbody>
            </table>
        </div>
    );
};

export default SpotUsers;
