import { Outlet } from "react-router-dom";
import Header from "../Components/Header/Header";
import styles from "./Page_styles.module.scss";

const Layout = () => {
    return (
        <div>
          {/* A "layout route" is a good place to put markup you want to
              share across all the pages on your site, like navigation. -> add header and footer here */}
          <Header />
    
          <div className={styles["page-margin"]}>
           <Outlet/>
          </div>
        </div>
      );
};
  
export default Layout;
