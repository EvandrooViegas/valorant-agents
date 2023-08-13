package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"valorant-agents/utils"
)

func GetAgents() ([]Agent, error) {
	request, err := http.Get("https://valorant-api.com/v1/agents")
	utils.HandleErr(err)
	defer request.Body.Close()

	data, err := ioutil.ReadAll(request.Body)
	utils.HandleErr(err)

	var response Response
	err = json.Unmarshal(data, &response)
	utils.HandleErr(err)

	return response.Data, nil
}

func FilterAgents(filters Filter) []Agent {
	var filteredList []Agent
	addedAgents := make(map[string]bool)

	agents, err := GetAgents()
	utils.HandleErr(err)

	if filters.Name == "" && len(filters.Roles) == 4 {
		return agents
	} else {
		for _, agent := range agents {
			if shouldIncludeAgent(agent, filters, addedAgents) {
				filteredList = append(filteredList, agent)
				addedAgents[agent.DisplayName] = true
			}
		}
		return filteredList
	}
}

func shouldIncludeAgent(
	agent Agent,
	filters Filter,
	addedAgents map[string]bool,
) bool {
	// if the agent was already added, then ignore it
	if addedAgents[agent.DisplayName] {
		return false
	}

	// if the name filter was not passed, filter by the roles
	if filters.Name == "" {
		for _, role := range filters.Roles {
			if role.Name == agent.Role.Name { return true }
		}
	// if the roles filter was not passed, filter by the name
	} else if len(filters.Roles) == 0 {
		if utils.StringContains(agent.DisplayName, filters.Name) { return true }
		
	} else if utils.StringContains(agent.DisplayName, filters.Name) {
		for _, role := range filters.Roles {
			if role.Name == agent.Role.Name { return true }
		}
	}
	return false
}
