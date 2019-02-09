import Router from 'preact-router';
import { h, render } from 'preact';
import Repositories from './pages/Repositories';
/** @jsx h */

const Main = () => (
  <div class="container">
    <Router>
      <Repositories default />
    </Router>
  </div>
);

render(<Main />, document.body);
