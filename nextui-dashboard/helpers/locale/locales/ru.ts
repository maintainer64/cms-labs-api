const ru = {
  Login: {
    PageName: 'Вход',
    FieldEmail: 'Почта',
    FieldPassword: 'Пароль',
    Submit: 'Войти',
    ErrorFieldEmailNotEmpty: 'Это должна быть почта',
    ErrorFieldEmailRequired: 'Email не указан',
    ErrorFieldPasswordRequired: 'Пароль не указан'
  },
  SSO: {
    Wait: 'Подождите...',
    OR: 'или'
  },
  Auth: {
    ErrorPageTitle: 'Ошибка авторизации',
    ErrorPageDescription: 'Попробуйте перейти на страницу входа в систему или авторизоваться через курс',
    MainTitle: 'СУиМ Лаб',
    MainDescription: 'Система Управления и Мониторинга лаболаторных работ. Сделано на Hero UI',
    MainChangeLanguage: 'Change language'
  },
  CompaniesDropdown: {
    Title: 'СУиМ Лаб',
    Description: 'УрФУ'
  },
  Sidebar: {
    Home: 'Главная',
    MainMenu: 'Главное меню',
    Users: 'Пользователи',
    Roles: 'Роли',
    Servers: 'Серверы',
    CurlRequests: 'API',
    ServiceCards: 'Сервисы',
    LTIIntegrations: 'LTIs',
    LTIRouting: 'Маршруты LTI',
    LTIAttempts: 'Попытки LTI',
    APIRequests: 'Запросы по API',
    AnyList: 'Список',
    Profile: 'Профиль',
    Edit: 'Редактирование',
    Save: 'Сохранить',
    Confirm: 'Подтвердить',
    Delete: 'Удалить',
    Close: 'Закрыть',
    ViewAll: 'Перейти'
  },
  LanguageSwitcher: {
    LanguageSwitch: 'Выберите язык'
  },
  UserNavBar: {
    SignedAs: 'Вход по',
    Logout: 'Выйти',
    PasswordChange: 'Сменить пароль',
    LanguageChange: 'Сменить язык'
  },
  Tables: {
    UsersTable: {
      Title: 'Все пользователи',
      TitleWidgetHome: 'Последние пользователи',
      SearchBar: 'Найти пользователя',
      ButtonAdd: 'Создать',
      ButtonEdit: 'Редактировать пользователя',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'ИМЯ', uid: 'name' },
        { name: 'РОЛИ', uid: 'role' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ]
    },
    LTIFormsTable: {
      Title: 'Интеграции LTI',
      ButtonAdd: 'Создать',
      SearchBar: 'Поиск интеграций',
      ButtonEdit: 'Редактировать',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ]
    },
    LTIAttemptsTable: {
      Title: 'Попытки LTI',
      TitleWidgetHome: 'Последние попытки',
      ButtonEdit: 'Редактировать',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'ПОЛЬЗОВАТЕЛЬ', uid: 'user' },
        { name: 'СЕРВЕР', uid: 'server' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ]
    },
    LTIRoutingTable: {
      Title: 'Маршруты LTI',
      ButtonAdd: 'Создать',
      SearchBar: 'Поиск маршрутов',
      ButtonEdit: 'Редактировать',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ]
    },
    CurlRequestTable: {
      Title: 'Запросы по API',
      ButtonAdd: 'Создать',
      SearchBar: 'Поиск API',
      ButtonEdit: 'Edit',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ]
    },
    PnetServersTable: {
      Title: 'Все серверы',
      ButtonAdd: 'Создать',
      SearchBar: 'Поиск серверов',
      ButtonEdit: 'Редактировать',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: 'ПОКАЗАТЕЛИ', uid: 'indicator' },
        { name: 'СТАТУС', uid: 'status' },
        { name: 'РОЛИ', uid: 'role' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ],
      ColumnIndicator: {
        UnitRate: 'Распределение',
        LastCountUsers: 'Лаб.'
      },
      ColumnStatus: {
        DisconnectDistribution: 'Отключен от трафика',
        ConnectDistribution: 'Принимает трафик',
        Activated: 'Активен',
        Deactivated: 'Деактивирован'
      }
    },
    RoleTable: {
      Title: 'Все роли',
      ButtonAdd: 'Создать',
      SearchBar: 'Поиск ролей',
      ButtonEdit: 'Редактировать',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ]
    },
    ServiceCardsTable: {
      Title: 'Сервисы',
      ButtonAdd: 'Создать',
      ButtonEdit: 'Редактировать',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: 'ПОЗИЦИЯ', uid: 'order' },
        { name: 'АКТИВЕН', uid: 'is_active' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ],
      ColumnStatus: {
        Activated: 'Активен',
        Deactivated: 'Деактивирован'
      }
    }
  },
  Forms: {
    SaveSuccess: 'Сохранено',
    SaveError: 'Ошибка при сохранении',
    DeleteSuccess: 'Удалено',
    DeleteError: 'Ошибка при удалении'
  },
  UserForm: {
    FieldID: 'ID',
    FieldName: 'Имя',
    FieldEmail: 'Почта',
    FieldUserRole: 'Роли',
    FieldGroupName: 'Группа',
    FieldExternalLTIID: 'LTI ID',
    FieldIsDeactivated: 'Деактивирован',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в',
    FieldRelationAttempts: 'Связные попытки'
  },
  LTIForm: {
    FieldID: 'ID',
    FieldName: 'Название',
    FieldBaseURI: 'Базовый URL LTI',
    ButtonBaseURI: 'Настройка URL',
    ButtonBaseURIMoodle: 'Moodle',
    ButtonBaseURIMoodleDescription: 'Установка URL по базовому адресу Moodle',
    DescriptionBaseURI: 'URL базовый инструмента LTI',
    FieldLTIAuthLoginUri: 'URL авторизации LTI',
    DescriptionLTIAuthLoginUri: 'Адрес с /mod/lti/auth.php',
    FieldLTIAuthTokenUri: 'URL токен LTI',
    DescriptionLTIAuthTokenUri: 'Адрес с /mod/lti/token.php',
    FieldTargetLinkUri: 'Ссылка CMS для входа через LTI',
    DescriptionTargetLinkUri: 'Ссылка вида https://cms-labs.com/api/v2/lti/launch',
    FieldLTIClientID: 'ID клиента LTI',
    FieldLTIDeployment: 'ID deployment LTI',
    FieldKeySetURI: 'URL для получения сертификатов LTI',
    DescriptionKeySetURI: 'Адрес с /mod/lti/certs.php',
    FieldSSOURL: 'SSO LTI URL',
    DescriptionSSOURL: 'Ссылка на элемент курса в инструменте LTI',
    FieldCreatedAt: 'Создана в',
    FieldUpdatedAt: 'Обновлена в',
    MoodleProviderParams: {
      SectionTitle: 'Параметры для Moodle',
      ToolURL: 'Базовый URL-адрес инструмента',
      LTIVersion: 'Версия LTI',
      LTIVersionValue: 'LTI 1.3',
      PublicKeyType: 'Тип открытого ключа',
      PublicKeyTypeValue: 'Ключ RSA',
      PublicKey: 'Открытый ключ',
      InitiateLoginURL: 'URL-адрес инициирования входа',
      RedirectionURI: 'URI перенаправления',
      DefaultLaunchContainer: 'Контейнер для запуска инструмента по умолчанию',
      DefaultLaunchContainerValue: 'Новое окно',
      IMSLTIAssignmentGradeServices: 'Службы заданий и оценок IMS LTI',
      IMSLTIAssignmentGradeServicesValue: 'Использовать эту службу для синхронизации оценок и управления столбцами',
      IMSLTINamesRoleProvisioning: 'Предоставление имен и ролей IMS LTI',
      IMSLTINamesRoleProvisioningValue:
        'Использовать эту службу для получения информации об участниках в соответствии с настройками конфиденциальности.',
      ToolSettings: 'Настройки инструмента',
      ToolSettingsValue: 'Использовать эту службу',
      ShareLauncherNameWithTool: 'Определять полное имя пользователя, запускающего инструмент',
      ShareLauncherNameWithToolValue: 'Всегда',
      ShareLauncherEmailWithTool: 'Определять адрес электронной почты пользователя, запускающего инструмента',
      ShareLauncherEmailWithToolValue: 'Всегда',
      AcceptGradesTool: 'Принимать оценки от инструмента',
      AcceptGradesToolValue: 'Всегда'
    },
    DeletePopup: {
      Title: 'Удаление сущности LTI-Forms',
      Description: 'При удалении интеграция между LMS системой будет прекращена'
    }
  },
  CurlRequest: {
    FieldID: 'ID',
    FieldName: 'Название',
    FieldTimeout: 'Таймаут',
    FieldTimeoutDescription: 'Прекращение запроса в секундах',
    FieldCurl: 'Curl',
    FieldCurlDescription: 'Поддерживаются параметры переданные через $ ($userID, $sessionID)',
    FieldCreatedAt: 'Создано в',
    FieldUpdatedAt: 'Обновлено в',
    DeletePopup: {
      Title: 'Удаление запроса по API',
      Description: 'При удалении будет невозможно вызвать API'
    }
  },
  LTIFormAttempt: {
    FieldID: 'ID попытки',
    FieldRoomNumber: 'Номер комнаты',
    FieldUserId: 'ID пользователя',
    FieldUserEmail: 'Почта пользователя',
    FieldUserName: 'ФИО пользователя',
    FieldLTIRoutingID: 'ID маршрута',
    FieldLTIRoutingName: 'Название маршрута',
    FieldPNETServer: 'Сервер',
    FieldExpiredAt: 'Закреплен до',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в',
    DeletePopup: {
      Title: 'Удаление сущности LTI-Attempt',
      Description:
        'При удалении остальные системы не смогут взаимодействовать с пользователем. Безопасно, если пользователь завершил работу'
    }
  },
  PnetServers: {
    FieldID: 'ID',
    FieldName: 'Название',
    FieldURL: 'URL',
    FieldType: 'Тип сервера',
    FieldClientID: 'ClientID',
    FieldClientIDDescription: 'Уникальное название сервера для SSO',
    FieldAllowedRoles: 'Разрешенные роли',
    FieldAllowedRolesDescription: 'Только эти роли смогут авторизоваться',
    FieldIsActive: 'Активный',
    FieldUnitRate: 'Процент распределения',
    FieldMinutesForDisconnect: 'Минуты для автоотключения',
    DescriptionMinutesForDisconnect: 'Используется для отключения распределения пользователей',
    FieldMaxCountUsers: 'Максимальное кол-во пользователей',
    DescriptionMaxCountUsers: 'Используется для отключения распределения пользователей',
    FieldToken: 'Токен',
    FieldLastOnlineStatus: 'Последняя актиность в',
    FieldLastCountUsers: 'Последнее кол-во запущенных лабораторных работ',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в',
    DeletePopup: {
      Title: 'Удаление сущности Server',
      Description: 'При удалении интеграция между сервером и CMS системой будет прекращена'
    },
    StatsChart: {
      Title: 'Статистика серверов',
      Status: 'Статус',
      StatusValueActive: 'Активные',
      StatusValueAll: 'Все',
      OrderBy: 'Сортировка',
      OrderByValueUnitRate: '% распр.',
      OrderByValueLastCountUsers: 'кол-во лаб.',
      OrderByValueCreatedAt: 'дата'
    }
  },
  Roles: {
    FieldID: 'ID',
    FieldName: 'Название',
    FieldCode: 'Код',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в',
    DeletePopup: {
      Title: 'Удаление сущности Role',
      Description: 'Данная роль отключится у всех пользователей'
    }
  },
  ServiceCards: {
    FieldID: 'ID',
    FieldImageURL: 'URL для картинки',
    FieldURL: 'URL для перехода',
    FieldIsActive: 'Активный',
    FieldName: 'Название',
    FieldDescription: 'Описание',
    FieldOrder: 'Сортировка',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в',
    DeletePopup: {
      Title: 'Удаление сущности Service Card',
      Description: 'При удалении сервис перестанет отображаться на главном экране'
    }
  },
  PnetServersQueue: {
    RoundRobinChart: {
      Title: 'PNET распределение'
    },
    LTIAttemptRoom: {
      ErrorPageTitle: 'Ошибка при распределении',
      ErrorPageRefresh: 'Попробовать снова',
      ModalTitle: 'Подключение к лабораторной работе',
      ModalDescription: 'Сообщите этот номер комнаты другому человеку или напишите его сами.',
      ModalDescriptionChange: 'Нажмите кнопку "изменить" для перехода.',
      ButtonCopy: 'Скопировать',
      ButtonRollback: 'Вернуть',
      ButtonChange: 'Изменить',
      RoomChangeSuccess: 'Комната изменена',
      RoomChangeError: 'Комната не изменена'
    },
    LTIAttemptConfirm: {
      ModalDescription: 'Вы хотите подключиться к лабораторной работе?',
      ButtonAdminPanel: 'Нет',
      ButtonLab: 'Подключиться'
    }
  },
  LTIRouting: {
    FieldID: 'ID',
    FieldName: 'Название',
    SectionRouteParams: 'Параметры маршрутизации',
    SectionRouteParamsDescription: 'Задайте параметры для выполнения действий по данному маршруту',
    FieldLTITitle: 'Название элемента курса',
    FieldLTIDescription: 'Описание элемента курса',
    FieldLTITaskID: 'ID или адрес элемента курса',
    FieldLTIParamsTask: 'Query LTI для элемента курса',
    SectionActionParams: 'Действие при подключении',
    FieldCollaboration: 'Кол-во человек для совместной работы',
    FieldPinnedSessionMinutes: 'Закрепить сессию в минутах',
    FieldPNETLabsType: 'Тип подключения',
    FieldPNETLabsTypeDefault: 'Прямое',
    FieldPNETLabsTypeDefaultDescription: 'Отправка параметров без верификации',
    FieldPNETLabsTypeCurl: 'Запрос по API',
    FieldPNETLabsTypeCurlDescription: 'Отправляется API-ID запроса для выполнения',
    FieldPNETLabsTypeSSO: 'SSO',
    FieldPNETLabsTypeSSODescription: 'Только подключение к cms системе',
    FieldPNETLabsTypeClabgate: 'Clabgate',
    FieldPNETLabsTypeClabgateDescription: 'Развертка в CI/CD системе',
    FieldPNETLabsPath: 'Подключение к лабораторной работе',
    FieldPNETTestPath: 'Запуск тестов лабораторной работы',
    FieldPNETServer: 'Автоматическое подключение к серверу',
    FieldPNETServerDefault: 'Автоматически',
    FieldIsDefault: 'Маршрут по умолчанию',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в',
    DeletePopup: {
      Title: 'Удаление сущности LTI-Routing',
      Description: 'Маршрутизация по данному заданию из LMS системы будет прекращена'
    }
  },
  UserFormPasswordChange: {
    FieldOldPassword: 'Текущий пароль',
    FieldNewPassword: 'Новый пароль',
    FieldAgainPassword: 'Подтверждение нового пароля'
  },
  RoleBasedAccess: {
    Title: 'Доступ запрещен',
    Description: 'Извините, доступ для просмотра этой страницы запрещен с вашей ролью',
    Button: 'Вернуться назад'
  },
  Topology: {
    Menu: {
      OpenLogs: 'Открыть логи топологии',
      RestartTopology: 'Перезагрузить топологию',
      RestartTopologyTitle: 'Перезагрузить топологию',
      RestartTopologyDescription:
        'Происходит применение исходных настроект для топологии, сбрасываются все подключения, поднимается обновленная конфигурация системы',
      CopyTokenModal: {
        Title: 'Скопировано',
        Description: 'Токен доступа скопирован в буфер обмена'
      },
      ConnectToKubectl: 'Подключиться к kubectl',
      ConnectToKubectlTitle: 'Инструкция для подключения к kubectl',
      RemoveTopology: 'Удалить топологию',
      RemoveTopologyTitle: 'Удалить топологию',
      RemoveTopologyDescription: 'Удалить топологию и все ресурсы связанные с ней?',
      RemoveTopologyDescriptionSuccess: 'Топология успешно удалена'
    },
    Connect: {
      Error: 'Внутренняя ошибка',
      ErrorModalViewTitle: 'Отображение топологии',
      ErrorModalConnectTitle: 'Подключение к топологии',
      ErrorModalRetry: 'Попробовать снова',
      WaitModalTitle: 'Топология ещё загружается',
      WaitModalDescription: 'Вы можете просмотреть логи загрузки топологии',
      WaitModalButtonText: 'Просмотреть логи'
    },
    Terminal: {
      ReadyStatus: 'Готов',
      NotReadyStatus: 'Не готов',
      Restarts: 'Перезапусков',
      ConnectionError: 'Ошибка подключения',
      TerminalConnected: 'Терминал подключен',
      ConnectionClosed: 'Подключение закрыто',
      UnknownReason: 'Неизвестная ошибка'
    }
  }
};
export default ru;
