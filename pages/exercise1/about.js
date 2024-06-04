import Link from 'next/link';
import NavBar from '../../components/NavBar';

const ContactPage = () => {
    return (
      <div>
        <NavBar />
        <h1>Contact Page</h1>
        <p>Please fill out the form below:</p>
        <form>
          <input type="text" placeholder="Name" />
          <input type="email" placeholder="Email" />
          <textarea placeholder="Message"></textarea>
          <button type="submit">Submit</button>
        </form>
        <br />
      </div>
    );
  }
  
  export default ContactPage;
  