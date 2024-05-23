import GenerateButton from "../Components/GenerateButton/GenerateButton";
import InputField from "../Components/InputField";
import MultipleChoiceSelector from "../Components/MultipleChoiceSelector";
import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";
import RangeSelector from "../Components/RangeSelector";
import './template.scss';
import {useEffect} from "react";

const RoomTemplate = () => {
    useEffect(() => {
        const token = localStorage.getItem('token');
        if (!token) {
            window.location.href = '/login';
        }

        const checkToken = async () => {
            try {
                const response = await fetch(`${process.env.REACT_APP_BACKEND_URL}/check-token-valid`, { headers: { 'Authorization': `Bearer ${token}` } });
                if (!response.ok) {
                    if (response.status === 401) {
                        localStorage.removeItem('token');
                        window.location.href = '/login';
                        return;
                    }
                }
            } catch (error) {
                console.error('Error checking token:', error);
            }
        };

        checkToken();
    }, []);
    return (
        <>
            <PageHeader heading="Create Room Poll" link="/selectTemplate"/>
            <div className='template'>
            <PollHeader />
            <p className='question'>How do you like the space overall?</p>
                <RangeSelector min={0} max={10} step={1} onChange={function (value: number): void { }} /> <br />
                <p className='question'>What would you change about the space if you could?</p>
                <InputField label={"Changes"} placeholder={"I would change..."} onChange={function (value: string): void { }} />
                <p className='question'>How do you like the furnishings and decorations of the room?</p>
                <RangeSelector min={0} max={10} step={1} onChange={function (value: number): void { }} /> <br />
                <p className='question'>How would you describe the room?</p>
                <MultipleChoiceSelector options={['Stylish', 'Cozy', 'Sterile', 'Spacious', 'Overwhelming', 'Ugly']} onChange={function (option: string): void { }} />
                <p className='question'>What piece of furniture is missing in the room?</p>
                <InputField label={"Furniture"} placeholder={"I would like to have..."} onChange={function (value: string): void { }} />
                <p className='heading'>
                    4. Everything Correct? Then Generate Your Poll!
                </p>
                <div className='generateButton'>
                    <GenerateButton label={""} onClick={function (): void {} } />
                </div>
            </div>
        </>
     );
  };
  
  export default RoomTemplate;