import React from 'react';
import Button from '@mui/material/Button';
import c from './MainButton.module.scss';
import { Link } from "react-router-dom";

interface MainButtonProps {
  text: string;
  link?: string;
  onClick?: () => void; 
}

const MainButton: React.FC<MainButtonProps> = ({ text, link, onClick }) => {
  const handleClick = () => {
    if (onClick) {
      onClick();
    }
  };

  return (
    <div>
      {link ? (
        <Button className={c.Button} component={Link} to={link} variant="contained" onClick={handleClick}>
          {text}
        </Button>
      ) : (
        <Button className={c.Button} variant="contained" onClick={handleClick}>
          {text}
        </Button>
      )}
    </div>
  );
};

export default MainButton;
