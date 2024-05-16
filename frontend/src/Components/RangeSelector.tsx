import React from 'react';
import Slider from '@mui/material/Slider';

interface SliderProps {
  min: number;
  max: number;
  step: number;
  onChange: (value: number) => void;
}

const SliderComponent: React.FC<SliderProps> = ({ min, max, step, onChange }) => {
  const defaultValue = Math.ceil((max - min) / 2 + min);

  const handleSliderChange = (event: Event, value: number | number[]) => {
    if (!Array.isArray(value)) {
      onChange(value);
    }
  };

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
        width: '100%',
        color: '#DBF881'
      }}
      onChange={handleSliderChange} 
    />
  );
};

export default SliderComponent;
