import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Replace with your backend server URL

export const getSpots = () => axios.get(`${API_BASE_URL}/spots`);

export const getActiveSpotUserCount = () => axios.get(`${API_BASE_URL}/spot-users/active-count`);

// Get users for a spot
export const getUsersForSpot = (spotId) => {
    return axios.get(`${API_BASE_URL}/spot-users/active-users/${spotId}`);
};

// Add user to a spot
export const addUserToSpot = (requestData) => {
  return axios.post(`${API_BASE_URL}/spot-users`, requestData);
};

// Delete user from a spot
export const deleteUserFromSpot = (userId) => {
    return axios.put(`${API_BASE_URL}/spot-users/user/${userId}`);
};
