import React from 'react';

export default function ExperienceCard({ job }) {
  return (
    <div className="experience-card">
      <p className="job-title">{job.company} | {job.vacancy}</p>
      <ul className="job-highlights">
        {job.highlights.map((highlight, index) => (
          <li key={index}>{highlight}</li>
        ))}
      </ul>
    </div>
  );
}