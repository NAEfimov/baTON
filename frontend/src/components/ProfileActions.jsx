import React from "react";
import { MatchIcon, InfoIcon, ShareIcon, YearsIcon } from "./Icons";


export default function ProfileActions({ onInfo, onShare, matchingScore, yearsExperience }) {
  return (
    <div className="profile-card-actions">
      <div className="action-item action-badge">
        <span className="badge-value">{matchingScore}%</span>
        <MatchIcon />
      </div>
      <div className="action-item action-badge">
        <span className="badge-value">{yearsExperience}</span>
        <YearsIcon />
      </div>
      <button className="action-item action-button action-button--info" onClick={onInfo}>
        <InfoIcon />
      </button>
      <button className="action-item action-button action-button--share" onClick={onShare}>
        <ShareIcon />
      </button>
    </div>
  );
}