import React from 'react';
import Button from '@mui/material/Button';
import c from './MainButton.module.scss';
import { Link } from "react-router-dom";

interface MainButtonProps {
  text: string;
  link: string;
}

const MainButton: React.FC<MainButtonProps> = ({ text, link }) => {
  return (
    <div>
      <Button className={c.btn} component={Link} to={link} variant="contained">
        {text}
      </Button>
    </div>
  );
};

export default MainButton;
