import { Route, Routes } from 'react-router-dom';
import LoginPage from '@/app/(auth)/layout';
import LanguagePage from '@/app/(app)/lang/page';
import * as React from 'react';
import HomePage from '@/app/(app)/home/page';
import { AccountsPage, AccountsPageEdit, ProfilePagePasswordChange } from '@/app/(app)/accounts/page';
import { RoutesLocation } from '@/components/routes';
import { AuthProvidersPage, AuthProvidersPageEdit } from '@/app/(app)/auth-providers/page';
import { ServersPage, ServersPageEdit } from '@/app/(app)/servers/page';
import { LTIRoutingPage, LTIRoutingPageEdit } from '@/app/(app)/lti-routings/page';
import {
  LTIAttemptPageConfirm,
  LTIAttemptsPageCreate,
  LTIAttemptsPageEdit,
  LTIAttemptsPage
} from '@/app/(app)/lti-attempts/page';
import { ServiceCardsPage, ServiceCardsPageEdit } from '@/app/(app)/service-cards/page';
import { RolesListPage, RolesPageEdit } from '@/app/(app)/roles/page';
import { TopologyDevicePageView, TopologyPageConnect, TopologyPageView } from '@/app/(app)/topology/page';
import { DocsKubectlTopology } from '@/app/(docs)/page';
import { TargetsPage, TargetsPageEdit } from '@/app/(app)/targets/page';

const RoutesAdmin = () => {
  return (
    <Routes>
      <Route path={RoutesLocation.accounts()} element={<AccountsPage />} />
      <Route path={RoutesLocation.accountsEdit()} element={<AccountsPageEdit />} />
      <Route path={RoutesLocation.accountsCreate()} element={<AccountsPageEdit />} />
      <Route path={RoutesLocation.profileChangePassword()} element={<ProfilePagePasswordChange />} />
      <Route path={RoutesLocation.roles()} element={<RolesListPage />} />
      <Route path={RoutesLocation.rolesEdit()} element={<RolesPageEdit />} />
      <Route path={RoutesLocation.rolesCreate()} element={<RolesPageEdit />} />
      <Route path={RoutesLocation.authProviders()} element={<AuthProvidersPage />} />
      <Route path={RoutesLocation.authProvidersEdit()} element={<AuthProvidersPageEdit />} />
      <Route path={RoutesLocation.authProvidersCreate()} element={<AuthProvidersPageEdit />} />
      <Route path={RoutesLocation.ltiRouting()} element={<LTIRoutingPage />} />
      <Route path={RoutesLocation.ltiRoutingEdit()} element={<LTIRoutingPageEdit />} />
      <Route path={RoutesLocation.ltiRoutingCreate()} element={<LTIRoutingPageEdit />} />
      <Route path={RoutesLocation.ltiRedirect()} element={<LTIAttemptPageConfirm />} />
      <Route path={RoutesLocation.ltiRedirectCreate()} element={<LTIAttemptsPageCreate />} />
      <Route path={RoutesLocation.ltiAttemptEdit()} element={<LTIAttemptsPageEdit />} />
      <Route path={RoutesLocation.ltiAttempts()} element={<LTIAttemptsPage />} />
      <Route path={RoutesLocation.servers()} element={<ServersPage />} />
      <Route path={RoutesLocation.serversEdit()} element={<ServersPageEdit />} />
      <Route path={RoutesLocation.serversCreate()} element={<ServersPageEdit />} />
      <Route path={RoutesLocation.serviceCards()} element={<ServiceCardsPage />} />
      <Route path={RoutesLocation.serviceCardsEdit()} element={<ServiceCardsPageEdit />} />
      <Route path={RoutesLocation.serviceCardsCreate()} element={<ServiceCardsPageEdit />} />
      <Route path={RoutesLocation.targets()} element={<TargetsPage />} />
      <Route path={RoutesLocation.targetsEdit()} element={<TargetsPageEdit />} />
      <Route path={RoutesLocation.targetsCreate()} element={<TargetsPageEdit />} />
      <Route path={RoutesLocation.topologyView()} element={<TopologyPageView />} />
      <Route path={RoutesLocation.topologyDevices()} element={<TopologyDevicePageView />} />
      <Route path={RoutesLocation.topologyConnect()} element={<TopologyPageConnect />} />
      <Route path={RoutesLocation.docsTopologyKubectl()} element={<DocsKubectlTopology />} />
      <Route path={RoutesLocation.login()} element={<LoginPage />} />
      <Route path={RoutesLocation.language()} element={<LanguagePage />} />
      <Route path={RoutesLocation.home()} element={<HomePage />} />
    </Routes>
  );
};

export default RoutesAdmin;
