import GenerateButton from "../Components/GenerateButton/GenerateButton";
import InputField from "../Components/InputField";
import MultipleChoiceSelector from "../Components/MultipleChoiceSelector";
import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";
import RangeSelector from "../Components/RangeSelector";
import './template.scss';
import {useEffect} from "react";

const PartyTemplate = () => {

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (!token) {
            window.location.href = '/login';
        }

        const checkToken = async () => {
            try {
                const response = await fetch(`${process.env.REACT_APP_BACKEND_URL}/check-token-validToken`, { headers: { 'Authorization': `Bearer ${token}` } });
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
            <PageHeader heading="Create Party Poll" link="/select-template" />
            <div className="template">
            <PollHeader />
                <p className="question">Which songs should definitely be played tonight? 📻</p>
                <InputField label={"Songs"} placeholder={"I would like to listen to..."} onChange={function (value: string): void { }} />
                <p className="question">What is your current alcohol level? 📈</p>
                <RangeSelector min={0} max={5} step={1} onChange={function (value: number): void { }} /> <br />
                <p className="question">What alcohol level have you set as your goal for today? 🍺</p>
                <RangeSelector min={0} max={5} step={1} onChange={function (value: number): void { }} /><br />
                <p className="question">What is your favortite party activity?</p>
                <MultipleChoiceSelector options={['Dancing 💃', 'Shout along to party hits or karaoke 🎤', 
                'PartyGames (Bierpong, Rage-Cage, etc.) 🍻 ', 'Chilling and chatting a bit outside with friends 🗨️']} onChange={function (option: string): void { }} />
                <p className="question">Which snacks or drinks would you like for the next party? 🍔</p>
                <InputField label={"Snack/Drink"} placeholder={"I would like to eat/drink..."} onChange={function (value: string): void { }} />
                <p className="heading">
                    4. Everything Correct? Then Generate Your Poll!
                </p>
                <div className="generateButton">
                    <GenerateButton label={""} onClick={function (): void { } } />
                </div>
            </div>
        </>
    );
  };
  
  export default PartyTemplate;