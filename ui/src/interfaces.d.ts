
interface BuildHistory {
  id: string,
  start: number,
  duration: number,
  status: string
}

interface ProjectSommary {
  name: string,
  lastBuild: BuildHistory
}

interface Project {
  projectName: string
  builds: Array<BuildHistory>
}