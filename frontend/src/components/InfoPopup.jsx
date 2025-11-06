import React from 'react';
import '../css/InfoPopup.css';

export default function InfoPopup({ title, content, onClose }) {
  if (!content) {
    return null;
  }

  const handleModalClick = (e) => e.stopPropagation();

  return (
    <div className="popup-overlay" onClick={onClose}>
      <div className="popup-modal" onClick={handleModalClick}>
        <h3 className="popup-title">{title}</h3>
        <p className="popup-content">{content}</p>
        <button className="popup-close-btn" onClick={onClose}>
          Close
        </button>
      </div>
    </div>
  );
}