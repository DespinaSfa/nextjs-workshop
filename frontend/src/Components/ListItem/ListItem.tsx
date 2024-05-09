import IconButton from '@mui/material/IconButton';
import ArrowForwardIcon from '@mui/icons-material/ArrowForward';
import "./ListItem.scss"

const ListItem = ({ title, description } : { title: string,  description: string}) => {
    return (
        <div className="list-item">
        <div className="text">
        <p className="list-title">{title}</p>
        { /* <p className="poll-description">-</p>
        <p className="poll-description">{description}</p> */}
       </div>
        <IconButton aria-label="delete" size="large" sx={{ color: "#ffffff" }}>
          <ArrowForwardIcon />
        </IconButton>
      </div>
    )
}

export default ListItem;