// src/screens/Recruiter.jsx
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
      // 1. POST the vacancy to the backend
      const postResponse = await fetch('/backend/vacancies', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(vacancyPayload),
      });
      if (!postResponse.ok) throw new Error(`POST /vacancies failed: ${postResponse.status}`);
      console.log("POST /vacancies successful.");

      // 2. GET the matched candidates using the recruiter's ID
      const idToFetch = vacancyPayload.telegram_id;
      const getResponse = await fetch(`/backend/vacancies/matches?telegram_id=${idToFetch}`);
      if (!getResponse.ok) throw new Error(`GET /vacancies/matches failed: ${getResponse.status}`);

      const fetchedCandidates = await getResponse.json();
      console.log("GET /vacancies/matches successful. Fetched profiles:", fetchedCandidates);

      // 3. Update state with the fetched candidates
      if (fetchedCandidates && fetchedCandidates.length > 0) {
        setMatchedCandidates(fetchedCandidates);
        setCurrentIndex(0); // Reset carousel to the first profile
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

  const handleDislike = (e) => {
    e.stopPropagation();
    const currentCandidate = matchedCandidates[currentIndex];
    console.log("Disliked:", currentCandidate.name);
  };
  
  const handleSuperLike = (e) => {
    e.stopPropagation();
    const currentCandidate = matchedCandidates[currentIndex];
    console.log("Super-liked:", currentCandidate.name);
  };

  const handleLike = (e) => {
    e.stopPropagation();
    const currentCandidate = matchedCandidates[currentIndex];
    console.log("Liked:", currentCandidate.name);
  };

  const handleInfo = (e) => {
    e.stopPropagation();
    const currentCandidate = matchedCandidates[currentIndex];
    alert(`More info about ${currentCandidate.name}`);
  };

  const handleShare = (e) => {
    e.stopPropagation();
    const currentCandidate = matchedCandidates[currentIndex];
    alert(`Share ${currentCandidate.name}'s profile`);
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
        onDislike={handleDislike}
        onSuperLike={handleSuperLike}
        onLike={handleLike}
        onInfo={handleInfo}
        onShare={handleShare}
      />
    </div>
  );
}