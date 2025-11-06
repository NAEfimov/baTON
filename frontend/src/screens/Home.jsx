import { Link } from 'react-router-dom';
import '../css/Home.css';
import batonImage from '../assets/baton_bw.png';

export default function Home() {
  return (
    <div className="app-container">
       <img src={batonImage} alt="A fresh loaf of baton bread" className="home-image" />
      <h1>Choose Your Role</h1>
      <p>Select your profile type to get started.</p>

      <div className="card">
        <div className="button-group">
          <Link to="/recruiter" className="button">
            I am a Recruiter
          </Link>
          <Link to="/candidate" className="button">
            I am a Candidate
          </Link>

        </div>
      </div>
    </div>
  );
}