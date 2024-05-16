import React, { useState } from 'react';
import InputField from "../Components/InputField";
import PersonSharpIcon from '@mui/icons-material/PersonSharp';
import VpnKeySharpIcon from '@mui/icons-material/VpnKeySharp';
import c from "./Login.module.scss";
import MainButton from '../Components/MainButton/MainButton';
import PageHeader from '../Components/PageHeader/PageHeader';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleUsernameChange = (value: string) => {
    setUsername(value);
  };

  const handlePasswordChange = (value: string) => {
    setPassword(value);
  };

  const handleSubmit = async () => {
    try {
      const response = await fetch('http://localhost:3001/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
      });
  
      if (response.ok) {
        const data = await response.json();
        if (data.token) {
          localStorage.setItem('token', data.token);
          window.location.href = '/dashboard';
          console.log('Login successful');
        } else {
          setError('Invalid username or password');
          console.error('Login failed');
        }
      } else {
        setError('An error occurred. Please try again.');
        console.error('An error occurred during login');
      }
    } catch (error) {
      setError('An error occurred. Please try again.');
      console.error('Error occurred during login:', error);
    }
  };
  
  return (
    <>
      <PageHeader heading=" " link="/"/>
    <div className={c.container}>
      <h1 className={c.title}>Party Poll</h1> 
      <hr className={c.separator} /> 
      <p>Log in to see your polls!</p>
      <InputField
        startIcon={<PersonSharpIcon className={c.personSVG} />}
        label={"Username"}
        placeholder={"Username"}
        onChange={handleUsernameChange}
      />
      <InputField
        startIcon={<VpnKeySharpIcon className={c.personSVG} />}
        label={"Password"}
        placeholder={"Password"}
        type="password"
        onChange={handlePasswordChange}
      /> 
      {error && <p className={`${c.error} ${c.errorMessage}`}>{error}</p>}
      <MainButton text={"Submit"} onClick={handleSubmit} />
    </div>
</>
  );
};
  
export default Login;
