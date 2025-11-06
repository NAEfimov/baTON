import { Route, Routes } from 'react-router-dom';
import Home from './screens/Home';
import Candidate from './screens/Candidate';
import Recruiter from './screens/Recruiter';
import AppLayout from './components/AppLayout';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import './css/index.css';

export default function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route element={<AppLayout />}>
          <Route path="/recruiter" element={<Recruiter />} />
          <Route path="/candidate" element={<Candidate />} />
        </Route>
      </Routes>
      <ToastContainer />
    </div>
  );
}