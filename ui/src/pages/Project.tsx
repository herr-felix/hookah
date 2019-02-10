import { h, Component } from 'preact';
import BuildWidget from '../components/BuildWidget';
import BuildAPI from '../api/build';
import { DateTime } from "luxon";
/** @jsx h */

interface ProjectProps {
  projectName?: string
}

interface ProjectState {
  builds: BuildHistory[]
}

export default class ProjectPage extends Component<ProjectProps, ProjectState> {
  constructor(props: ProjectProps) {
    super(props)
    this.state = {
      builds: []
    }
  }

  fetchProject() {
    BuildAPI.GetBuildsByProject(this.props.projectName).then((builds) => {
      this.setState({builds: builds});
    })
  }

  componentDidMount() {
    this.fetchProject();
  }

  render(props: ProjectProps, state: ProjectState) {
    const builds = state.builds.map(b => <BuildWidget build={b}/>);

    if (state.builds.length === 0) {
      return <h2>Loading...</h2>
    }

    return <div>
      <div class="panel">
        <div class="head title">{props.projectName}</div>
        <div className="content">
        <table>
          <tr>
            <td>
              <strong>Last build:&nbsp;</strong>
              {DateTime.fromSeconds(state.builds[0].start).toFormat("yyyy LLL dd HH:mm:ss")}
            </td>
          </tr>
        </table>
        </div>
      </div>
      {builds}
    </div>
  }
}
