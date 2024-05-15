import GenerateButton from "../Components/GenerateButton/GenerateButton";
import InputField from "../Components/InputField";
import MultipleChoiceSelector from "../Components/MultipleChoiceSelector";
import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";
import RangeSelector from "../Components/RangeSelector";
import './template.scss';

const WeddingTemplate = () => {
    return (
        <>
            <PageHeader heading="Create Wedding Poll" link="/selectTemplate"/>
            <div className='template'>
                <PollHeader />
                <p className='question'>Who did you invite to the wedding?</p>
                <MultipleChoiceSelector options={['Bride', 'Groom', 'Both']} onChange={function (option: string): void {} } />
                <p className='question'>How long have you known the bride and groom?</p>
                <RangeSelector min={0} max={30} step={1} onChange={function (value: number): void {} } />
                <p className='question'>How do you know the bride and groom?</p>
                <InputField label={"History"} placeholder={"I know you..."} onChange={function (value: string): void {} } />
                <p className='question'>What was your highlight of the wedding?</p>
                <MultipleChoiceSelector options={['Wedding Ceremony', 'Food', 'Wedding dance', 'Program', 'After Party']} onChange={function (option: string): void {} } />
                <p className='question'>What do you wish the bride and groom?</p>
                <InputField label={"Wishes"} placeholder={"I wish you..."} onChange={function (value: string): void {} } />
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
  
  export default WeddingTemplate;