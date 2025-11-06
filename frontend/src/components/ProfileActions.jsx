import React from "react";
import { LocationIcon, EducationIcon, MatchIcon, YearsIcon } from "./Icons";


export default function ProfileActions({ onShowLocation, onShowEducation, matchingScore, yearsExperience }) {
  return (
    <div className="profile-card-actions">
      <div  className="action-item action-badge">
        <span>{matchingScore}%</span>
        <span className="badge-value">match</span>
        {/* <MatchIcon /> */}
      </div >
      <div className="action-item action-badge">
        <span>{yearsExperience}</span>
        <span className="badge-value">years</span>
        {/* <YearsIcon /> */}
      </div>
      <button className="action-item action-badge" onClick={onShowLocation}>
        <LocationIcon />
        <span className="badge-value">location</span>
      </button>
      <button className="action-item action-badge" onClick={onShowEducation}>
        <EducationIcon />
        <span className="badge-value">education</span>
      </button>
    </div>
  );
}