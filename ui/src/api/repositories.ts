
export default class RepositoryAPI {
  static async GetAllSummaries() : Promise<ProjectSommary[]> {
    return [
      {
        name: "Test",
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "success",
        }
      },
      {
        name: "Website",
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "failed",
        }
      },
      {
        name: "ProfilNav",
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "success",
        }
      },
      {
        name: "MeetingDashboard",
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "success",
        }
      }
    ];
  }

  static async GetDetailsByName(name: string) : Promise<Project> {
    return {
      projectName: 'Website',
      builds: [
        {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 343,
          status: "success"
        },
        {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1542745380,
          duration: 343,
          status: "failed"
        },
        {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1541744380,
          duration: 343,
          status: "success"
        },
        {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1533745380,
          duration: 343,
          status: "success"
        },
        {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1522745380,
          duration: 343,
          status: "failed"
        }
      ]
    }
  }
}