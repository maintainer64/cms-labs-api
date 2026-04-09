const en = {
  Login: {
    PageName: 'Login',
    TabSSO: 'Moodle',
    TabInternal: 'Internal',
    FieldEmail: 'Email',
    FieldEmailDescription: 'Enter your business email address',
    FieldPassword: 'Password',
    Submit: 'Login',
    ErrorFieldEmailNotEmpty: 'This field must be an email',
    ErrorFieldEmailRequired: 'Email is required',
    ErrorFieldPasswordRequired: 'Password is required'
  },
  SSO: {
    Wait: 'Please wait...',
    OR: 'or'
  },
  Auth: {
    ErrorPageTitle: 'Authorization Error',
    ErrorPageDescription: 'Try to go to the login page or log in through the course',
    MainTitle: 'CMS LABS',
    MainDescription: '[Admin] Control Management System labs',
    MainChangeLanguage: 'Сменить язык'
  },
  CompaniesDropdown: {
    Title: 'CMS LABS',
    Description: 'UrFU'
  },
  Sidebar: {
    Home: 'Home',
    MainMenu: 'Main Menu',
    Users: 'Users',
    Roles: 'Roles',
    Servers: 'Servers',
    Targets: 'Map',
    Collapse: 'Collapse',
    CurlRequests: 'API',
    ServiceCards: 'Services',
    AuthProviders: 'Auth providers',
    LTIRouting: 'LTI Routes',
    LTIAttempts: 'LTI Attempts',
    APIRequests: 'API Requests',
    AnyList: 'List',
    Profile: 'Profile',
    Edit: 'Edit',
    Save: 'Save',
    Confirm: 'Confirm',
    Cancel: 'Cancel',
    Delete: 'Delete',
    Close: 'Close',
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
      Title: 'All user',
      TitleWidgetHome: 'Latest user',
      SearchBar: 'Search user',
      ButtonAdd: 'Add user',
      ButtonEdit: 'Edit user',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'ROLES', uid: 'role' },
        { name: 'ACTIONS', uid: 'actions' }
      ]
    },
    AuthProvidersTable: {
      Title: 'Auth providers',
      ButtonAdd: 'Create',
      SearchBar: 'Search providers',
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
        { name: 'STATUS', uid: 'status' },
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
    CurlRequestTable: {
      Title: 'API Requests',
      ButtonAdd: 'Create',
      SearchBar: 'Search API',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'ACTIONS', uid: 'actions' }
      ]
    },
    PnetServersTable: {
      Title: 'All Servers',
      ButtonAdd: 'Create',
      SearchBar: 'Search All Servers',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'INDICATORS', uid: 'indicator' },
        { name: 'STATUS', uid: 'status' },
        { name: 'ROLES', uid: 'role' },
        { name: 'ACTIONS', uid: 'actions' }
      ],
      ColumnIndicator: {
        UnitRate: 'Distribution',
        LastCountUsers: 'Labs'
      },
      ColumnStatus: {
        DisconnectDistribution: 'Disconnected from traffic',
        ConnectDistribution: 'Accepts traffic',
        Activated: 'Active',
        Deactivated: 'Deactivated'
      }
    },
    RoleTable: {
      Title: 'Roles',
      ButtonAdd: 'Create',
      SearchBar: 'Search All Roles',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'NAME', uid: 'name' },
        { name: 'ACTIONS', uid: 'actions' }
      ]
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
    FieldUserRole: 'Roles',
    FieldGroupName: 'Group',
    FieldExternalLTIID: 'LTI ID',
    FieldIsDeactivated: 'Is deactivated',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    FieldRelationAttempts: 'Relation attempts'
  },
  AuthProvider: {
    FieldID: 'ID',
    FieldName: 'Name',
    FieldType: 'Тип',
    FieldTypeLTI: 'LTI',
    FieldTypeLTIDescription: 'Connect to course',
    FieldTypeLDAP: 'LDAP',
    FieldTypeLDAPDescription: 'Connect to lightweight directory',
    FieldBaseURI: 'Base URL',
    ButtonBaseURI: 'Settings URL',
    ButtonBaseURIMoodle: 'Moodle',
    ButtonBaseURIMoodleDescription: 'Setup URI on base URL Moodle service',
    DescriptionBaseURI: 'URL base service provider',
    FieldLTIAuthLoginUri: 'URL authorization LTI',
    DescriptionLTIAuthLoginUri: 'Address suffix /mod/lti/auth.php',
    FieldLTIAuthTokenUri: 'URL token LTI',
    DescriptionLTIAuthTokenUri: 'Address suffix /mod/lti/token.php',
    FieldTargetLinkUri: 'Link on CMS for login from LTI',
    DescriptionTargetLinkUri: 'Example: https://cms-labs.com/api/v2/lti/launch',
    FieldLTIClientID: 'ClientID LTI',
    FieldLTIDeployment: 'DeploymentID LTI',
    FieldKeySetURI: 'URL certificates LTI',
    FieldDN: 'DN',
    DescriptionKeySetURI: 'Address suffix /mod/lti/certs.php',
    FieldSSOURL: 'SSO LTI URL',
    DescriptionSSOURL: 'Link to the course element in the LTI tool',
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
  CurlRequest: {
    FieldID: 'ID',
    FieldName: 'Name',
    FieldCurl: 'Curl',
    FieldTimeout: 'Timeout',
    FieldTimeoutDescription: 'Request termination in seconds',
    FieldCurlDescription: 'Parameters passed via $ are supported ($userID, $SessionID)',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting an API request',
      Description: 'It will be impossible to call the API when deleting it'
    }
  },
  AuthProviderAttempt: {
    FieldID: 'ID attempt',
    FieldRoomNumber: 'Room number',
    FieldUserId: 'User id',
    FieldUserEmail: 'User email',
    FieldUserName: 'User name',
    FieldLTIRoutingID: 'Routing id',
    FieldLTIRoutingName: 'Routing name',
    FieldPNETServer: 'Pnet server',
    FieldExpiredAt: 'Expired at',
    FieldStatus: 'Status',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting an Auth provider entity',
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
    FieldAllowedRoles: 'Allowed roles',
    FieldAllowedRolesDescription: 'Only these roles will be able to log in',
    FieldIsActive: 'Active',
    FieldUnitRate: 'Percentage of distribution',
    FieldMinutesForDisconnect: 'Minutes for auto-shutdown',
    DescriptionMinutesForDisconnect: 'Used to disable user allocation',
    FieldMaxCountUsers: 'Max count user for auto-shutdown',
    DescriptionMaxCountUsers: 'Used to disable user allocation',
    FieldToken: 'Token',
    FieldLastOnlineStatus: 'Last active at',
    FieldLastCountUsers: 'Last count work labs',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting a Server entity',
      Description: "When you take the Server out, the Server and CMS won't work together anymore"
    },
    StatsChart: {
      Title: 'Server stats',
      Status: 'Status',
      StatusValueActive: 'Active',
      StatusValueAll: 'All',
      OrderBy: 'Sort',
      OrderByValueUnitRate: '% rate.',
      OrderByValueLastCountUsers: 'count lab',
      OrderByValueCreatedAt: 'date'
    }
  },
  Roles: {
    FieldID: 'ID',
    FieldName: 'Name',
    FieldCode: 'Code',
    FieldCreatedAt: 'Created at',
    FieldUpdatedAt: 'Updated at',
    DeletePopup: {
      Title: 'Deleting a Role entity',
      Description: 'This role will be disabled for all user'
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
      ErrorPageRefresh: 'Try again',
      ModalTitle: 'Connecting to laboratory',
      ModalDescription: 'Tell this room number to another person or write it yourself.',
      ModalDescriptionChange: 'Press the "change" button to change the room.',
      ButtonCopy: 'Copy',
      ButtonRollback: 'Rollback',
      ButtonChange: 'Change',
      RoomChangeSuccess: 'The room has been changed',
      RoomChangeError: 'The room has not been changed'
    },
    LTIAttemptConfirm: {
      ModalDescription: 'Do you want to connect to the lab?',
      ButtonAdminPanel: 'No',
      ButtonLab: 'Connect'
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
    FieldPNETLabsTypeDefaultDescription: 'Send params to other service no validation',
    FieldPNETLabsTypeCurl: 'API Request',
    FieldPNETLabsTypeCurlDescription: 'Send identifier API curl request',
    FieldPNETLabsTypeSSO: 'SSO',
    FieldPNETLabsTypeSSODescription: 'Only connect to cms system',
    FieldPNETLabsTypeClabgate: 'Clabgate',
    FieldPNETLabsTypeClabgateDescription: 'Deploy into CI/CD system',
    FieldPNETLabsPath: 'Connecting to laboratory work',
    FieldPNETTestPath: 'Running lab tests',
    FieldPNETServer: 'Auto connect to server',
    FieldPNETServerDefault: 'Auto',
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
  },
  RoleBasedAccess: {
    Title: 'Forbidden',
    Description: 'Sorry, access to view this page is denied with your role',
    Button: 'Go back'
  },
  Target: {
    Title: 'Map services',
    ButtonAdd: 'Create',
    SearchBar: 'Search',
    ButtonEdit: 'Edit',
    FieldID: 'ID',
    FieldName: 'Name',
    FieldDescription: 'Description',
    FieldType: 'Type',
    FieldTags: 'Tags',
    FieldLinks: 'Links',
    FieldCreatedAt: 'Created',
    FieldUpdatedAt: 'Updated',
    FieldSynchronizedAt: 'Synchronized',
    Types: {
      server: 'Server',
      virtual: 'Virtual Server',
      service: 'Service',
      module: 'Module'
    },
    Relation: {
      Title: 'Relations',
      Parent: 'Parent',
      AddParent: 'Add relation',
      RemoveParent: 'Remove relation',
      SelectParent: 'Select relation',
      NoParents: 'No relations'
    },
    Addon: {
      Title: 'Addons',
      Connect: 'Connect',
      NoAddons: 'Not connected addons',
      ConnectTitle: 'Connect addon',
      Name: 'Name',
      Type: 'Type',
      Disconnect: 'Disconnect',
      Reset: 'Reset',
      ResetTitle: 'Reset tokens',
      ResetDescription:
        'Access tokens will be recreated. Data in the resource will not be changed or deleted. Continue?',
      SelectAddon: 'Select addon',
      DeleteRequest: 'Delete request',
      DeleteConfirmStep1: 'Create backup of data',
      DeleteConfirmStep2: 'Completely clear application data',
      DeleteConfirmStep3: 'The addon will be permanently deleted and cannot be restored',
      DeleteConfirmWarning: 'Are you sure you want to delete this addon?',
      DeleteConfirmButton: 'Yes, delete',
      DeleteCancelButton: 'Cancel',
      DeletePendingMessage: 'Waiting for deletion confirmation from another user',
      DeleteRequestedBy: 'Deletion requested by user',
      DeleteConfirmRequired: 'Deletion confirmation required',
      Config: 'Configuration',
      DatabaseSize: 'Used',
      Cluster: 'Cluster',
      Namespace: 'Namespace',
      Registry: 'Image registry',
      ApiKeyInVault: 'API key in Vault',
      RotateCredentials: 'Rotate credentials',
      RotateCredentialsTitle: 'Rotate credentials',
      RotateCredentialsDescription: 'Keys and passwords will be recreated. Old keys will become invalid. Continue?',
      RevokeDeleteRequest: 'Revoke delete request',
      RevokeDeleteRequestTitle: 'Revoke delete request',
      RevokeDeleteRequestDescription: 'Delete request will be canceled. Continue?',
      Expires: 'Expires'
    },
    User: {
      Title: 'Users',
      AddTitle: 'Add user',
      EditTitle: 'Edit permissions',
      Add: 'Add user',
      Remove: 'Remove',
      Email: 'Email',
      Roles: 'Roles',
      UserLabel: 'User',
      UserPlaceholder: 'Start typing name',
      RolesLabel: 'Roles',
      SelectUser: 'Select user',
      SelectRole: 'Select role',
      AddButton: 'Add user',
      NoUsers: 'No linked users',
      NoRoles: 'No roles'
    },
    Roles: [
      {
        key: 'vault_viewer',
        value: 'Vault Viewer'
      },
      {
        key: 'vault_writer',
        value: 'Vault Writer'
      },
      {
        key: 'editor',
        value: 'Editor'
      },
      {
        key: 'nominal',
        value: 'Nominal'
      },
      {
        key: 'nominal',
        value: 'Harbor Access'
      }
    ],
    DeletePopup: {
      Title: 'Delete target',
      Description: 'Are you sure you want to delete this object?'
    }
  },
  Topology: {
    Menu: {
      OpenLogs: 'Open logs topology',
      RestartTopology: 'Restart topology',
      RestartTopologyTitle: 'Restart topology',
      RestartTopologyDescription:
        'The initial settings for the topology are applied, all connections are reset, and the updated system configuration is raised',
      CopyTokenModal: {
        Title: 'Copied',
        Description: 'Access token copied to clipboard'
      },
      ConnectToKubectl: 'Connect to kubectl',
      ConnectToKubectlTitle: 'Manual of connection to system',
      RemoveTopology: 'Remove topology',
      RemoveTopologyTitle: 'Remove topology',
      RemoveTopologyDescription: 'Delete topology and all resources associated with it?',
      RemoveTopologyDescriptionSuccess: 'Topology success removed'
    },
    Connect: {
      Error: 'Internal error',
      ErrorModalViewTitle: 'Topology display',
      ErrorModalConnectTitle: 'Connection to topology',
      ErrorModalRetry: 'Try again',
      WaitModalTitle: 'The topology is still loading',
      WaitModalDescription: 'You can view the logs of topology loading',
      WaitModalButtonText: 'View logs'
    },
    Terminal: {
      ReadyStatus: 'Ready',
      NotReadyStatus: 'Not ready',
      Restarts: 'Restarts',
      ConnectionError: 'Connection error',
      TerminalConnected: 'Terminal connected',
      ConnectionClosed: 'Connection closed',
      UnknownReason: 'Unknown reason'
    }
  }
};
export default en;
