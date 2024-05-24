import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
import './App.css';
import Login from './components/Login';
import Register from './components/Register';
import Home from './components/Home';

const App = () => {
  const [token, setToken] = useState(localStorage.getItem('token') || '');

  const handleLogin = (token) => {
    setToken(token);
    localStorage.setItem('token', token);
  };

  const handleLogout = () => {
    setToken('');
    localStorage.removeItem('token');
  };

  return (
    <Router>
      <div className="app">
        <Switch>
          <Route path="/login">
            {token ? <Redirect to="/" /> : <Login setToken={handleLogin} />}
          </Route>
          <Route path="/register">
            {token ? <Redirect to="/" /> : <Register />}
          </Route>
          <Route path="/">
            {token ? <Home token={token} handleLogout={handleLogout} /> : <Redirect to="/login" />}
          </Route>
        </Switch>
      </div>
    </Router>
  );
};

export default App;
