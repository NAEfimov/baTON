import React, { useState } from 'react';
import '../css/Home.css';
import { candidatePayload } from '../data/candidate.js';
import CandidatureProfile from './CandidatureProfile';
import Loading from '../components/LoadingScreen';

export default function Candidate() {
  const [isLoading, setIsLoading] = useState(false);
  const [fetchedProfile, setFetchedProfile] = useState(null);
  const [error, setError] = useState(null);

  const handleUploadResume = async () => {
    setIsLoading(true);
    setError(null);

    try {
      const postResponse = await fetch('/backend/candidates', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(candidatePayload),
      });
      if (!postResponse.ok) throw new Error(`POST failed: ${postResponse.status}`);

      const idToFetch = candidatePayload.telegram_id;
      const getResponse = await fetch(`/backend/candidates?telegram_id=${idToFetch}`);
      if (!getResponse.ok) throw new Error(`GET failed: ${getResponse.status}`);

      const getResult = await getResponse.json();

      const completeProfile = {
        ...getResult,
        id: candidatePayload.telegram_id,
        photoUrl: `https://t.me/i/userpic/320/${candidatePayload.username}.jpg`,
        telegramUsername: candidatePayload.username
      };      
      setFetchedProfile(completeProfile);

    } catch (error) {
      console.error("An error occurred:", error);
      setError(error.message);
    } finally {
      setIsLoading(false);
    }
  };

  if (isLoading) {
    return <Loading message="Creating profile..." />;
  }

  if (error) {
    return <div className="candidate-container">Error: {error} <GoBack /></div>;
  }

  if (fetchedProfile) {
    return (
      <CandidatureProfile
        candidate={fetchedProfile}
        isFirst={true}
        isLast={true}
        onNext={() => {}}
        onPrevious={() => {}}
        matchingScore={fetchedProfile.matching_score || 85}
        yearsExperience={fetchedProfile.years}
        onInfo={() => alert('Info button clicked!')}
        onShare={() => alert('Share button clicked!')}
      />
    );
  }

  return (
    <div className="app-container">
      <div className="card">
        <button className="button" onClick={handleUploadResume}>
          Upload Resume
        </button>
      <div/>
    </div>
    </div>
  );
}