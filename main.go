package main

import buildingspaces "./buildingspaces/docker"
import "./sqlitestore"

func main() {
	// Initialize the builspace
	dockerBuildSpace := buildingspaces.NewDockerBuildingSpace()

	store := sqlitestore.NewSqliteStore()

	// Start the server
	server := NewServer(dockerBuildSpace)

	server.Listen()
}
