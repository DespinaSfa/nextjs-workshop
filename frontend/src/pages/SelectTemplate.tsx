import { Outlet, Link } from "react-router-dom";
import c from './Page_styles.module.scss';
import Button from "@mui/material/Button";
import PageHeader from "../Components/PageHeader/PageHeader";

const SelectTemplate = () => {
    return(
        <div className={c.content}>
            <PageHeader heading="Select poll type" link="/"/>
            <div className={c.selectContainer}>
                <Button className={`${c.selectButton} ${c.party}`} variant="contained" component={Link} to='partyTemplate'>
                    <div className={c.buttonText}>Party</div>
                </Button>
                <Button className={`${c.selectButton} ${c.room}`} variant="contained" component={Link} to='roomTemplate'>
                    <div className={c.buttonText}>Room</div>
                </Button>
                <Button className={`${c.selectButton} ${c.wedding}`} variant="contained" component={Link} to='weddingTemplate'>
                        
                    <div className={c.buttonText}>Wedding</div>
                </Button>
            </div>

            <Outlet />
        </div>
    )
  };
  
  export default SelectTemplate;