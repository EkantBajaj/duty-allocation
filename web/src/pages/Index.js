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
            <h1>Spots</h1>
            <table>
                <thead>
                <tr>
                    <th>Name</th>
                    <th>Total People</th>
                    <th>Min People</th>
                    <th>User Count</th>
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
