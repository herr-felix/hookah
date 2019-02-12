
export default class BuildAPI {

  static async getOutputById(buildId: string): Promise<string> {
    return await fetch(`/build/output/${buildId}`).then((resp) => resp.text())
  }

  static async GetLatestBuilds(): Promise<BuildHistory[]> {
    return await fetch('/builds').then((resp) => resp.json())
  }

  static async GetBuildsByProject(projectName: string): Promise<BuildHistory[]> {
    return await fetch(`/builds/${projectName}`).then((resp) => resp.json())
  }

}