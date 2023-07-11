import React from 'react';

const Spot = ({ name, userCount, color, totalPeople , minPeople, onClick }) => (
    <tr onClick={onClick} >
        <td>{name}</td>
        <td>{totalPeople}</td>
        <td>{minPeople}</td>
        <td style={{ color }}>{userCount}</td>
    </tr>
);

export default Spot;
