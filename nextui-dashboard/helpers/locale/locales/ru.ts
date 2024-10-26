const ru = {
    "Login": {
        "PageName": "Вход",
        "FieldEmail": "Почта",
        "FieldPassword": "Пароль",
        "Submit": "Войти",
        "ErrorFieldEmailNotEmpty": "Это должна быть почта",
        "ErrorFieldEmailRequired": "Email не указан",
        "ErrorFieldPasswordRequired": "Пароль не указан",
    },
    "Auth": {
        "ErrorPageTitle": "Ошибка авторизации",
        "ErrorPageDescription": "Попробуйте перейти на страницу входа в систему или авторизоваться через курс",
        "MainTitle": "СУиМ Лаб",
        "MainDescription": "Система Управления и Мониторинга лаболаторных работ. Сделано на Next.js",
        "MainChangeLanguage": "Сменить язык",
    },
    "CompaniesDropdown": {
        "Title": "СУиМ Лаб",
        "Description": "УрФУ",
        "ContentService": "Сервисы",
    },
    "Sidebar": {
        "Home": "Главная",
        "MainMenu": "Главное меню",
        "Users": "Пользователи",
        "Servers": "Серверы",
        "AnyList": "Список",
        "Profile": "Профиль",
        "Edit": "Редактирование",
        "Save": "Сохранить",
    },
    "LanguageSwitcher": {
        "LanguageSwitch": "Выберите язык",
    },
    "UserNavBar": {
        "SignedAs": "Вход по",
        "Logout": "Выйти",
        "PasswordChange": "Сменить пароль",
        "LanguageChange": "Сменить язык",
    },
    "Tables": {
        "UsersTable": {
            "Title": "Все пользователи",
            "SearchBar": "Найти пользователя",
            "ButtonAdd": "Создать",
            "ButtonEdit": "Редактировать пользователя",
            "Columns": [
                {name: 'ID', uid: 'id'},
                {name: 'ИМЯ', uid: 'name'},
                {name: 'РОЛЬ', uid: 'role'},
                {name: 'ГРУППА', uid: 'group'},
                {name: 'СТАТУС', uid: 'status'},
                {name: 'ДЕЙСТВИЕ', uid: 'actions'},
            ]
        }
    },
    "Forms": {
        "SaveSuccess": "Сохранено",
        "SaveError": "Ошибка при сохранении",
    },
    "UserForm": {
        "FieldID": "ID",
        "FieldName": "Имя",
        "FieldEmail": "Почта",
        "FieldUserRole": "Роль",
        "FieldUserRoleStudent": "Студент",
        "FieldUserRoleInstructor": "Преподаватель",
        "FieldUserRoleAdmin": "Администратор",
        "FieldGroupName": "Группа",
        "FieldExternalLTIID": "LTI ID",
        "FieldIsActive": "Активный",
        "FieldCreatedAt": "Создан в",
        "FieldUpdatedAt": "Обновлен в"
    },
    "UserFormPasswordChange": {
        "FieldOldPassword": "Текущий пароль",
        "FieldNewPassword": "Новый пароль",
        "FieldAgainPassword": "Подтверждение нового пароля",
    }
}
export default ru