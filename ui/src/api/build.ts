
export default class BuildAPI {

  static async GetLatestBuilds(): Promise<BuildHistory[]> {
    return await fetch('/builds').then((resp) => resp.json())
  }

  static async GetBuildsByProject(projectName: string): Promise<BuildHistory[]> {
    return await fetch(`/builds/${projectName}`).then((resp) => resp.json())
  }

}