import React from 'react';
import Button from '@mui/material/Button';
import c from './MainButton.module.scss';

interface MainButtonProps {
  label: string;
  onClick: () => void;
}

const MainButton: React.FC<MainButtonProps> = ({ label, onClick }) => {
  return (
    <div className={c.btn}>
      <Button variant="contained" onClick={onClick}>
        {label}
      </Button>
    </div>
  );
};

export default MainButton;
