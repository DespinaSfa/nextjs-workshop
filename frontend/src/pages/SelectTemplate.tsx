import { Outlet, Link } from "react-router-dom";
import './selectTemplate.scss';
import Button from "@mui/material/Button";
import PageHeader from "../Components/PageHeader/PageHeader";
import {useEffect} from "react";

const SelectTemplate = () => {
    useEffect(() => {
        const token = localStorage.getItem('token');
        if (!token) {
            window.location.href = '/login';
        }

        const checkToken = async () => {
            try {
                const response = await fetch(`${process.env.REACT_APP_BACKEND_URL}/check-token-valid`, {headers: { 'Authorization': `Bearer ${token}` } });
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

    return(
        <div className="content">
            <PageHeader heading="Select poll type" link="/dashboard"/>
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