import React from 'react';
import Slider from '@mui/material/Slider';

interface SliderProps {
  min: number;
  max: number;
  step: number;
}

const SliderComponent: React.FC<SliderProps> = ({ min, max, step }) => {
  const defaultValue = (max - min) / 2 + min; // Set default value to the midpoint
  
  return (
    <Slider
      defaultValue={defaultValue}
      min={min}
      max={max}
      step={step}
      marks
      valueLabelDisplay='auto'
      aria-label='slider'
      sx={{
        width: 1000,
        color: '#DBF881'
      }}
    />
  );
};

export default SliderComponent;
