import { h, Component } from 'preact';
import RepositoryAPI from '../api/repositories';
import BuildWidget from '../components/BuildWidget';
import { DateTime } from "luxon";
/** @jsx h */

interface RepositoryProps {
  repoId?: string;
}

interface RepositoryState {
  repo: Repository
}

export default class RepositoryPage extends Component<RepositoryProps, RepositoryState> {
  constructor(props: RepositoryProps) {
    super(props)
    this.state = {
      repo: {
        id: null,
        name: '',
        builds: []
      }
    }
  }

  fetchRepo() {
    RepositoryAPI.GetDetailsById(this.props.repoId).then((repo) => {
      this.setState({repo: repo});
    })
  }

  componentDidMount() {
    this.fetchRepo();
  }

  render(props: RepositoryProps, state: RepositoryState) {
    const builds = state.repo.builds.map(b => <BuildWidget build={b}/>);

    if (state.repo.id === null) {
      return <h2>Loading...</h2>
    }

    return <div class="repo-details">
      <div class="panel">
        <div class="head title">{state.repo ? state.repo.name : ''}</div>
        <div className="content">
        <table>
          <tr>
            <td>
              <strong>Last build:&nbsp;</strong>
              {DateTime.fromSeconds(state.repo.builds[0].start).toFormat("yyyy LLL dd HH:mm:ss")}
            </td>
          </tr>
        </table>
        </div>
      </div>
      {builds}
    </div>
  }
}
