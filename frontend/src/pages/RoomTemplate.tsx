import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";

const RoomTemplate = () => {
    return (
        <>
            <PageHeader heading="Create Room Pool" link="/selectTemplate"/>
            <PollHeader></PollHeader>
            <p> Und  hier kommt dann das Template hin... Fragen & Antwortoptionen siehe Notion</p>
        </>
     );
  };
  
  export default RoomTemplate;