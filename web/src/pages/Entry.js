import React, { useState, useEffect } from 'react';
import { getGates, postEntryExit } from '../services/api'; // Adjust the import path to match your file structure
import '../Entry.css';
import Counts from '../components/Count';
const EntryExitForm = () => {
    const [gate, setGate] = useState('');
    const [count, setCount] = useState(0);
    const [gates, setGates] = useState([]);
    const [updateCounts, setUpdateCounts] = useState(false);

    useEffect(() => {
        getGates()
            .then(data => {
                setGates(data);
            })
            .catch(error => {
                console.error(error);
            });
    }, []);

    const handleSubmit = (isExit) => {
        const payload = {
            gate_id: parseInt(gate),
            count: isExit ? -count : count,
        };

        postEntryExit(payload)
            .then(data => {
                console.log(data);
                setCount(0); // Reset the count field
            })
            .catch(error => {
                console.error(error);
            });
            setUpdateCounts(!updateCounts);
    };

    return (
        <div className="form-container">
      <h2>Gate Entry/Exit Form</h2>
      <form>
        <label>
          Gate:
          <select value={gate} onChange={e => setGate(e.target.value)}>
            <option value="">Select a gate</option>
            {gates.map((gate) => (
              <option key={gate.ID} value={gate.ID}>{gate.name}</option>
            ))}
          </select>
        </label>
        <label>
          Count:
          <input type="number" value={count} onChange={e => setCount(Number(e.target.value))} />
        </label>
        <div className="button-group">
          <button type="button" onClick={() => handleSubmit(false)} className="entry-button">
            Entry
          </button>
          <button type="button" onClick={() => handleSubmit(true)} className="exit-button">
            Exit
          </button>
        </div>
      </form>
        <Counts key={updateCounts}/>
    </div>
    );
};

export default EntryExitForm;