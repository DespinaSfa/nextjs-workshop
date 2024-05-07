import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";

const WeddingTemplate = () => {
    return (
        <>
            <PageHeader heading="Create Wedding Pool" link="/selectTemplate"/>
            <PollHeader></PollHeader>
            <p> Und  hier kommt dann das Template hin... Fragen & Antwortoptionen siehe Notion</p>
        </>
     );
  };
  
  export default WeddingTemplate;