import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Replace with your backend server URL

const api = axios.create({
    baseURL: API_BASE_URL,
});

// Function to set the user token in the Authorization header
const setAuthToken = (token) => {
    if (token) {
        console.log("setting token");
        api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    } else {
        console.log("deleting token");
        delete api.defaults.headers.common['Authorization'];
    }
};

// Get the user token from local storage (if available) on app load
const userToken = localStorage.getItem('token');
if (userToken) {
    setAuthToken(userToken);
}
console.log(userToken + " is the user token");

// API requests
export const login = (badgeID, password) => {
    const requestData = {
        badge_id: badgeID,
        password: password,
    };
    return api.post('/users/login', requestData);
};

export const getSpots = () => api.get('/spots');

export const getActiveSpotUserCount = () => api.get('/spot-users/active-count');

// Get users for a spot
export const getUsersForSpot = (spotId) => {
    return api.get(`/spot-users/active-users/${spotId}`);
};

// Add user to a spot
export const addUserToSpot = (requestData) => {
    return api.post('/spot-users', requestData);
};

// Delete user from a spot
export const deleteUserFromSpot = (userId) => {
    return api.put(`/spot-users/user/${userId}`);
};

// Create user API endpoint
export const createUser = async (userData) => {
    try {
        const response = await api.post('/users', userData);
        return response.data;
    } catch (error) {
        console.error(error);
        throw new Error('Failed to create user');
    }
};

// Get gates API endpoint
export const getGates = async () => {
    try {
      const response = await api.get('/gate/');
      return response.data;
    } catch (error) {
      console.error(error);
      throw new Error('Failed to get gates');
    }
  };

  export const postEntryExit = async (payload) => {
    try {
      const response = await api.post('/gate/entry', payload);
      return response.data;
    } catch (error) {
      console.error(error);
      throw new Error('Failed to post entry/exit');
    }
  };

  export const fetchTotalCount = async () => {
    try {
      const response = await api.get('/gate/total');
      return response.data.count;
    } catch (error) {
      console.error(error);
      throw new Error('Failed to get total count');
    }
  };

  export const fetchGateCounts = async () => {
    try {
      const response = await api.get('/gate/count');
      return response.data.map(item => ({ gate: item.name, count: item.count }));
    } catch (error) {
      console.error(error);
      throw new Error('Failed to get gate counts');
    }
  };

export default api;
