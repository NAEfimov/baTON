// src/components/Button.jsx

import React from 'react';
import { Link } from 'react-router-dom';

function Button({ to, children, className }) {
  return (
    <Link to={to} className={className}>
      {children}
    </Link>
  );
}

export default Button;