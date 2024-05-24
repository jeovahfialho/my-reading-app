import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Home = ({ token, handleLogout }) => {
  const [day, setDay] = useState(1);
  const [reading, setReading] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchReading(day);
  }, [day]);

  const fetchReading = async (day) => {
    try {
      const response = await axios.get(`http://localhost:8080/readings/${day}`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      setReading(response.data);
      setError('');
    } catch (err) {
      setError('Reading not found');
      setReading(null);
    }
  };

  const nextDay = () => setDay(day + 1);
  const previousDay = () => day > 1 && setDay(day - 1);

  return (
    <div>
      <h1>Bible Reading Plan</h1>
      <button onClick={handleLogout}>Logout</button>
      {error && <p className="error">{error}</p>}
      {reading && (
        <div className="reading">
          <h2>Day {reading.Day}</h2>
          <p><strong>Period:</strong> {reading.Period}</p>
          <p><strong>First Reading:</strong> {reading.FirstReading}</p>
          <p><strong>Second Reading:</strong> {reading.SecondReading}</p>
          <p><strong>Third Reading:</strong> {reading.ThirdReading}</p>
        </div>
      )}
      <div className="navigation">
        <button onClick={previousDay} disabled={day === 1}>Previous</button>
        <button onClick={nextDay}>Next</button>
      </div>
    </div>
  );
};

export default Home;
