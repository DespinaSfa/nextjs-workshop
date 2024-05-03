import { Outlet, Link } from "react-router-dom";
import c from './Page_styles.module.scss';
import Button from "@mui/material/Button";
import { IconButton } from "@mui/material";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

const SelectTemplate = () => {
    return(
        <div className={c.content}>
            <div className={c.container}>
                <IconButton component={Link} to='/'>
                    <ArrowBackIcon  className={c.backButton} />
                </IconButton>
                <p className={c.heading}>Select poll type</p>
            </div>
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