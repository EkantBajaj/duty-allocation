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
            <h1>Radha Soami Satsang Beas- Saharanpur Major Center</h1>
            {/* Entry/Exit Form button */}
        <   Link to="/entry-exit-form">
            <button className="entry-exit-form-button">
                Gate Entry
            </button>
        </Link>
            {/* Logout button */}
            <button onClick={handleLogout} className="logout-button">
                Logout
            </button>
        </div>
    );
};

export default Header;
