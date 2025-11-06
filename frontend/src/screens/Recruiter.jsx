// src/screens/Recruiter.jsx
import React, { useState, useEffect } from "react";
import CandidatureProfile from "./CandidatureProfile";
import GoBack from "../components/GoBack";

const candidateList = [
  {
    id: 213123124,
    name: "Alexandros Lappas",
    username: "AlexLuthor135",
    photoUrl: "https://i.pravatar.cc/400?u=AlexLuthor135",
    experience: [
      { company: "42 Wolfsburg", vacancy: "Pedago Manager", highlights: ["Led the pedagogical team.", "Built a certificate generation platform.", "Automated peer evaluation."] },
      { company: "42 Berlin", vacancy: "Pedago Intern", highlights: ["Prepared and facilitated bootcamp workshops.", "Developed automated monitoring tools.", "Built API solutions."] },
    ],
    matching_score: 92,
    years: 2
  },
  {
    id: 987654321,
    name: "Elena Volkova",
    username: "elena_volkova",
    photoUrl: "https://i.pravatar.cc/400?u=elena_volkova",
    experience: [
      { company: "Yandex", vacancy: "Frontend Developer", highlights: ["Developed core UI features for Yandex.Music.", "Improved page load times by 30%.", "Mentored junior developers."] },
      { company: "Acme Ltd.", vacancy: "Junior UI Engineer", highlights: ["Built reusable React components.", "Wrote unit and integration tests."] },
    ],
    matching_score: 93,
    years: 3
  },
  {
    id: 112233445,
    name: "Ben Carter",
    username: "ben_carter",
    photoUrl: "https://i.pravatar.cc/400?u=ben_carter",
    experience: [
      { company: "Google", vacancy: "Backend Engineer", highlights: ["Worked on the Google Cloud Storage API.", "Designed and implemented a new caching layer."] },
    ],
    matching_score: 94,
    years: 4
  }
];

export default function Recruiter() {
  const [candidates, setCandidates] = useState([]);
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    // Simulating an API call
    setCandidates(candidateList);
  }, []);

  const handleNext = () => {
    setCurrentIndex(prevIndex => Math.min(prevIndex + 1, candidates.length - 1));
  };

  const handlePrevious = () => {
    setCurrentIndex(prevIndex => Math.max(prevIndex - 1, 0));
  };

  const handleDislike = (e) => {
    e.stopPropagation();
    console.log("Disliked:", singleCandidate.name);
    setShowCard(false);
  };
  
  const handleSuperLike = (e) => {
    e.stopPropagation();
    console.log("Super-liked:", singleCandidate.name);
    // Add logic here
    setShowCard(false);
  };

  const handleLike = (e) => {
    e.stopPropagation();
    console.log("Liked:", singleCandidate.name);
    // Add logic here
    setShowCard(false);
  };

  const handleInfo = (e) => {
    e.stopPropagation();
    // Add logic here
    alert(`More info about ${singleCandidate.name}`);
  };

  const handleShare = (e) => {
    e.stopPropagation();
    // Add logic here
    alert(`Share ${singleCandidate.name}'s profile`);
  };

  if (candidates.length === 0) {
    return <div>Loading candidates...</div>;
  }

  const currentCandidate = candidates[currentIndex];
  const isFirst = currentIndex === 0;
  const isLast = currentIndex === candidates.length - 1;

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