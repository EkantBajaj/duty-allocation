import React from 'react';
import { Link } from 'react-router-dom';
import '../Header.css'; // Import the Header.css file

const Header = () => {
    const handleLogout = () => {
        // Clear the token from local storage
        localStorage.removeItem('token');
        // Reload the page to log the user out
        window.location.reload();
    };

    return (
        <div className="header">
            {/* Home logo to redirect to the root page */}
            <Link to="/">
                <img src="/rssb-logo.jpeg" alt="Home" className="logo-image" />
            </Link>
             {/* Create User button */}
            <Link to="/create-user" className="create-user-button">
                Create User
            </Link>
            {/* Logout button */}
            <button onClick={handleLogout} className="logout-button">
                Logout
            </button>
        </div>
    );
};

export default Header;
