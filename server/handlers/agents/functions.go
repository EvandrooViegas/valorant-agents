package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

	for _, agent := range agents {
		if shouldIncludeAgent(agent, filters, addedAgents) {
			filteredList = append(filteredList, agent)
			addedAgents[agent.DisplayName] = true
		}
	}
	return filteredList
}

func shouldIncludeAgent(
	agent Agent, 
	filters Filter, 
	addedAgents map[string]bool,
) bool {
	if addedAgents[agent.DisplayName] {
		return false
	}
	for _, role := range filters.Roles {
		if role.Name == agent.Role.Name {
			return true
		}
	}
	if filters.Name != "" && strings.Contains(strings.ToLower(agent.DisplayName), strings.ToLower(filters.Name)) {
		return true
	}

	return false
}