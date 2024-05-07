import IconButton from '@mui/material/IconButton';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import React from 'react';
import { Link } from "react-router-dom";
import c from './PageHeader.module.scss';

interface PageHeaderProps {
  heading: string;
  link: string;
}

const PageHeader: React.FC<PageHeaderProps> = ({ heading, link }) => {
  return (
    <div className={c.container}>
        <IconButton component={Link} to={link}>
            <ArrowBackIcon  className={c.backButton} />
        </IconButton>
        <p className={c.heading}>{heading}</p>
    </div>
  );
};

export default PageHeader;