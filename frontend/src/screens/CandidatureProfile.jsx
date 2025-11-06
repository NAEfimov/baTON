import React, { useState, useEffect } from "react";
import "../css/CandidatureProfile.css";
import GoBack from "../components/GoBack";
import ExperienceCarousel from "../components/ExperienceCarousel";
import ProfileActions from "../components/ProfileActions";
import InfoPopup from "../components/InfoPopup";
import { PrevIcon, NextIcon, TelegramIcon, VerifiedIcon } from "../components/Icons";

export default function CandidatureProfile({ candidate, onNext, onPrevious, isFirst, isLast, matchingScore, yearsExperience, onSuperLike, onInfo, onShare, }) {
  const [[jobIndex, direction], setJobIndex] = useState([0, 0]);
  const [popupContent, setPopupContent] = useState(null);
  const jobs = candidate.experience;

  const handleShowLocation = () => {
    setPopupContent({
      title: "Location",
      content: candidate.location || "Location not provided."
    });
  };

  const handleShowEducation = () => {
    setPopupContent({
      title: "Education",
      content: candidate.education || "Education details not provided."
    });
  };
  
  const handleClosePopup = () => {
    setPopupContent(null);
  };

  useEffect(() => {
    setJobIndex([0, 0]);
  }, [candidate.id]);


  const paginate = (newDirection) => {
    setJobIndex(prev => {
        const newIndex = prev[0] + newDirection;
        if (newIndex < 0 || newIndex >= jobs.length) return prev;
        return [newIndex, newDirection];
    });
  };

  const handleDragEnd = (e, { offset, velocity }) => {
    const swipe = Math.abs(offset.x) * velocity.x;
    if (swipe < -10000) {
      paginate(1);
    } else if (swipe > 10000) {
      paginate(-1);
    }
  };
  return (
    <div className="profile-card">
      <div className="profile-card-image-container">
        <img src={candidate.photoUrl} alt={candidate.name} className="profile-card-image"/>
      </div>
      <div className="profile-header">
      <h2 className="profile-card-name">{candidate.name}</h2>
        <ProfileActions
          onShowLocation={handleShowLocation}
          onShowEducation={handleShowEducation}
          matchingScore={matchingScore}
          yearsExperience={yearsExperience}
        />
      </div>
      <div className="profile-card-content">
        
        <section className="profile-card-section">
          <h3 className="section-title">Experience ({jobIndex + 1} of {jobs.length})</h3>
          <ExperienceCarousel
            jobs={jobs}
            jobIndex={jobIndex}
            direction={direction}
            onPaginate={paginate}
            onDragEnd={handleDragEnd}
          />
          <div className="pagination-dots">
            {jobs.map((_, i) => (
              <button
                key={i}
                className={`dot ${i === jobIndex ? 'active' : ''}`}
                onClick={() => setJobIndex([i, i > jobIndex ? 1 : -1])}
              />
            ))}
          </div>
        </section>
        <div className="profile-card-navigation">
          <button className="nav-button" onClick={onPrevious} disabled={isFirst}>
            <PrevIcon />
          </button>
          
          <a
            href={`https://t.me/${candidate.telegramUsername}`}
            target="_blank"
            rel="noopener noreferrer"
            className="nav-button nav-button--telegram"
          >
            <TelegramIcon />
          </a>

          <button className="nav-button" onClick={onNext} disabled={isLast}>
            <NextIcon />
          </button>
        </div>
      </div>
      <GoBack />
      {popupContent && (
        <InfoPopup 
          title={popupContent.title}
          content={popupContent.content}
          onClose={handleClosePopup}
        />
      )}
    </div>
  );
}