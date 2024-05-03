import React from 'react';
import Button from '@mui/material/Button';
import c from './GenerateButton.module.scss';

interface GenerateButtonProps {
    label: string;
    onClick: () => void;
}

const GenerateButton: React.FC<GenerateButtonProps> = ({ label, onClick }) => {
    return (
        <Button className={c.generateButton} variant="contained" onClick={onClick}>
            {label}
        </Button>
    );
};

export default GenerateButton;