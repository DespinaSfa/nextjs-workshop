import GenerateButton from "../Components/GenerateButton/GenerateButton";
import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";
import c from './Page_styles.module.scss';

const RoomTemplate = () => {
    return (
        <>
            <PageHeader heading="Create Room Pool" link="/selectTemplate"/>
            <div className={c.template}>
                <PollHeader></PollHeader>
                <p> Und  hier kommt dann das Template hin... Fragen & Antwortoptionen siehe Notion</p>
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