import { Route, Routes } from 'react-router-dom';
import LanguagePage from '@/app/(app)/lang/page';
import * as React from 'react';
import LoginPage, { LoginError } from '@/app/(auth)/layout';
import { RoutesLocation } from '@/components/routes';
import { LTIAttemptsPageCreate } from '@/app/(app)/lti-attempts/page';

const RoutesStudent = () => {
  return (
    <Routes>
      <Route path={RoutesLocation.language()} element={<LanguagePage />} />
      <Route path={RoutesLocation.login()} element={<LoginPage />} />
      <Route path={RoutesLocation.home()} element={<LoginError />} />
      <Route path={RoutesLocation.ltiRedirect()} element={<LTIAttemptsPageCreate />} />
      <Route path={RoutesLocation.ltiRedirectCreate()} element={<LTIAttemptsPageCreate />} />
    </Routes>
  );
};

export default RoutesStudent;
