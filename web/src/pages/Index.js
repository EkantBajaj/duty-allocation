import React, { useEffect, useState } from 'react';
import Spot from '../components/Spot';
import {useNavigate} from 'react-router-dom';
import { getSpots, getActiveSpotUserCount } from '../services/api';
import '../index.css'
import Header from "../components/Header";
import SpotUsers from "./SpotUsers";
const Index = () => {
    const [spots, setSpots] = useState([]);
    const [spotUserCounts, setSpotUserCounts] = useState([]);
    const [selectedSpot, setSelectedSpot] = useState(null);
    const navigate = useNavigate();
    const currentDate = new Date();
    const formattedDate = currentDate.toLocaleDateString('en-US');
    const currentDayName = currentDate.toLocaleDateString('en-US', { weekday: 'long' });
    const InchargeDayMapping = {
        0: 'Om Pal', // Sunday
        1: 'Deepak Makkar', // Monday
        2: 'Pappu Ji', // Tuesday
        3: 'Brijesh Kumar', // Wednesday
        4: 'Vipin Bhandari', // Thursday
        5: 'Rupak Yadav', // Friday
        6: 'Rohit Khera', // Saturday
    };
    const currentDay = currentDate.getDay();
    const currentPerson = InchargeDayMapping[currentDay];

    useEffect(() => {
        // Fetch spots
        getSpots()
            .then((response) => setSpots(response.data|| []))
            .catch((error) => console.error(error));

        // Fetch spot user counts
        getActiveSpotUserCount()
            .then((response) => setSpotUserCounts(response.data|| []))
            .catch((error) => console.error(error));
    }, []);
    const handleSpotClick = (spotId) => {
        const spot = spots.find((spot) => spot.ID === spotId);
        navigate('/user/'+ spotId+'/'+spot.Name)
    };

    return (
        <div>
            <Header />
            <div className='responsive-text' style={{ display: 'flex', justifyContent: 'space-between' }}>
            <h1>Sewadars Deployment Summary - Security Department</h1>
            <h2>{`${formattedDate}, ${currentDayName}, ${currentPerson}`}</h2>
        </div>
            <div style={{ display: 'flex', justifyContent: 'center' }}>
    <p><a href="/create-user" className="create-user-button"> New Sewadar</a></p>
            </div>
            <table className='responsive-table'>
                <thead>
                <tr>
                    <th>Sewa Point</th>
                    <th>Minimum Sewadars Required</th>
                    <th>Sewadars Deployed</th>
                </tr>
                </thead>
                <tbody>
            {spots.map((spot) => {
                // Find spot user count
                const spotUserCount = spotUserCounts.find((count) => count.ID === spot.ID);
                const userCount = spotUserCount ? spotUserCount.UserCount : 0;

                // Determine color based on user count
                let color = "";

                if (userCount < spot.MinPeople) {
                    color = "red";
                } else if (userCount >= spot.TotalPeople) {
                    color = "green";
                } else {
                    color = "orange";
                }

                return (
                    <Spot key={spot.ID} name={spot.Name} userCount={userCount} color={color} minPeople={spot.MinPeople} totalPeople={spot.TotalPeople} onClick={() => handleSpotClick(spot.ID)} />
                );
            })}</tbody></table>
            {selectedSpot && <SpotUsers spot={selectedSpot} />}
        </div>
    );
};

export default Index;
