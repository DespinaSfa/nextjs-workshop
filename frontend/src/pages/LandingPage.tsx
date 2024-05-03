import { Outlet } from "react-router-dom";
import c from './Page_styles.module.scss';
import MainButton from "../Components/MainButton/MainButton";

const LandingPage = () => {
    return(
        <div className={c.landingPage}>
            <h1 className={c.icon}>Party Poll</h1>
            <p className={c.text}>Tired of boring parties? Spice things up with PartyPoll! Vote on songs, review the restroom, you name it! Plus, no need for sober sign-ups - just scan the QR code and vote instantly!</p>
            <div className={c.button}>
                <MainButton text="Create Poll" link="/selectTemplate"/>
            </div>

            <Outlet />
        </div>
    )
  };
  
  export default LandingPage;
  