import React, { useState, useEffect } from 'react';
import { Route, BrowserRouter, Routes } from 'react-router-dom';
import Index from './pages/Index';
import SpotUsers from './pages/SpotUsers';
import Login from './pages/Login';
const App = () => {
    const [userToken, setUserToken] = useState(localStorage.getItem('token') || '');

    // Function to handle user login
    const handleLogin = (token) => {
        // Save the token to local storage
        localStorage.setItem('token', token);
        console.log(token+" is the token saved ")
        // Set the userToken state to the received token
        setUserToken(token);
        setTimeout(() => {
            window.location.reload();
        }, 50)
    };


    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    {/* Render the Login component as the landing page if userToken is not set */}
                    {userToken ? (
                        <Route path="/" element={<Index />} />
                    ) : (
                        <Route path="/" element={<Login onLogin={handleLogin} />} />
                    )}
                    <Route path="/user/:id/:name" element={<SpotUsers />} />

                </Routes>
            </BrowserRouter>
        </div>
    );
};

export default App;
