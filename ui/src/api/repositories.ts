
export default
  class RepositoryAPI {
  static GetAllSummaries() {
    return [
      {
        name: "Test",
        id: '23fwa3fa3',
        lastBuild: {
          duration: 312,
          status: "success",
          lastOutputLine: "Tigidou"
        }
      },
      {
        name: "Website",
        id: '8asd7f8f734',
        lastBuild: {
          duration: 312,
          status: "failed",
          lastOutputLine: "Tigidou"
        }
      },
      {
        name: "ProfilNav",
        id: '3498v4hv98',
        lastBuild: {
          duration: 312,
          status: "success",
          lastOutputLine: "Tigidou"
        }
      },
      {
        name: "MeetingDashboard",
        id: '0934fn0vns04',
        lastBuild: {
          duration: 312,
          status: "success",
          lastOutputLine: "Tigidou"
        }
      }
    ];
    
  }
}