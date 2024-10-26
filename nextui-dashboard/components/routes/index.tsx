export const RoutesLocation = {
    login: () => '/login',
    language: () => '/lang',
    accounts: () => '/accounts',
    accountsEdit: (id = ":id") => `/accounts/edit/${id}`,
    accountsCreate: () => `/accounts/create`,
    profileChangePassword: () => `/profile/password`,
    tasks: () => '/tasks',
    home: () => '/',
}