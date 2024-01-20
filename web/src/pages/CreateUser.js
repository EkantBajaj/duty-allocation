import { useState } from 'react';
import { createUser } from '../services/api';
import '../CreateUser.css';
import Header from "../components/Header";

const validateMobileNumber = (number) => {
    const regex = /^\d{10}$/;
    return regex.test(number);
};

const validateAadharNumber = (number) => {
    const regex = /^\d{16}$/;
    return regex.test(number);
};

const calculateAge = (birthDate) => {
    const ageDifference = Date.now() - new Date(birthDate).getTime();
    const ageDate = new Date(ageDifference);
    return Math.abs(ageDate.getUTCFullYear() - 1970);
};

const fieldConfig = [
    { label: 'Name*', state: 'name', type: 'text', required: true },
    { label: 'Email*', state: 'email', type: 'email', required: false },
    { label: 'Relative Name', state: 'relative_name', type: 'text', required: true },
    { label: 'Gender', state: 'gender', type: 'select', options: ['Male', 'Female'], required: true },
    { label: 'Mobile Number', state: 'mobile_number', type: 'tel', required: true },
    { label: 'Aadhar Number', state: 'aadhar_number', type: 'text', required: true },
    { label: 'Address', state: 'address', type: 'text', required: true },
    { label: 'City', state: 'city', type: 'text', required: true },
    { label: 'Pincode', state: 'pin_code', type: 'text', required: true },
    { label: 'Emergency Number', state: 'emergency_number', type: 'tel', required: true },
    { label: 'Birth Date', state: 'birth_date', type: 'date', required: true },
    { label: 'Initiation Date', state: 'initiation_date', type: 'date', required: false },
    { label: 'Qualification', state: 'qualification', type: 'select', options: ['Nil', '8th', '10th', '12th', 'Diploma', 'Graduate', 'Post-Graduate'],required: true },
    { label: 'Profession', state: 'profession', type: 'text', required: true },
    { label: 'Marital Status', state: 'marital_status', type: 'select', options: ['Married', "Un-Married", "Divorcee", "Widower"], required: true },
    { label: 'Blood Group', state: 'blood_group', type: 'select', options: ["A+","B+","O+","AB+","A-","B-","O-","AB-"], required: true },
    { label: 'Department', state: 'department', type: 'select', options: ["Security"], required: true },
    { label: 'Introduced By (Badge-ID)', state: 'introduced_by', type: 'text', required: true },
    { label: 'Center', state: 'center', type: 'select', options: ["Saharanpur-1"], required: true },
    { label: 'Sub Center', state: 'sub_center', type: 'select', options: ["Chilkana","Rasoolpur","Phandpuri","Pather","Nakur","Badgaon","Bhadarpur","Ghateda","Talheri Buzurug","Sona-Arjanpur","Saharanpur-2","Nanauta","Shermau","Ambheta Peer","Chhutmalpur","Gurunanakpura","Bilaspura","Islam Nagar","Gangoh","Randhol","Lakhnoti kalan","Behat","Hariya Baans","Deoband","Nagal","Manjhol","Gangoh","Santour","Teetron","Muzaffarabad","Beharigarh","Talapur"], required: true },
    { label: 'Remarks', state: 'remarks', type: 'text', required: true },
    { label: 'Initiated', state: 'initiated', type: 'checkbox', required: true },
];

const CreateUser = () => {
    const [formData, setFormData] = useState(
        Object.fromEntries(fieldConfig.map(({ state, type, options }) => [state, type === 'checkbox' ? false : type === 'select' ? options[0] : '']))
    );

    const handleChange = (e, state) => {
        setFormData({
            ...formData,
            [state]: e.target.type === 'checkbox' ? e.target.checked : e.target.value || false,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!validateMobileNumber(formData.mobile_number)) {
            alert('Please enter a valid 10-digit mobile number.');
            return;
        }
        if (!validateMobileNumber(formData.emergency_number)) {
            alert('Please enter a valid 10-digit emergency number.');
            return;
        }
        if (!validateAadharNumber(formData.aadhar_number)) {
            alert('Please enter a valid 16-digit aadhaar number.');
            return;
        }
        if (calculateAge(formData.birth_date) < 15) {
            alert('The person should be 15 years or older.');
            return;
        }
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
            <h2>New Sewadar</h2>
            <form onSubmit={handleSubmit}>
                {fieldConfig.map(({ label, state, type, options, required }) => (
                    <div key={state} className='form-field'>
                        <label htmlFor={state}>{label}:</label>
                        {type === 'select' ? (
                            <select
                                id={state}
                                value={formData[state]}
                                onChange={(e) => handleChange(e, state)}
                                required={required}
                            >
                                {options.map((option, i) => <option key={i} value={option}>{option}</option>
                                )}
                            </select>
                        ) : (
                            <input
                                type={type}
                                id={state}
                                value={formData[state]}
                                onChange={(e) => handleChange(e, state)}
                                required={required}
                            />
                        )}
                    </div>
                ))}

                <button type="submit">Create User</button>
            </form>
        </div>
    </div>
);
};

export default CreateUser;