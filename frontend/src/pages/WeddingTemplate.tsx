import PageHeader from "../Components/PageHeader/PageHeader";
import PollHeader from "../Components/PollHeader/PollHeader";

const WeddingTemplate = () => {
    return (
        <>
            <PageHeader heading="Create Wedding Pool" link="/selectTemplate"/>
            <PollHeader></PollHeader>
            <p>Frage 1</p>
        </>
     );
  };
  
  export default WeddingTemplate;