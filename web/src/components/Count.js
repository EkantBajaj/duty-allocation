import React, { useEffect, useState } from 'react';
import { fetchTotalCount, fetchGateCounts } from '../services/api'; // Import the API functions
import '../Count.css';

function Counts() {
    const [totalCount, setTotalCount] = useState(0);
    const [gateCounts, setGateCounts] = useState([]);

    useEffect(() => {
        fetchTotalCount().then(setTotalCount);
        fetchGateCounts().then(data => setGateCounts(data));
    }, []);

    return (
        <div className="counts-container">
            <div className="total-count">
                <h3>Total Count</h3>
                <p>{totalCount}</p>
            </div>
            <div className="gate-counts">
                <h3>Gate Counts</h3>
                {gateCounts.map((gateCount, index) => (
                    <p key={index}>{gateCount.gate}: {gateCount.count}</p>
                ))}
            </div>
        </div>
    );
}

export default Counts;