import { Outlet, Link } from "react-router-dom";
import Header from "../Components/Header/Header";

const Layout = () => {
    return (
        <div>
          {/* A "layout route" is a good place to put markup you want to
              share across all the pages on your site, like navigation. -> add header and footer here */}
          <Header />
    
          <Outlet />
        </div>
      );
};
  
export default Layout;
