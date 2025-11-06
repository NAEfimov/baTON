import React, { useState } from 'react';
import GoBack from '../components/GoBack';
import '../css/Home.css';
import { candidatePayload } from '../data/candidate.js';
import CandidatureProfile from './CandidatureProfile';

export default function Candidate() {
  const [isLoading, setIsLoading] = useState(false);
  const [fetchedProfile, setFetchedProfile] = useState(null);
  const [error, setError] = useState(null);

  const handleUploadResume = async () => {
    setIsLoading(true);
    setError(null);

    try {
      // 1. POST request using the full payload
      const postResponse = await fetch('/backend/candidates', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(candidatePayload),
      });
      if (!postResponse.ok) throw new Error(`POST failed: ${postResponse.status}`);
      console.log("POST successful.");

      // 2. GET request using the telegram_id from our source data
      const idToFetch = candidatePayload.telegram_id;
      const getResponse = await fetch(`/backend/candidates?telegram_id=${idToFetch}`);
      if (!getResponse.ok) throw new Error(`GET failed: ${getResponse.status}`);

      const getResult = await getResponse.json(); // This is the CandidatePublicDTO
      console.log("GET successful. Received from API:", getResult);

      // 3. *** THE CRITICAL STEP: Augment the data ***
      // Your API returns a DTO, but your component needs more (photoUrl, etc.).
      // We'll create a new object that combines the API data with our mock data.
      const completeProfile = {
        ...getResult, // Spreads all fields from the DTO (Name, Experience, Years, etc.)
        id: candidatePayload.telegram_id, // The component's useEffect key needs an 'id'
        photoUrl: `https://i.pravatar.cc/400?u=${candidatePayload.telegram_id}`, // Add a mock photo
        telegramUsername: candidatePayload.username // Add the username from our source payload
      };
      
      console.log("Augmented profile for UI:", completeProfile);
      setFetchedProfile(completeProfile);

    } catch (error) {
      console.error("An error occurred:", error);
      setError(error.message);
    } finally {
      setIsLoading(false);
    }
  };

  // --- UI Rendering Logic ---

  if (isLoading) {
    return <div className="candidate-container">Creating and fetching profile...</div>;
  }

  if (error) {
    return <div className="candidate-container">Error: {error} <GoBack /></div>;
  }

  if (fetchedProfile) {
    // If we have a profile, render the CandidatureProfile component
    return (
      <CandidatureProfile
        candidate={fetchedProfile}
        isFirst={true}
        isLast={true}
        onNext={() => {}}
        onPrevious={() => {}}
        matchingScore={fetchedProfile.matching_score || 85} // Use score from API or fallback
        yearsExperience={fetchedProfile.years} // Use years from API
        onInfo={() => alert('Info button clicked!')}
        onShare={() => alert('Share button clicked!')}
      />
    );
  }

  // The initial view: show the button
  return (
    <div className="app-container">
      <div className="card">
        <button className="button" onClick={handleUploadResume}>
          Upload Resume
        </button>
      <div/>
    </div>
      <GoBack />
    </div>
  );
}