import React from 'react';
import c from './Header.module.scss';
import MainButton from '../MainButton/MainButton';

class Header extends React.Component {
    render() {
      return (
        <>
          <div className={c.header}>
            <h1 className={c.icon}>Party Poll</h1>
            {
            //TODO: if logged in display instead of Login button the profil menu -> change when authentication is implemented
            }
            <MainButton text='Login' link='/login'/>
          </div>
        </>
      );
    }
}

export default Header;

