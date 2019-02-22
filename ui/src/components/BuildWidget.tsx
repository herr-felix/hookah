import { h, Component } from 'preact';
/** @jsx h */

interface BuildWidgetProps {
  build: BuildHistory
}

interface BuildWidgetState {
  toggleOpen: boolean
}

export default class BuildWidget extends Component<BuildWidgetProps, BuildWidgetState> {
  constructor(props: BuildWidgetProps) {
    super(props)
    this.state = {
      toggleOpen: false
    }
  }

  flip() {
    this.setState({toggleOpen: !this.state.toggleOpen});
  }

  render(props: BuildWidgetProps, state: BuildWidgetState) {
    return <div class="repo-details">
      <div class={"panel " + this.props.build.status}>
        <div class="head clickable" onClick={this.flip.bind(this)}>
          <span>
            {this.props.build.status === 'success' ? 'PASS' : 'FAIL'}&nbsp;-&nbsp;
          </span>
          {this.props.build.name}
        </div>
        <div class="content">
          <div class="code">{this.props.build.output}</div>
        </div>
      </div>
    </div>
  }
}