import { h, Component } from 'preact';
import { route } from 'preact-router';
import BuildAPI from '../api/build';
/** @jsx h */

interface ProjectsState {
  latestBuilds: BuildHistory[]
}

export default class ProjectsPage extends Component<any, ProjectsState> {
  private refreshTimer = 0;
  constructor(props: any) {
    super(props)

    this.state = {
      latestBuilds: []
    }
  }

  fetchSummaries() {
    BuildAPI.GetLatestBuilds().then((builds) => {
      this.setState({ latestBuilds: builds })
    })
  }

  componentDidMount() {
    this.fetchSummaries()
    this.refreshTimer = setInterval(() => {
      this.fetchSummaries()
    }, 10000); // 10 seconds
  }  

  componentWillUnmount() {
    clearInterval(this.refreshTimer);
  }

  clickProject(projectName: string) {
    route(`/project/${projectName}`)
  }

  render(props: any, state: ProjectsState) {
    let boxes = state.latestBuilds.map(b => {
      return <div key={b.projectName} onClick={(e) => {this.clickProject(b.projectName)}} class={"project-boxes " + b.status}>
        <div class="project-name">
          {b.projectName}
          <span>
            &nbsp;-&nbsp;{b.status === 'success' ? 'PASS' : 'FAIL'}
          </span>
        </div>
      </div>
    })
    return <div class="boxes">{boxes}</div>
  }
}
