import { h, Component } from 'preact';
import { route } from 'preact-router';
import RepositoryAPI from '../api/repositories';
/** @jsx h */

interface RepositoriesState {
  summaries: any[]
}

export default class Repositories extends Component<any, RepositoriesState> {
  private refreshTimer = 0;
  constructor(props: any) {
    super(props)

    this.state = {
      summaries: []
    }
  }

  fetchSummaries() {
    this.setState({
      summaries: RepositoryAPI.GetAllSummaries()
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

  clickRepo(id: string) {
    console.log(`boo ${id}`);
    route(`/repo/${id}`)
  }

  render(props: any, state: RepositoriesState) {
    let boxes = state.summaries.map(s => {
      return <div key={s.id} onClick={(e) => {this.clickRepo(s.id)}} class={"repo-summary-box " + s.lastBuild.status}>
        <span class="repo-name">
          {s.name}
        </span>
      </div>;
    })
    return <div class="boxes">{boxes}</div>
  }
}

