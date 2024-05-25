import React, { useState, useEffect } from 'react';
import c from './Header.module.scss';
import MainButton from '../MainButton/MainButton';

const Header = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [username, setUsername] = useState('');

    useEffect(() => {
        const token = localStorage.getItem('token');
        const user = localStorage.getItem('username');
        if (token) {
            setIsLoggedIn(true);
            setUsername(user || '');
        }
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('username');
        setIsLoggedIn(false);
        setUsername('');
        window.location.href = '/';
    };

    if (!isLoggedIn) {
        return (
            <div className={c.header}>
                <h1 className={c.icon}>Party Poll</h1>
                {<MainButton text='Login' link='/login'/>}
            </div>
        );
    } else {
        return (
            <div className={c.header}>
                <h1 className={c.icon}>Party Poll</h1>
                {<div className={c.user_handling}>
                    <h4 className={c.username}>{username}</h4>
                    <MainButton text='Logout' onClick={handleLogout}/>
                </div>}
            </div>
        );
    }
};

export default Header;

