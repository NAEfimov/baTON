import React from "react";
import { Link } from "react-router-dom";
import "../css/GoBack.css";

const BackArrowIcon = () => (
  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor">
    <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
  </svg>
);

export default function GoBack({}) {
  return (
    <Link to="/" className="go-back-link">
      <BackArrowIcon />
    </Link>
  );
}