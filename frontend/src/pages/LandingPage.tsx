import { Outlet, Link } from "react-router-dom";

const LandingPage = () => {
    return(
        <>
        <h1>Landing Page</h1>
        <p>Gehe zum Dashboard:</p>
        <button>
            <Link to="/dashboard">Dashboard</Link>
        </button>
        <p>Gehe zum Login:</p>
        <button>
            <Link to="/login">Login</Link>
        </button>

        <Outlet />
        </>
    )
  };
  
  export default LandingPage;
  