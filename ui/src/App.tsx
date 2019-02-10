import Router from 'preact-router';
import { h, render } from 'preact';
import RepositoriesPage from './pages/Repositories';
import RepositoryPage from './pages/Repository';
/** @jsx h */

const Main = () => (
  <div class="container">
    <Router>
      <RepositoriesPage default />
      <RepositoryPage path='/repo/:repoId' />
    </Router>
  </div>
);

render(<Main />, document.body);
