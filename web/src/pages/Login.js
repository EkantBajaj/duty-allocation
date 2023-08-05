import React, { useState } from 'react';
import { login } from '../services/api';
import '../Login.css';

const Login = ({ onLogin }) => {
    const [badgeID, setBadgeID] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            // Call the login API with badgeID and password
            const response = await login(badgeID, password);

            // If login is successful, call the onLogin function with the received token
            onLogin(response.data.token);
        } catch (error) {
            setError('Invalid badge ID or password. Please try again.');
            setTimeout(() => {
                setError('');
            }, 5000);
        }
    };

    return (
        <div className="login-container">
            <div className="login-content-box">
            <form className="login-form" onSubmit={handleSubmit}>
                <img src="/rssb-logo.jpeg" alt="Login" className="login-image" />
                <div className="form-group">
                    <input type="text" value={badgeID} onChange={(e) => setBadgeID(e.target.value)} placeholder="Badge ID" />
                </div>
                <div className="form-group">
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" />
                </div>
                {error && <div className="error-message">{error}</div>}
                <button type="submit" className="login-button">Login</button>
            </form>
            </div>
        </div>
    );
};

export default Login;
