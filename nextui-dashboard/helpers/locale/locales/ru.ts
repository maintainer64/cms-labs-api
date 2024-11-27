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
  Auth: {
    ErrorPageTitle: 'Ошибка авторизации',
    ErrorPageDescription: 'Попробуйте перейти на страницу входа в систему или авторизоваться через курс',
    MainTitle: 'СУиМ Лаб',
    MainDescription: 'Система Управления и Мониторинга лаболаторных работ. Сделано на Next.js',
    MainChangeLanguage: 'Сменить язык'
  },
  CompaniesDropdown: {
    Title: 'СУиМ Лаб',
    Description: 'УрФУ',
    ContentService: 'Сервисы'
  },
  Sidebar: {
    Home: 'Главная',
    MainMenu: 'Главное меню',
    Users: 'Пользователи',
    Servers: 'Серверы',
    LTIIntegrations: 'LTIs',
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
        { name: 'РОЛЬ', uid: 'role' },
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
    PnetServersTable: {
      Title: 'Серверы PNET',
      ButtonAdd: 'Создать',
      SearchBar: 'Поиск серверов',
      ButtonEdit: 'Редактировать',
      Columns: [
        { name: 'ID', uid: 'id' },
        { name: 'НАЗВАНИЕ', uid: 'name' },
        { name: '%', uid: 'unitRate' },
        { name: 'СТАТУС', uid: 'status' },
        { name: 'ДЕЙСТВИЕ', uid: 'actions' }
      ],
      ColumnStatus: {
        DisconnectDistribution: 'Отключен от трафика',
        ConnectDistribution: 'Принимает трафик',
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
    FieldUserRole: 'Роль',
    FieldUserRoleStudent: 'Студент',
    FieldUserRoleInstructor: 'Преподаватель',
    FieldUserRoleAdmin: 'Администратор',
    FieldGroupName: 'Группа',
    FieldExternalLTIID: 'LTI ID',
    FieldIsActive: 'Активный',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в'
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
    FieldCreatedAt: 'Создана в',
    FieldUpdatedAt: 'Обновлена в',
    MoodleProviderParams: {
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
  PnetServers: {
    FieldID: 'ID',
    FieldName: 'Название',
    FieldURL: 'URL',
    FieldIsActive: 'Активный',
    FieldUnitRate: 'Процент распределения',
    FieldMinutesForDisconnect: 'Минуты для автоотключения',
    DescriptionMinutesForDisconnect: 'Используется для отключения распределения пользователей',
    FieldMaxCountUsers: 'Максимальное кол-во пользователей',
    DescriptionMaxCountUsers: 'Используется для отключения распределения пользователей',
    FieldToken: 'Токен',
    FieldLastOnlineStatus: 'Последняя актиность в',
    FieldLastCountUsers: 'Последнее кол-во пользователей',
    FieldCreatedAt: 'Создан в',
    FieldUpdatedAt: 'Обновлен в',
    DeletePopup: {
      Title: 'Удаление сущности PNET-Server',
      Description: 'При удалении интеграция между PNET сервером и CMS системой будет прекращена'
    },
    StatsChart: {
      Title: 'PNET статистика',
      Status: 'Статус',
      StatusValueActive: 'Активные',
      StatusValueAll: 'Все',
      OrderBy: 'Сортировка',
      OrderByValueUnitRate: '% распр.',
      OrderByValueLastCountUsers: 'кол-во польз.',
      OrderByValueCreatedAt: 'дата'
    }
  },
  PnetServersQueue: {
    RoundRobinChart: {
      Title: 'PNET распределение'
    }
  },
  UserFormPasswordChange: {
    FieldOldPassword: 'Текущий пароль',
    FieldNewPassword: 'Новый пароль',
    FieldAgainPassword: 'Подтверждение нового пароля'
  }
};
export default ru;
