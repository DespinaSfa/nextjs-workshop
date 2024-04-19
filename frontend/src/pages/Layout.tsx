import { Outlet, Link } from "react-router-dom";

const Layout = () => {
    return (
        <div>
          {/* A "layout route" is a good place to put markup you want to
              share across all the pages on your site, like navigation. -> add header and footer here */}
          <nav>
            <ul>
              <li>
                <Link to="/">LandingPage</Link>
              </li>
              <li>
                <Link to="/login">Login</Link>
              </li>
              <li>
                <Link to="/dashboard">Dashboard</Link>
              </li>
            </ul>
          </nav>
    
          <hr />
    
          <Outlet />
        </div>
      );
};
  
export default Layout;
