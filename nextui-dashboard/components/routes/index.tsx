export const RoutesLocation = {
  login: () => '/login',
  language: () => '/lang',
  accounts: () => '/accounts',
  accountsEdit: (id = ':id') => `/accounts/edit/${id}`,
  accountsCreate: () => `/accounts/create`,
  ltiForms: () => '/lti-forms',
  ltiFormsEdit: (id = ':id') => `/lti-forms/edit/${id}`,
  ltiFormsCreate: () => `/lti-forms/create`,
  ltiRouting: () => '/lti-routing',
  ltiRoutingEdit: (id = ':id') => `/lti-routing/edit/${id}`,
  ltiRoutingCreate: () => `/lti-routing/create`,
  pnetServers: () => '/pnet-servers',
  pnetServersEdit: (id = ':id') => `/pnet-servers/edit/${id}`,
  pnetServersCreate: () => `/pnet-servers/create`,
  profileChangePassword: () => `/profile/password`,
  tasks: () => '/tasks',
  home: () => '/'
};
