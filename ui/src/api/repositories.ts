
export default class RepositoryAPI {
  static async GetAllSummaries() : Promise<RepositorySommary[]> {
    return [
      {
        name: "Test",
        id: '23fwa3fa3',
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "success",
        }
      },
      {
        name: "Website",
        id: '8asd7f8f734',
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "failed",
        }
      },
      {
        name: "ProfilNav",
        id: '3498v4hv98',
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "success",
        }
      },
      {
        name: "MeetingDashboard",
        id: '0934fn0vns04',
        lastBuild: {
          id: "2cd73a64a0e10f42c61a38799132afee",
          start: 1549745380,
          duration: 312,
          status: "success",
        }
      }
    ];
  }

  static async GetDetailsById(repoId: string) : Promise<Repository> {
    return {
      id: '2cd73a64a0e10f42c61a38799132afee',
      name: 'Website',
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