import { Route, Routes } from 'react-router-dom';
import LoginPage from '@/app/(auth)/layout';
import LanguagePage from '@/app/(app)/lang/page';
import * as React from 'react';
import HomePage from '@/app/(app)/home/page';
import { AccountsPage, AccountsPageEdit, ProfilePagePasswordChange } from '@/app/(app)/accounts/page';
import { RoutesLocation } from '@/components/routes';
import { LTIFormsPage, LTIFormsPageEdit } from '@/app/(app)/lti-forms/page';
import { PnetServersPage, PnetServersPageEdit } from '@/app/(app)/pnet-servers/page';
import { LTIRoutingPage, LTIRoutingPageEdit } from '@/app/(app)/lti-routings/page';
import { LTIAttemptsPageCreate, LTIAttemptsPageEdit, LTIAttemptsPageUser } from '@/app/(app)/lti-attempts/page';
import { ServiceCardsPage, ServiceCardsPageEdit } from '@/app/(app)/service-cards/page';

const RoutesAdmin = () => {
  return (
    <Routes>
      <Route path={RoutesLocation.accounts()} element={<AccountsPage />} />
      <Route path={RoutesLocation.accountsEdit()} element={<AccountsPageEdit />} />
      <Route path={RoutesLocation.accountsCreate()} element={<AccountsPageEdit />} />
      <Route path={RoutesLocation.profileChangePassword()} element={<ProfilePagePasswordChange />} />
      <Route path={RoutesLocation.ltiForms()} element={<LTIFormsPage />} />
      <Route path={RoutesLocation.ltiFormsEdit()} element={<LTIFormsPageEdit />} />
      <Route path={RoutesLocation.ltiFormsCreate()} element={<LTIFormsPageEdit />} />
      <Route path={RoutesLocation.ltiRouting()} element={<LTIRoutingPage />} />
      <Route path={RoutesLocation.ltiRoutingEdit()} element={<LTIRoutingPageEdit />} />
      <Route path={RoutesLocation.ltiRoutingCreate()} element={<LTIRoutingPageEdit />} />
      <Route path={RoutesLocation.ltiRedirect()} element={<LTIAttemptsPageCreate />} />
      <Route path={RoutesLocation.ltiAttemptUser()} element={<LTIAttemptsPageUser />} />
      <Route path={RoutesLocation.ltiAttemptEdit()} element={<LTIAttemptsPageEdit />} />
      <Route path={RoutesLocation.pnetServers()} element={<PnetServersPage />} />
      <Route path={RoutesLocation.pnetServersEdit()} element={<PnetServersPageEdit />} />
      <Route path={RoutesLocation.pnetServersCreate()} element={<PnetServersPageEdit />} />
      <Route path={RoutesLocation.serviceCards()} element={<ServiceCardsPage />} />
      <Route path={RoutesLocation.serviceCardsEdit()} element={<ServiceCardsPageEdit />} />
      <Route path={RoutesLocation.serviceCardsCreate()} element={<ServiceCardsPageEdit />} />
      <Route path={RoutesLocation.login()} element={<LoginPage />} />
      <Route path={RoutesLocation.language()} element={<LanguagePage />} />
      <Route path={RoutesLocation.home()} element={<HomePage />} />
    </Routes>
  );
};

export default RoutesAdmin;
