import React from 'react';
import Button from '@mui/material/Button';
import c from './MainButton.module.scss';

interface MainButtonProps {
  text: string;
}

const MainButton: React.FC<MainButtonProps> = ({ text }) => {
  return (
    <div className={c.btn}>
      <Button variant="contained">
        {text}
      </Button>
    </div>
  );
};

export default MainButton;
