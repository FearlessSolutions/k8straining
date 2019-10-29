This is the React Reference Architecture  It is designed for a standalone app, but could easily be combined with any other API stack.

The React app was created using create-react-app.  It has been augmented with Storybook and Cypress to ease the development process.  [Storybook]("https://storybook.js.org/" "Storybook Homepage") is a quick, easy way to view components in isolation.  [Cypress]("http://cypress.io" "Cypress Homepage") is a pure javascript integration test tool.

Helpful Scripts:

To accommodate a container first development philosophy and remove the requirement to install any dependencies locally you can run  `./scripts/container-npm.sh "<command here>"` which will run NPM commands inside your containers.  For example to add [React-bootstrap](https://react-bootstrap.github.io/) run the following:

`./scripts/container-npm.sh "install react-bootstrap"`


Make Commands:
* `make run_app` - runs just the [app](http://localhost:3000), with hot reloading whenever changes are made
* `make test_app` - runs Cypress and jest test suites in their containers
* `make run_app_dev` -  runs the [app](http://localhost:3000), and [storybook](http://localhost:9009). This is probably the best default to just leave running. 
* `make down` - stops all running docker containers