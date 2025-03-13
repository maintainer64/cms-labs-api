const en = {
  Login: {
    PageName: 'Login',
    FieldEmail: 'Email',
    FieldPassword: 'Password',
    Submit: 'Login',
    ErrorFieldEmailNotEmpty: 'This field must be an email',
    ErrorFieldEmailRequired: 'Email is required',
    ErrorFieldPasswordRequired: 'Password is required'
  },
  SSO: {
    Wait: 'Please wait...'
  },
  Auth: {
    ErrorPageTitle: 'Authorization Error',
    ErrorPageDescription: 'Try to go to the login page or log in through the course',
    MainTitle: 'CMS LABS',
    MainDescription: '[Admin] Control Management System labs by Hero UI',
    MainChangeLanguage: 'Change language'
  },
  CompaniesDropdown: {
    Title: 'CMS LABS',
    Description: 'UrFU'
  },
  Sidebar: {
    Home: 'Home',
    MainMenu: 'Main Menu',
    Users: 'Users',
    Servers: 'Servers',
    ServiceCards: 'Services',
    LTIIntegrations: 'LTIs',
    LTIRouting: 'LTI Routes',
    LTIAttempts: 'LTI Attempts',
    AnyList: 'List',
    Profile: 'Profile',
    Edit: 'Edit',
    Save: 'Save',
    Confirm: 'Confirm',
    Close: 'Close',
    Delete: 'Delete',
    ViewAll: 'View all'
  },
  LanguageSwitcher: {
    LanguageSwitch: 'Select language'
  },
  UserNavBar: {
    SignedAs: 'Signed in as',
    Logout: 'Log out',
    PasswordChange: 'Change password',
    LanguageChange: 'Change language'
  },
  Tables: {
    UsersTable: {
      Title: 'All users',
      TitleWidgetHome: 'Latest users',
      SearchBar: 'Search users',
      ButtonAdd: 'Add user',
      ButtonEdit: 'Edit user',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'ROLE', uid: 'role' },
        { name: 'ACTIONS', uid: 'actions' }
      ]
    },
    LTIFormsTable: {
      Title: 'LTI Integration',
      ButtonAdd: 'Create',
      SearchBar: 'Search integrations',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'ACTIONS', uid: 'actions' }
      ]
    },
    LTIAttemptsTable: {
      Title: 'LTI Attempts',
      TitleWidgetHome: 'Latest attempts',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'USER', uid: 'user' },
        { name: 'SERVER', uid: 'server' },
        { name: 'NAME', uid: 'name' },
        { name: 'ACTIONS', uid: 'actions' }
      ]
    },
    LTIRoutingTable: {
      Title: 'LTI Routing',
      ButtonAdd: 'Create',
      SearchBar: 'Search route',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'ACTIONS', uid: 'actions' }
      ]
    },
    PnetServersTable: {
      Title: 'PNET Servers',
      ButtonAdd: 'Create',
      SearchBar: 'Search PNET Servers',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: '%', uid: 'unitRate' },
        { name: 'STATUS', uid: 'status' },
        { name: 'ACTIONS', uid: 'actions' }
      ],
      ColumnStatus: {
        DisconnectDistribution: 'Disconnected from traffic',
        ConnectDistribution: 'Accepts traffic',
        Activated: 'Active',
        Deactivated: 'Deactivated'
      }
    },
    ServiceCardsTable: {
      Title: 'Services',
      ButtonAdd: 'Create',
      ButtonEdit: 'Edit Service-Card',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'ORDER', uid: 'order' },
        { name: 'ACTIVE', uid: 'is_active' },
        { name: 'ACTIONS', uid: 'actions' }
      ],
      ColumnStatus: {
        Activated: 'Active',
        Deactivated: 'Deactivated'
      }
    }
  },
  Forms: {
    SaveSuccess: 'Saved',
    SaveError: 'Error when saving',
    DeleteSuccess: 'Deleted',
    DeleteError: 'Error when deleting'
  },
  UserForm: {
    FieldID: 'ID',
    FieldName: 'Name',
    FieldEmail: 'Email',
    FieldUserRole: 'Role',
    FieldUserRoleStudent: 'Student',
    FieldUserRoleInstructor: 'Instructor',
    FieldUserRoleAdmin: 'Admin',
    FieldGroupName: 'Group',
    FieldExternalLTIID: 'LTI ID',
    FieldIsDeactivated: 'Is deactivated',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    FieldRelationAttempts: 'Relation attempts'
  },
  LTIForm: {
    FieldID: 'ID',
    FieldName: 'Name',
    FieldBaseURI: 'Base URL LTI',
    ButtonBaseURI: 'Settings URL',
    ButtonBaseURIMoodle: 'Moodle',
    ButtonBaseURIMoodleDescription: 'Setup URI on base URL Moodle service',
    DescriptionBaseURI: 'URL base service provider LTI',
    FieldLTIAuthLoginUri: 'URL authorization LTI',
    DescriptionLTIAuthLoginUri: 'Address suffix /mod/lti/auth.php',
    FieldLTIAuthTokenUri: 'URL token LTI',
    DescriptionLTIAuthTokenUri: 'Address suffix /mod/lti/token.php',
    FieldTargetLinkUri: 'Link on CMS for login from LTI',
    DescriptionTargetLinkUri: 'Example: https://cms-labs.com/api/v2/lti/launch',
    FieldLTIClientID: 'ClientID LTI',
    FieldLTIDeployment: 'DeploymentID LTI',
    FieldKeySetURI: 'URL certificates LTI',
    DescriptionKeySetURI: 'Address suffix /mod/lti/certs.php',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    MoodleProviderParams: {
      SectionTitle: 'Params for moodle',
      ToolURL: 'Tool URL',
      LTIVersion: 'LTI version',
      LTIVersionValue: 'LTI 1.3',
      PublicKeyType: 'Public key type',
      PublicKeyTypeValue: 'RSA key',
      PublicKey: 'Public key',
      InitiateLoginURL: 'Initiate login URL',
      RedirectionURI: 'Redirection URI(s)',
      DefaultLaunchContainer: 'Default launch container',
      DefaultLaunchContainerValue: 'New window',
      IMSLTIAssignmentGradeServices: 'IMS LTI Assignment and Grade Services',
      IMSLTIAssignmentGradeServicesValue: 'Use this service for grade sync and column management',
      IMSLTINamesRoleProvisioning: 'IMS LTI Names and Role Provisioning',
      IMSLTINamesRoleProvisioningValue: "Use this service to retrieve members' information as per privacy settings",
      ToolSettings: 'Tool Settings',
      ToolSettingsValue: 'Use this service',
      ShareLauncherNameWithTool: "Share launcher's name with tool",
      ShareLauncherNameWithToolValue: 'Always',
      ShareLauncherEmailWithTool: "Share launcher's email with tool",
      ShareLauncherEmailWithToolValue: 'Always',
      AcceptGradesTool: 'Accept grades from the tool',
      AcceptGradesToolValue: 'Always'
    },
    DeletePopup: {
      Title: 'Deleting an LTI-Forms entity',
      Description: 'Upon removal, integration between the LMS system will be terminated'
    }
  },
  LTIFormAttempt: {
    FieldID: 'ID attempt',
    FieldUserId: 'User id',
    FieldUserEmail: 'User email',
    FieldUserName: 'User name',
    FieldLTIRoutingID: 'Routing id',
    FieldLTIRoutingName: 'Routing name',
    FieldPNETServer: 'Pnet server',
    FieldExpiredAt: 'Expired at',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting an LTI-Attempt entity',
      Description:
        'When you delete it, the other systems will not be able to interact with the user. It is safe if the user has completed the work'
    }
  },
  PnetServers: {
    FieldID: 'ID',
    FieldName: 'Name',
    FieldURL: 'URL',
    FieldType: 'Server type',
    FieldClientID: 'ClientID',
    FieldClientIDDescription: 'Unique server name for SSO',
    FieldIsActive: 'Active',
    FieldUnitRate: 'Percentage of distribution',
    FieldMinutesForDisconnect: 'Minutes for auto-shutdown',
    DescriptionMinutesForDisconnect: 'Used to disable user allocation',
    FieldMaxCountUsers: 'Max count users for auto-shutdown',
    DescriptionMaxCountUsers: 'Used to disable user allocation',
    FieldToken: 'Token',
    FieldLastOnlineStatus: 'Last active at',
    FieldLastCountUsers: 'Last count users',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting a PNET Server entity',
      Description: "When you take the PNET out, the PNET and CMS won't work together anymore"
    },
    StatsChart: {
      Title: 'PNET stats',
      Status: 'Status',
      StatusValueActive: 'Active',
      StatusValueAll: 'All',
      OrderBy: 'Sort',
      OrderByValueUnitRate: '% rate.',
      OrderByValueLastCountUsers: 'count users',
      OrderByValueCreatedAt: 'date'
    }
  },
  ServiceCards: {
    FieldID: 'ID',
    FieldImageURL: 'Image URL',
    FieldURL: 'Redirect URL',
    FieldIsActive: 'Active',
    FieldName: 'Title',
    FieldDescription: 'Description',
    FieldOrder: 'Order',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting a Service Card entity',
      Description: 'When you delete the service, it will no longer be displayed on the main screen'
    }
  },
  PnetServersQueue: {
    RoundRobinChart: {
      Title: 'PNET distribution'
    },
    LTIAttemptRoom: {
      ErrorPageTitle: 'Distribution error',
      ErrorPageRefresh: 'Try again'
    }
  },
  LTIRouting: {
    FieldID: 'ID',
    FieldName: 'Name',
    SectionRouteParams: 'Route params',
    SectionRouteParamsDescription: 'Set the parameters for performing actions on this route',
    FieldLTITitle: 'Name course element',
    FieldLTIDescription: 'Description course element',
    FieldLTITaskID: 'ID or URI course element',
    FieldLTIParamsTask: 'Query LTI course element',
    SectionActionParams: 'Action when connecting',
    FieldCollaboration: 'Number of people to work together',
    FieldPinnedSessionMinutes: 'Pinned session on minutes',
    FieldPNETLabsType: 'Type connection',
    FieldPNETLabsTypeDefault: 'Default',
    FieldPNETLabsTypeFile: 'File',
    FieldPNETLabsTypeEnumeration: 'Enumeration',
    FieldPNETLabsPath: 'Connecting to laboratory work',
    FieldPNETTestPath: 'Running lab tests',
    FieldIsDefault: 'Route is default',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting an LTI-Routing entity',
      Description: 'Routing for this task from the LMS system will be terminated'
    }
  },
  UserFormPasswordChange: {
    FieldOldPassword: 'Current password',
    FieldNewPassword: 'New password',
    FieldAgainPassword: 'New password again'
  }
};
export default en;
