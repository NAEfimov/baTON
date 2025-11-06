import { Link } from 'react-router-dom';
import '../css/Home.css';

export default function Home() {
  return (
    <div className="app-container">
      <h1>Choose Your Role</h1>
      <p>Select your profile type to get started.</p>

      <div className="card">
        <div className="button-group">
          <Link to="/candidate" className="button">
            I am a Candidate
          </Link>

          <Link to="/recruiter" className="button">
            I am a Recruiter
          </Link>
        </div>
      </div>
    </div>
  );
}