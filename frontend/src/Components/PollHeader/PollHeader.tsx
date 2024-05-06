import React from 'react';
import InputField from '../InputField';

const PollHeader = () => {
    return (
       <div>
         <p>
            1. Select A Fancy Name For Your Poll
         </p>
         <InputField label={'Heading'} placeholder={'Name of your poll'} onChange={function (value: string): void {
            //TODO: Save Inputs
         } } />
          <p>
            2. Write A Nice Description
         </p>
         <InputField label={'Description'} placeholder={'This poll is about...'} onChange={function (value: string): void {
            //TODO: Save Inputs
         } } />
          <p>
            3. Check The Poll
         </p>
       </div> 
    );
  };
  
  export default PollHeader;