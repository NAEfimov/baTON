import React, { useState } from "react";
import CandidatureProfile from "./CandidatureProfile";
import Loading from "../components/LoadingScreen";
import { vacancyPayload } from "../data/vacancy.js";

export default function Recruiter() {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [matchedCandidates, setMatchedCandidates] = useState([]);
  const [currentIndex, setCurrentIndex] = useState(0);

  const handleFindMatches = async () => {
    setIsLoading(true);
    setError(null);

    try {
      const postResponse = await fetch('/backend/vacancies', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(vacancyPayload),
      });
      if (!postResponse.ok) throw new Error(`POST /vacancies failed: ${postResponse.status}`);
      const idToFetch = vacancyPayload.telegram_id;
      const getResponse = await fetch(`/backend/vacancies/matches?telegram_id=${idToFetch}`);
      if (!getResponse.ok) throw new Error(`GET /vacancies/matches failed: ${getResponse.status}`);

      const fetchedCandidates = await getResponse.json();
      if (fetchedCandidates && fetchedCandidates.length > 0) {
        const candidatesWithPhotos = fetchedCandidates.map(candidate => {
          const uniqueIdentifier = candidate.username;
          return {
            ...candidate,
            photoUrl: `https://t.me/i/userpic/320/${uniqueIdentifier}.jpg`
          };
        });        
        setMatchedCandidates(candidatesWithPhotos);
        setCurrentIndex(0);
      } else {
        setError("No matching candidates found.");
      }

    } catch (err) {
      console.error("An error occurred:", err);
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleNext = () => {
    setCurrentIndex(prevIndex => Math.min(prevIndex + 1, matchedCandidates.length - 1));
  };

  const handlePrevious = () => {
    setCurrentIndex(prevIndex => Math.max(prevIndex - 1, 0));
  };
  
  const handleInfo = (e) => {
    e.stopPropagation();
    const currentCandidate = matchedCandidates[currentIndex];
    alert(`More info about ${currentCandidate.name}`);
  };

  if (isLoading) {
    return <Loading message="Finding matches..." />;
  }
  if (error) {
    return <div className="app-container">Error: {error}</div>;
  }
  if (matchedCandidates.length === 0) {
    return (
      <div className="app-container">
        <div className="card">
          <h1>Find Candidates</h1>
          <p>Upload your vacancy to find the best matching talent.</p>
          <button className="button" onClick={handleFindMatches}>
            Find Matches
          </button>
        </div>
      </div>
    );
  }
  const currentCandidate = matchedCandidates[currentIndex];
  const isFirst = currentIndex === 0;
  const isLast = currentIndex === matchedCandidates.length - 1;

  return (
    <div>
      <CandidatureProfile 
        key={currentCandidate.id} 
        candidate={currentCandidate}
        onNext={handleNext}
        onPrevious={handlePrevious}
        isFirst={isFirst}
        isLast={isLast}
        matchingScore={currentCandidate.matching_score}
        yearsExperience={currentCandidate.years}
        onInfo={handleInfo}
      />
    </div>
  );
}