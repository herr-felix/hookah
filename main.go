package main

import buildingspaces "./buildingspaces/docker"
import "./sqlitestore"

func main() {
	// Initialize the builspace
	dockerBuildSpace := buildingspaces.NewDockerBuildingSpace()

	queuer := NewBuildQueuer(dockerBuildSpace)

	// Initialize the store
	store, err := sqlitestore.NewSqliteStore("./hookah.db")
	if err != nil {
		panic(err)
	}

	// Start the server
	server := NewServer(queuer, store)

	server.Listen()
}
