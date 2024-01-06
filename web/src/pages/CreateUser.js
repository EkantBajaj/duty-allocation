import { useState } from 'react';
import { createUser } from '../services/api';
import '../CreateUser.css';
import Header from "../components/Header";


const fieldConfig = [
    { label: 'Name', state: 'name', type: 'text', required: true },
    { label: 'Email', state: 'email', type: 'email', required: false },
    { label: 'Relative Name', state: 'relative_name', type: 'text', required: true },
    { label: 'Gender', state: 'gender', type: 'text', required: true },
    { label: 'Mobile Number', state: 'mobile_number', type: 'tel', required: true },
    { label: 'Aadhar Number', state: 'aadhar_number', type: 'text', required: true },
    { label: 'Address', state: 'address', type: 'text', required: true },
    { label: 'City', state: 'city', type: 'text', required: true },
    { label: 'Badge ID', state: 'badge_id', type: 'text', required: false },
    { label: 'Pincode', state: 'pin_code', type: 'text', required: true },
    { label: 'Emergency Number', state: 'emergency_number', type: 'tel', required: true },
    { label: 'Birth Date', state: 'birth_date', type: 'date', required: true },
    { label: 'Initiation Date', state: 'initiation_date', type: 'date', required: true },
    { label: 'Qualification', state: 'qualification', type: 'text', required: true },
    { label: 'Profession', state: 'profession', type: 'text', required: true },
    { label: 'Marital Status', state: 'marital_status', type: 'text', required: true },
    { label: 'Blood Group', state: 'blood_group', type: 'text', required: true },
    { label: 'Department', state: 'department', type: 'text', required: true },
    { label: 'Zone Badge ID', state: 'zone_badge_id', type: 'text', required: true },
    { label: 'Zone Department', state: 'zone_department', type: 'text', required: true },
    { label: 'Introduced By', state: 'introduced_by', type: 'text', required: true },
    { label: 'Center', state: 'center', type: 'text', required: true },
    { label: 'Sub Center', state: 'sub_center', type: 'text', required: true },
    { label: 'Remarks', state: 'remarks', type: 'text', required: true },
    { label: 'Password', state: 'password', type: 'password', required: true },
    { label: 'Initiated', state: 'initiated', type: 'checkbox' },
    { label: 'Photo Badge', state: 'photo_badge', type: 'checkbox' },
    { label: 'Badge Printed', state: 'badge_printed', type: 'checkbox' },
];

const CreateUser = () => {
    const [formData, setFormData] = useState(
        Object.fromEntries(fieldConfig.map(({ state, type }) => [state, type === 'checkbox' ? false : '']))
    );

    const handleChange = (e, state) => {
        setFormData({
            ...formData,
            [state]: e.target.type === 'checkbox' ? e.target.checked : e.target.value || false,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const response = await createUser(formData);
            alert(`User created successfully. Badge ID: ${response.badgeId}`);
            e.target.reset();
        } catch (error) {
            console.error(error);

            // Show error message in pop-up
            alert('Failed to create user');
        }
        window.location.reload();
    };

    return  (
    <div>
        <Header />
        <div className="create-user">
            <h2>Create User</h2>
            <form onSubmit={handleSubmit}>
                {fieldConfig.map(({ label, state, type, required }) => (
                    <div key={state}>
                        <label htmlFor={state}>{label}:</label>
                        <input
                            type={type}
                            id={state}
                            value={formData[state]}
                            onChange={(e) => handleChange(e, state)}
                            required={required}
                        />
                    </div>
                ))}

                <button type="submit">Create User</button>
            </form>
        </div>
    </div>
);
};

export default CreateUser;