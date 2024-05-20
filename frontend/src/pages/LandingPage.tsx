import { Outlet } from "react-router-dom";
import './landingPage.scss';
import MainButton from "../Components/MainButton/MainButton";

const LandingPage = () => {
    return(
        <div className="landingPage">
            <h1 className="icon">Party Poll</h1>
            <p className="text">Tired of boring parties? Spice things up with PartyPoll! Vote on songs, review the restroom, you name it! Plus, no need for sober sign-ups - just scan the QR code and vote instantly!</p>
            <div className="button">
                <MainButton text="Login" link="/login"/>
            </div>

            <Outlet />
        </div>
    )
  };
  
  export default LandingPage;
  