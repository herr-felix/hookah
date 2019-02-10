import Router from 'preact-router';
import { h, render } from 'preact';
import ProjectsPage from './pages/Projects';
import ProjectPage from './pages/Project';
/** @jsx h */

const Main = () => (
  <div class="container">
    <Router>
      <ProjectsPage default />
      <ProjectPage path='/project/:projectName' />
    </Router>
  </div>
);

render(<Main />, document.body);
