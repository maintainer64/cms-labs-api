import { Route, Routes } from 'react-router-dom';
import LanguagePage from '@/app/(app)/lang/page';
import * as React from 'react';
import LoginPage, { LoginError } from '@/app/(auth)/layout';
import { RoutesLocation } from '@/components/routes';
import { LTIAttemptsPageCreate } from '@/app/(app)/lti-attempts/page';
import TopologyPageView from '@/app/(app)/topology/page';
import { TargetsPage, TargetsPageEdit } from '@/app/(app)/targets/page';
import SessionPage from '@/app/(app)/session/page';

const RoutesStudent = () => {
  return (
    <Routes>
      <Route path={RoutesLocation.language()} element={<LanguagePage />} />
      <Route path={RoutesLocation.login()} element={<LoginPage />} />
      <Route path={RoutesLocation.home()} element={<LoginError />} />
      <Route path={RoutesLocation.targets()} element={<TargetsPage />} />
      <Route path={RoutesLocation.targetsEdit()} element={<TargetsPageEdit />} />
      <Route path={RoutesLocation.targetsCreate()} element={<TargetsPageEdit />} />
      <Route path={RoutesLocation.ltiRedirect()} element={<LTIAttemptsPageCreate />} />
      <Route path={RoutesLocation.ltiRedirectCreate()} element={<LTIAttemptsPageCreate />} />
      <Route path={RoutesLocation.topologyView()} element={<TopologyPageView />} />
      <Route path={RoutesLocation.session()} element={<SessionPage />} />
      <Route path={RoutesLocation.sessionTopology()} element={<TopologyPageView />} />
    </Routes>
  );
};

export default RoutesStudent;
