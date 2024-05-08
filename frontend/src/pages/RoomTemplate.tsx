import GenerateButton from "../Components/GenerateButton/GenerateButton";
import InputField from "../Components/InputField";
import MultipleChoiceSelector from "../Components/MultipleChoiceSelector";
import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";
import RangeSelector from "../Components/RangeSelector";
import c from './Page_styles.module.scss';

const RoomTemplate = () => {
    return (
        <>
            <PageHeader heading="Create Room Poll" link="/selectTemplate"/>
            <div className={c.template}>
            <p className={c.question}>How do you like the space overall?</p>
                <RangeSelector min={0} max={10} step={1} onChange={function (value: number): void { }} />
                <p className={c.question}>What would you change about the space if you could?</p>
                <InputField label={"Changes"} placeholder={"I would change..."} onChange={function (value: string): void { }} />
                <p className={c.question}>How do you like the furnishings and decorations of the room?</p>
                <RangeSelector min={0} max={10} step={1} onChange={function (value: number): void { }} />
                <p className={c.question}>How would you describe the room?</p>
                <MultipleChoiceSelector options={['Stylish', 'Cozy', 'Sterile', 'Spacious', 'Overwhelming', 'Ugly']} onChange={function (option: string): void { }} />
                <p className={c.question}>What piece of furniture is missing in the room?</p>
                <InputField label={"Furniture"} placeholder={"I would like to have..."} onChange={function (value: string): void { }} />
                <p className={c.heading}>
                    4. Everything Correct? Then Generate Your Poll!
                </p>
                <div className={c.generateButton}>
                    <GenerateButton label={""} onClick={function (): void {} } />
                </div>
            </div>
        </>
     );
  };
  
  export default RoomTemplate;