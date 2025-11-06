import React from 'react';
import { VerifiedIcon } from './Icons';

export default function ExperienceCard({ job }) {
  return (
    <div className="experience-card">
      <div className="experience-card-header">
      <p className="job-title">{job.company} | {job.vacancy}</p>
      <div className={`verification-badge ${job.verified ? 'is-verified' : 'is-not-verified'}`}>
          <VerifiedIcon />
        </div>
      </div>
      <ul className="job-highlights">
        {job.highlights.map((highlight, index) => (
          <li key={index}>{highlight}</li>
        ))}
      </ul>
    </div>
  );
}