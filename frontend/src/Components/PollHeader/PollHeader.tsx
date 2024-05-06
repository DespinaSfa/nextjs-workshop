import React from 'react';
import InputField from '../InputField';
import c from './PollHeader.module.scss';

const PollHeader = () => {
    return (
       <div>
         <p className={c.heading}>
            1. Select A Fancy Name For Your Poll
         </p>
         <InputField label={'Heading'} placeholder={'Name of your poll'} onChange={function (value: string): void {
            //TODO: Save Inputs
         } } />
          <p className={c.heading}>
            2. Write A Nice Description
         </p>
         <InputField label={'Description'} placeholder={'This poll is about...'} onChange={function (value: string): void {
            //TODO: Save Inputs
         } } />
          <p className={c.heading}>
            3. Check The Poll
         </p>
       </div> 
    );
  };
  
  export default PollHeader;