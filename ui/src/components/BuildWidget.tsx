import { h, Component } from 'preact';
import BuildAPI from '../api/build';
/** @jsx h */

interface BuildWidgetProps {
  build: BuildHistory
}

interface BuildWidgetState {
  toggleOpen: boolean
  output: string
}

export default class BuildWidget extends Component<BuildWidgetProps, BuildWidgetState> {
  constructor(props: BuildWidgetProps) {
    super(props)
    this.state = {
      toggleOpen: false,
      output: null
    }
  }

  fetchBuildOutput() {
    BuildAPI.getOutputById(this.props.build.id).then((output) => {
      let state = this.state as BuildWidgetState
      state.output = output
      this.setState(state)
    })
  }

  renderBuildDetails() {
    if (!this.state.toggleOpen) {
      return null
    }
    const outputPart = this.state.output ? <div class="code">{this.state.output}</div> : 'Loading...'
    return <div class="content">
      {outputPart}
    </div>
  }

  flipOpen() {
    let state = this.state as BuildWidgetState
    state.toggleOpen = !state.toggleOpen
    if (state.output === null) {
      this.fetchBuildOutput()
    }
    this.setState(state);
  }

  render(props: BuildWidgetProps, state: BuildWidgetState) {
    return <div class="repo-details">
      <div class={"panel " + this.props.build.status}>
        <div class="head clickable" onClick={this.flipOpen.bind(this)}>
          {(new Date(this.props.build.start*1000)).toJSON()}
        </div>
        {this.renderBuildDetails()}
      </div>
    </div>
  }
}