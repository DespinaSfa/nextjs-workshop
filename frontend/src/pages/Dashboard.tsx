import "./dashboard.scss"
import ListItem from '../Components/ListItem/ListItem';
import { useEffect, useState } from "react";

interface Poll {
  title: string;
  description: string;
  pollType: string;
}

const Dashboard = () => {
  const [polls, setPolls] = useState<Poll[]>([]);

  useEffect(() => {
    const fetchPolls = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_BACKEND_URL}/polls`);
        if (!response.ok) {
          throw new Error('Failed to fetch polls');
        }
        const data = await response.json();
        setPolls(data);
      } catch (error) {
        console.error('Error fetching polls:', error);
      }
    };

    fetchPolls();
  }, []); 

    return (
      <>
        <h2 className="dashboard-title">Your Polls</h2>
        <div className="poll-list">
          {polls.map((poll, index) => (
            <ListItem key={index} title={poll.title} description={poll.description} />
          ))}
        </div>
      </>
    );
  };
  
  export default Dashboard;
  