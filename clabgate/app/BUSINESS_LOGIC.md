# ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which
caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers on validate and send to UseCase layer (used in routes)
- `./app/models` folder for describe business models of your project
- `./app/queries` folder for describe queries for models of your project
- `./app/di` folder for Depends container
- `./app/usecases` folder for login and func controller
