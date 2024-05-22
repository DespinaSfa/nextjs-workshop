import { Outlet, Link } from "react-router-dom";
import './selectTemplate.scss';
import Button from "@mui/material/Button";
import PageHeader from "../Components/PageHeader/PageHeader";

const SelectTemplate = () => {
    return(
        <div className="content">
            <PageHeader heading="Select poll type" link="/"/>
            <div className="selectContainer">
                <Button className="selectButton party" variant="contained" component={Link} to='partyTemplate'>
                    <div className="buttonText">Party</div>
                </Button>
                <Button className="selectButton room" variant="contained" component={Link} to='roomTemplate'>
                    <div className="buttonText">Room</div>
                </Button>
                <Button className="selectButton wedding" variant="contained" component={Link} to='weddingTemplate'>
                        
                    <div className="buttonText">Wedding</div>
                </Button>
            </div>

            <Outlet />
        </div>
    )
  };
  
  export default SelectTemplate;