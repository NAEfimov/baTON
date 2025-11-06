import React from 'react';
import { Outlet } from 'react-router-dom';
import GoBack from './GoBack';

export default function AppLayout() {
  
  return (
    <>
      <Outlet />
      <GoBack />
    </>
  );
}