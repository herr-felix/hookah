
interface BuildHistory {
  id: string,
  start: number,
  duration: number,
  status: string
}

interface RepositorySommary {
  id: string,
  name: string,
  lastBuild: BuildHistory
}

interface Repository {
  id: string,
  name: string
  builds: Array<BuildHistory>
}